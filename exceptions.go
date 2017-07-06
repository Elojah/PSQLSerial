/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   exceptions.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/05 17:00:12 by leeios            #+#    #+#             */
/*   Updated: 2017/07/06 17:52:29 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

type UnconvertibleType struct {
	typename string
}

func (u UnconvertibleType) Error() string {
	return (fmt.Sprintf(`Type %s is not convertible into SQL type`, u.typename))
}

type NoColumnSpecified struct{}

func (n NoColumnSpecified) Error() string {
	return (`No cols field in query`)
}

type BadFormatColumns struct{}

func (b BadFormatColumns) Error() string {
	return (`Cols field is not an array`)
}

type BadFormatWhere struct{}

func (b BadFormatWhere) Error() string {
	return (`Where field is not a map`)
}

type ColumnNotString struct{}

func (c ColumnNotString) Error() string {
	return (`A column name is not a string`)
}

type ExpectedStructure struct{}

func (e ExpectedStructure) Error() string {
	return (`Expected a structure`)
}

type NotExistingColumn struct {
	col string
}

func (n NotExistingColumn) Error() string {
	return (fmt.Sprintf(`Column %s doesn't exist`, n.col))
}
