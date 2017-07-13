/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   types.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/06 11:37:14 by leeios            #+#    #+#             */
/*   Updated: 2017/07/10 11:57:13 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */
package PSQLSerial

type ID int
type Constraint int
type Index int

type Condition struct {
	And []Condition `json:"and"`
	Or  []Condition `json:"or"`
	Col string      `json:"col"`
	Op  string      `json:"op"`
	Rhs string      `json:"rhs"`
}

type Conditions struct {
	tables map[string]Condition
}
