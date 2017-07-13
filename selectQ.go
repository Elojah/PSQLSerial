/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   selectQ.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/06 11:35:21 by leeios            #+#    #+#             */
/*   Updated: 2017/07/10 20:33:35 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package PSQLSerial

import (
	"encoding/json"
	"fmt"
	"github.com/elojah/PSQLSerial/utils"
	"reflect"
	"strings"
)

func serializeSelectForeign(v reflect.Value) (string, error) {
	return ``, nil
}

func serializeSelectWhere(v reflect.Value, where interface{}, tablename string) (string, error) {
	if v.Kind() != reflect.Struct {
		return ``, ExpectedStructure{}
	}
	wclauses, ok := where.(map[string]interface{})
	if !ok {
		return ``, BadFormatWhere{}
	}
	wclauses_split := []string{}
	for key, wclause := range wclauses {
		if key == `and` || key == `or` {
			andclausesJson, ok := wclause.([]interface{})
			if !ok {
				return ``, BadFormatWhere{}
			}
			andclauses := []string{}
			for _, val := range andclausesJson {
				andclause, err := serializeSelectWhere(v, val, tablename)
				if err != nil {
					return ``, err
				}
				andclauses = append(andclauses, andclause)
			}
			wclauses_split = append(wclauses_split, strings.Join(andclauses, ` `+strings.ToUpper(key)+` `))
		} else {
			if !v.FieldByName(key).IsValid() {
				return ``, NotExistingColumn{col: key}
			}
			clauseJson, ok := wclause.(map[string]interface{})
			if !ok {
				return ``, BadFormatWhere{}
			}
			k, v := utils.TakeFirst(clauseJson)
			op, ok := k.(string)
			if !ok {
				return ``, BadFormatWhere{}
			}
			rhs, ok := v.(string)
			if !ok {
				return ``, BadFormatWhere{}
			}
			clause := utils.Join(` `, utils.Join(`.`, tablename, strings.ToLower(key)), op, rhs)
			wclauses_split = append(wclauses_split, clause)
		}
	}
	if len(wclauses_split) == 0 {
		return ``, nil
	}
	return `WHERE ` + strings.Join(wclauses_split, ` AND `), nil
}

func serializeSelectAllColumns(v reflect.Value, tablename string) (columns []string) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tags := field.Tag
		reftable := tags.Get(`sql_ref_table`)
		nntable := tags.Get(`sql_nn_table`)
		fktable := tags.Get(`sql_fk_table`)
		if reftable != `` || nntable != `` || fktable != `` {
			continue
		}
		columns = append(columns, utils.Join(`.`, tablename, strings.ToLower(field.Name)))
	}
	return
}

func serializeSelectCols(v reflect.Value, colJson interface{}, tablename string) (string, error) {
	if v.Kind() != reflect.Struct {
		return ``, fmt.Errorf(`Can't serialize a non-struct type`)
	}
	cols_split := []string{}
	for _, col := range cols {
		strcol, ok := col.(string)
		if !ok {
			return ``, ColumnNotString{}
		}
		if strcol == `*` {
			cols_split = append(cols_split, serializeSelectAllColumns(v, tablename)...)
			continue
		}
		if !v.FieldByName(strcol).IsValid() {
			return ``, NotExistingColumn{col: strcol}
		}
		cols_split = append(cols_split, utils.Join(`.`, tablename, strings.ToLower(strcol)))
	}
	return strings.Join(cols_split, `,`), nil
}

func serializeSelect(v reflect.Value, jsonQuery map[string]interface{}) (columns string, err error) {
	if v.Kind() != reflect.Struct {
		return ``, fmt.Errorf(`Can't serialize a non-struct type`)
	}
	tablename := strings.ToLower(v.Type().Name())

	return utils.Join(` `, `SELECT`, cols, `FROM`, tablename, joinclauses), nil
}

/*
filterCols is a json with allowed keys
{
	"id": "",
	"name": "",
	"image": {
		"id": "",
		"content": ""
	},
	"group": {
	}
}
e.g:
{
	"cols": ["*"],
	"foreign": {
		"author": {
			"cols": ["name"],
			"where": {
				"and": [{
					"last_update": {
						">": "2015-10-21"
					},
					"creation_date": {
						"<": "1987-02-21"
					}]
				}
			}
		}
	}
}
*/
func SelectQ(obj interface{}, jsonQuery []byte) (string, error) {
	var filter map[string]interface{}

	v := reflect.ValueOf(obj)
	err := json.Unmarshal(jsonQuery, &filter)
	if err != nil {
		return ``, err
	}
	return serializeSelect(v, filter)
}
