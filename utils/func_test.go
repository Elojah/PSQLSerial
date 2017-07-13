/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   func_test.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/09 20:26:21 by leeios            #+#    #+#             */
/*   Updated: 2017/07/10 11:16:08 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"testing"
)

func TestJoin(t *testing.T) {
	if Join(`.`, `a`, `b`, `c`) != `a.b.c` {
		t.Fail()
	}
}

func TesttakeFirst(t *testing.T) {
	slice := map[string]interface{}{
		`a`: `z`,
	}
	k, v := TakeFirst(slice)
	if k != `a` || v.(string) != `z` {
		t.Fail()
	}
}
