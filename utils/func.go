/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   func.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/09 19:42:44 by leeios            #+#    #+#             */
/*   Updated: 2017/07/10 11:16:46 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"strings"
)

func Join(sep string, elems ...string) string {
	acc := []string{}
	return strings.Join(append(acc, elems...), sep)
}

func TakeFirst(obj map[string]interface{}) (interface{}, interface{}) {
	for k, v := range obj {
		return k, v
	}
	return nil, nil
}
