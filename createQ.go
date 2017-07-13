/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   createQ.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/05 14:55:57 by leeios            #+#    #+#             */
/*   Updated: 2017/07/09 19:48:41 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package PSQLSerial

import (
	"reflect"
	"strings"
)

func serializeCreateCol(name string, typename string, attr string) (string, error) {
	gotypeToSQLType := map[string]string{
		`ID`:     `SERIAL`,
		`string`: `VARCHAR`,
		`int`:    `int`,
		`bool`:   `bool`,
		`Time`:   `timestamptz`,
	}
	sqltype, ok := gotypeToSQLType[typename]
	if !ok {
		return ``, UnconvertibleType{typename: typename}
	}
	col := []string{strings.ToLower(name), sqltype, attr}
	return strings.Join(col, ` `), nil
}

func serializeCreateRefTable(reftable string, ondelete string) string {
	if ondelete == `` {
		ondelete = `RESTRICT`
	}
	attr := []string{
		`REFERENCES`,
		reftable,
		`ON DELETE`,
		ondelete,
	}
	return strings.Join(attr, ` `)
}

func serializeCreate(v reflect.Value) (columns string, indexes string, err error) {
	columns_split := []string{}
	indexes_split := []string{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		name := field.Name
		tags := field.Tag
		typename := field.Type.Name()

		switch typename {
		case `Constraint`:
			columns_split = append(columns_split, tags.Get(`sql_attr`))
		case `Index`:
			indexes_split = append(indexes_split, tags.Get(`sql_attr`))
		default:
			reftable := tags.Get(`sql_ref_table`)
			nntable := tags.Get(`sql_nn_table`)
			fktable := tags.Get(`sql_fk_table`)

			if nntable != `` || fktable != `` {
				// Those columns are not represented as SQL columns so just ignore them
			} else if reftable != `` {
				// This is a foreign key so set an id as references
				column, err := serializeCreateCol(name, `ID`, serializeCreateRefTable(reftable, tags.Get(`sql_on_delete`)))
				if err != nil {
					return ``, ``, err
				}
				columns_split = append(columns_split, column)
			} else {
				// This is a regular column with no external dependencies
				column, err := serializeCreateCol(name, typename, tags.Get(`sql_attr`))
				if err != nil {
					return ``, ``, err
				}
				columns_split = append(columns_split, column)
			}
		}
	}
	return strings.Join(columns_split, `,`), strings.Join(indexes_split, `,`), nil
}

func CreateQ(obj interface{}) (string, error) {
	v := reflect.ValueOf(obj)
	columns, _, err := serializeCreate(v)
	if err != nil {
		return ``, err
	}
	queryterms := []string{
		`CREATE TABLE`,
		v.Type().Name(),
		`(`,
		columns,
		`);`,
	}
	query := strings.Join(queryterms, ` `)
	return query, nil
}
