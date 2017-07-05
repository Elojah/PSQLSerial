/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   structs_test.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: leeios <leeios@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/07/05 14:55:58 by leeios            #+#    #+#             */
/*   Updated: 2017/07/05 15:54:59 by leeios           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"time"
)

/*
	Basic
	sql_attr: Attribute SQL used in column creation

	N1 Joins
	sql_ref_table: For ID (foreign key) only. Foreign table name

	NN joins
	sql_nn_table: Jointable name for n<->n links
	sql_nn_self: Jointable column name referencing object id
	sql_nn_field: Jointable column name referencing field id

	1N joins
	sql_fk_table: Table where self is referenced as fk id
	sql_fk_self: Column name where self is referenced as fk id
*/

/*
	Domain implementation
	- Keep a newline between actual columns and nn joins fields
*/

/*TODO !!!
- Add indexes
- Find a way to specify a call with multiple join (other than nn or 1N/N1)
- Add deletion modes
*/

// Classic user model, rights in UserGroup
type User struct {
	Id       ID     `sql_attr:"PRIMARY_KEY"`
	Name     string `sql_attr:"NOT NULL"`
	Mail     string `sql_attr:"NOT NULL UNIQUE"`
	Password string `sql_attr:"NOT NULL"`

	Groups []UserGroup `sql_nn_table:"user_fk_usergroup",
							sql_nn_self:"user_id",
							sql_nn_field:"usergroup_id"`
	// Add formations/sessions/classes when multiple joins ok
}

type User_fk_UserGroup struct {
	User      User      `sql_attr:"NOT NULL" sql_ref_table:"user"`
	Usergroup UserGroup `sql_attr:"NOT NULL" sql_ref_table:"usergroup"`
}

// UserGroup with associated rights defined in validators
type UserGroup struct {
	Id   ID     `sql_attr:"PRIMARY_KEY"`
	Name string `sql_attr:"NOT NULL UNIQUE"`

	Users []User `sql_nn_table:"user_fk_usergroup",
					sql_nn_self:"usergroup_id",
					sql_nn_field:"user_id"`
	SessionsTeacher []Session `sql_fk_table:"session",
								sql_fk_self:"teachergroup_id"`
	SessionsStudent []Session `sql_fk_table:"session",
								sql_fk_self:"studentgroup_id"`
	ClassesTeacher []Class `sql_fk_table:"class",
							sql_fk_self:"teachergroup_id"`
	ClassesStudent []Class `sql_fk_table:"class",
							sql_fk_self:"studentgroup_id"`
}

// A formation is date independant, only basic informations
type Formation struct {
	Id          ID     `sql_attr:"PRIMARY_KEY"`
	Name        string `sql_attr:"NOT NULL"`
	Description string ``

	DefaultEvaluation Evaluation `sql_ref_table:"evaluation"`

	Sessions []Session `sql_fk_table:"session",
						sql_fk_self:"formation_id"`
}

// A session is a group of multiple classes e.g: a school year or a 2 week formation
// It has his own evaluations
// TeacherGroup and StudentGroup are the common factors between classes
type Session struct {
	// insert row "ALL" (id=0) in Session
	Id         ID        `sql_attr:"PRIMARY_KEY"`
	Start_time time.Time `sql_attr:"NOT NULL"`
	End_time   time.Time `sql_attr:"NOT NULL"`
	Comment    string    `` // e.g: 'Year 2016/2017'

	Formation    Formation `sql_attr:"NOT NULL" sql_ref_table:"formation"`
	TeacherGroup UserGroup `sql_ref_table:"usergroup"` // May be NULL if always defined in Class
	StudentGroup UserGroup `sql_ref_table:"usergroup"` // May be NULL if always defined in Class

	Classes []Class `sql_fk_table:"class",
						sql_fk_self:"session_id"`
	Evaluations []Evaluation `sql_nn_table:"session_fk_evaluation",
								sql_nn_self:"session_id",
								sql_nn_field:"evaluation_id"`

	Check_dates_c Constraint `sql_attr:"CHECK (start_time < end_time)"`
}

// A class is a short period course, timeline will be shown on calendar
type Class struct {
	Start_time time.Time `sql_attr:"NOT NULL"`
	End_time   time.Time `sql_attr:"NOT NULL"`
	Comment    string    `` // e.g: 'Saturday morning class'

	Session      Session   `sql_attr:"NOT NULL" sql_ref_table:"session"`
	TeacherGroup UserGroup `sql_ref_table:"usergroup"` // Defaulted to session TeacherGroup
	StudentGroup UserGroup `sql_ref_table:"usergroup"` // Defaulted to session StudentGroup

	Evaluations []Evaluation `sql_nn_table:"class_fk_evaluation",
								sql_nn_self:"class_id",
								sql_nn_field:"evaluation_id"`

	Check_dates Constraint `sql_constraints:"CHECK (start_time < end_time)"`
}

// An evaluation is a group of multiple questions and answers for an author
type Evaluation struct {
	Id          ID     `sql_attr:"PRIMARY_KEY"`
	Name        string `sql_attr:"NOT NULL"`
	Description string ``
	Anonymous   bool   `sql_attr:"NOT NULL"`

	Classes []Class `sql_nn_table:"class_fk_evaluation",
						sql_nn_self:"evaluation_id",
						sql_nn_field:"class_id"`
	Sessions []Session `sql_nn_table:"session_fk_evaluation",
						sql_nn_self:"evaluation_id",
						sql_nn_field:"session_id"`
}

// Evaluation per session e.g: After 1 month, 3 month, post 6 months
type Session_fk_Evaluation struct {
	Id           ID        `sql_attr:"PRIMARY_KEY"`
	Comment      string    `` // e.g: 'Post 3 month evaluation'
	Sending_date time.Time `sql_attr:"NOT NULL"`

	Session    Session    `sql_attr:"NOT NULL" sql_ref_table:"session"`
	Evaluation Evaluation `sql_attr:"NOT NULL" sql_ref_table:"evaluation"`
}

// Evaluation per class. We keep it separate from session evaluation because they have a different meaning
type Class_fk_Evaluation struct {
	Id           ID        `sql_attr:"PRIMARY_KEY"`
	Comment      string    ``                    // e.g: Practice class
	Sending_date time.Time `sql_attr:"NOT NULL"` // Should be defaulted to Class end_time

	Class      Class      `sql_attr:"NOT NULL" sql_ref_table:"class"`
	Evaluation Evaluation `sql_attr:"NOT NULL" sql_ref_table:"evaluation"`
}

type Question struct {
	Id            ID     `sql_attr:"PRIMARY_KEY"`
	RatingAllowed bool   `sql_attr:"NOT NULL"`
	Content       string `sql_attr:"NOT NULL"` // e.g: 'Was it a good formation ?'
}

type Evaluation_fk_Question struct {
	Evaluation Evaluation `sql_attr:"NOT NULL" sql_ref_table:"evaluation"`
	Question   Question   `sql_attr:"NOT NULL" sql_ref_table:"question"`
	N_order    int        `sql_attr:"NOT NULL"`

	Consistent_eval_c Constraint `sql_attr:"UNIQUE(evaluation_id, n_order)"`
}

type Answer struct {
	Id      ID     `sql_attr:"PRIMARY_KEY"`
	Content string `sql_attr:"NOT NULL"` // e.g: 'That was hum... ok'
	Rate    int    ``

	User       User       `sql_attr:"NOT NULL" sql_ref_table:"user"`
	Evaluation Evaluation `sql_attr:"NOT NULL" sql_ref_table:"evaluation"`
	Question   Question   `sql_attr:"NOT NULL" sql_ref_table:"question"`

	Consistent_answer_c Constraint `sql_attr:"UNIQUE(user_id, evaluation_id, question_id)"`
}
