/*******************************************************************/
/*
/* Blog project
/*
/* Copyright (C) 2019 Sergey Denisov.
/*
/* Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
/* Github: https://github.com/LittleBuster
/*	       https://github.com/futcamp
/*
/* This library is free software; you can redistribute it and/or
/* modify it under the terms of the GNU General Public Licence 3
/* as published by the Free Software Foundation; either version 3
/* of the Licence, or (at your option) any later version.
/*
/*******************************************************************/

package connector

import (
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type Base interface {
	Connect() error
	Close()
	Categories() ([]*Category, error)
	Posts(category int) ([]*Post, error)
}

type Database struct {
	engine *xorm.Engine
	ip     string
	user   string
	passwd string
	db     string
}

// NewDatabase make new struct
func NewDatabase(ip string, user string, passwd string, db string) Base {
	return &Database{
		ip:     ip,
		user:   user,
		passwd: passwd,
		db:     db,
	}
}

// Connect connecting to database
func (d *Database) Connect() error {
	var err error

	d.engine, err = xorm.NewEngine("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s",
		d.user, d.passwd, d.ip, d.db))
	if err != nil {
		return err
	}

	return nil
}

// Categories get categories list from database
func (d *Database) Categories() ([]*Category, error) {
	var categories []*Category

	rows, err := d.engine.Table("categories").Rows(&Category{})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		cat := &Category{}

		err := rows.Scan(cat)
		if err != nil {
			return nil, err
		}

		categories = append(categories, cat)
	}

	return categories, nil
}

// Posts get posts list from database
func (d *Database) Posts(category int) ([]*Post, error) {
	var posts []*Post

	rows, err := d.engine.Table("articles").Rows(&Post{Category: category})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &Post{}

		err := rows.Scan(post)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Close closing database
func (d *Database) Close() {
	d.engine.Close()
}
