/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   exceptions.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/05 17:00:12 by leeios            #+#    #+#             */
/*   Updated: 2017/07/05 17:17:52 by leeios           ###   ########.fr       */
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
