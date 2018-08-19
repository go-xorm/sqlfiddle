# SQL Fiddle API (UnOfficial)

## Purpose

This Go library is aimed to provide an API to operate http://sqlfiddle.com/

## Usage

```Go
fiddle := NewFiddle("")
res, err := fiddle.CreateSchema(Mysql5_6, `create table person(id int not null auto_increment,
    name varchar(8),
    birthday datetime,
    constraint pk__person primary key(id));`)

ret, err := fiddle.RunSQL(Mysql5_6, res.Code, "select * from person;")
```