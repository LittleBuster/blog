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

package server

type IndexResponse struct {
	Service string `json:"service"`
	Result  bool   `json:"result"`
}

type ErrorResponse struct {
	Service string `json:"service"`
	Result  bool   `json:"result"`
	Reason  string `json:"reason"`
}

type DataResponse struct {
	Service string      `json:"service"`
	Data    interface{} `json:"data"`
	Result  bool        `json:"result"`
}
