/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main_test.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/05 14:55:55 by leeios            #+#    #+#             */
/*   Updated: 2017/07/09 20:15:17 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package PSQLSerial

import (
	"fmt"
	"testing"
)

/*
User_fk_UserGroup
User
UserGroup
Formation
Session
Class
Evaluation
Session_fk_Evaluation
Class_fk_Evaluation
Question
Evaluation_fk_Question
Answer
*/
// func TestCreateQ(t *testing.T) {
// 	fmt.Println(CreateQ(User{}))
// 	fmt.Println(CreateQ(User_fk_UserGroup{}))
// 	fmt.Println(CreateQ(UserGroup{}))
// 	fmt.Println(CreateQ(Formation{}))
// 	fmt.Println(CreateQ(Session{}))
// 	fmt.Println(CreateQ(Class{}))
// 	fmt.Println(CreateQ(Evaluation{}))
// 	fmt.Println(CreateQ(Session_fk_Evaluation{}))
// 	fmt.Println(CreateQ(Class_fk_Evaluation{}))
// 	fmt.Println(CreateQ(Question{}))
// 	fmt.Println(CreateQ(Evaluation_fk_Question{}))
// 	fmt.Println(CreateQ(Answer{}))
// }

func TestSelectQ(t *testing.T) {
	fmt.Println(SelectQ(User{}, []byte(
		`{"cols": ["*"]}`,
	)))
	fmt.Println(SelectQ(Formation{}, []byte(
		`{"cols": ["Name"]}`,
	)))
	fmt.Println(SelectQ(Evaluation{}, []byte(
		`{"cols": ["Name", "Description"],
			"where": {"Anonymous": {"is": "TRUE"}}
		}`,
	)))

}
