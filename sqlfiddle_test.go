package sqlfiddle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiddle(t *testing.T) {
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
}
