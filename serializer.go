/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   serializer.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/05 14:55:57 by leeios            #+#    #+#             */
/*   Updated: 2017/07/05 17:36:38 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type ID int
type Constraint int
type Index int

type table_attributes struct {
	Cols        []string
	Constraints []string
	Indexes     []string
}

func serializeColumn(name string, typename string, attr string) (string, error) {
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

func serializeRefTableAttr(reftable string, ondelete string) string {
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

func getTableAttributes(obj interface{}) (attr table_attributes, err error) {
	v := reflect.ValueOf(obj)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		name := field.Name
		tags := field.Tag
		typename := field.Type.Name()

		switch typename {
		case `Constraint`:
			attr.Constraints = append(attr.Constraints, tags.Get(`sql_attr`))
		case `Index`:
			attr.Indexes = append(attr.Indexes, tags.Get(`sql_attr`))
		default:
			reftable := tags.Get(`sql_ref_table`)
			nntable := tags.Get(`sql_nn_table`)
			fktable := tags.Get(`sql_fk_table`)

			if nntable != `` || fktable != `` {
				// Those fields are not represented as SQL columns so just ignore them
			} else if reftable != `` {
				// This is a foreign key so set an id as references
				column, err := serializeColumn(name, `ID`, serializeRefTableAttr(reftable, tags.Get(`sql_on_delete`)))
				if err != nil {
					return attr, err
				}
				attr.Cols = append(attr.Cols, column)
			} else {
				// This is a regular column with no external dependencies
				column, err := serializeColumn(name, typename, tags.Get(`sql_attr`))
				if err != nil {
					return attr, err
				}
				attr.Cols = append(attr.Cols, column)
			}
		}
	}
	return attr, nil
}

func CreateTableQuery(obj interface{}) error {
	attr, err := getTableAttributes(obj)
	if err != nil {
		return err
	}
	fmt.Println(strings.Join(attr.Cols, `,`))
	return nil
}
