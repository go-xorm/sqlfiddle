// Copyright 2018 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlfiddle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiddle_CreateSchema4Mysql(t *testing.T) {
	fiddle := NewFiddle("")
	res, err := fiddle.CreateSchema(Mysql5_6, `create table person(id int not null auto_increment,
		name varchar(8),
		birthday datetime,
		constraint pk__person primary key(id));`)
	assert.NoError(t, err)
	fmt.Println(res)

	ret, err := fiddle.RunSQL(Mysql5_6, res.Code, "select * from person;")
	assert.NoError(t, err)
	fmt.Println(ret)

	ret, err = fiddle.RunSQL(Mysql5_6, res.Code, "select * from person1;")
	assert.Error(t, err)
	fmt.Println(err)
}

func TestFiddle_CreateSchema4Oracle(t *testing.T) {
	fiddle := NewFiddle("")
	res, err := fiddle.CreateSchema(Oracle11gR2, `create table table1(
       id number(9) not null primary key,
       a varchar2(40),
       b varchar2(40),
       c varchar2(40)
);`)
	assert.NoError(t, err)
	fmt.Println(res)

	ret, err := fiddle.RunSQL(Oracle11gR2, res.Code, "select * from table1;")
	assert.NoError(t, err)
	fmt.Println(ret)
}

func TestFiddle_CreateSchema4MssSQL(t *testing.T) {
	fiddle := NewFiddle("")
	res, err := fiddle.CreateSchema(Oracle11gR2, `create table table1(
       id int primary key,
       a varchar(40),
       b varchar(40),
       c varchar(40)
);`)
	assert.NoError(t, err)
	fmt.Println(res)

	ret, err := fiddle.RunSQL(Oracle11gR2, res.Code, "select * from table1;")
	assert.NoError(t, err)
	fmt.Println(ret)
}
