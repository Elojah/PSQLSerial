/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   serializer_test.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/05 14:55:55 by leeios            #+#    #+#             */
/*   Updated: 2017/07/05 17:32:23 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

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
func TestCreateTableQuery(t *testing.T) {
	fmt.Println(CreateTableQuery(User{}))
	fmt.Println(CreateTableQuery(User_fk_UserGroup{}))
	fmt.Println(CreateTableQuery(UserGroup{}))
	fmt.Println(CreateTableQuery(Formation{}))
	fmt.Println(CreateTableQuery(Session{}))
	fmt.Println(CreateTableQuery(Class{}))
	fmt.Println(CreateTableQuery(Evaluation{}))
	fmt.Println(CreateTableQuery(Session_fk_Evaluation{}))
	fmt.Println(CreateTableQuery(Class_fk_Evaluation{}))
	fmt.Println(CreateTableQuery(Question{}))
	fmt.Println(CreateTableQuery(Evaluation_fk_Question{}))
	fmt.Println(CreateTableQuery(Answer{}))
}
