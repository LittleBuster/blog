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

package configs

import (
	"encoding/json"
	"io/ioutil"
)

type Configurable interface {
	LoadFromFile(fileName string) error
	Settings() *RouterSettings
}

type Configs struct {
	settings RouterSettings
}

// NewConfigs make new struct
func NewConfigs() Configurable {
	return &Configs{
	}
}

// LoadFromFile loading configs from file
func (c *Configs) LoadFromFile(fileName string) error {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c.settings)
	if err != nil {
		return err
	}

	return nil
}

// Settings get configs pointer
func (c *Configs) Settings() *RouterSettings {
	return &c.settings
}
