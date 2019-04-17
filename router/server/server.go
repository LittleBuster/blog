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

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/LittleBuster/blog/router/services/post"
	"github.com/LittleBuster/blog/router/utils/configs"

	"github.com/buaazp/fasthttprouter"
	"github.com/google/logger"
	"github.com/valyala/fasthttp"

)

type Server interface {
	Start(ip string, port int, api string) error
}

type WebServer struct {
	configs configs.Configurable
}

// NewWebServer make new struct
func NewWebServer(cfg configs.Configurable) Server {
	return &WebServer{
		configs: cfg,
	}
}

// Start start web server
func (w *WebServer) Start(ip string, port int, api string) error {
	router := fasthttprouter.New()

	router.GET("/", w.IndexHandler)
	router.GET(fmt.Sprintf("/api/%s/post/categories", api), w.CatHandler)
	router.GET(fmt.Sprintf("/api/%s/post/articles/:cat", api), w.ArtHandler)

	return fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), router.Handler)
}

//ErrorResponse send error response
func (w *WebServer) ErrorResponse(ctx *fasthttp.RequestCtx, reason string) {
	errResp := ErrorResponse{
		Service: "post",
		Result:  false,
		Reason:  reason,
	}
	bytesOut, _ := json.Marshal(&errResp)

	logger.Error(reason)

	ctx.Response.Header.Add("Content-Type", "application/json")
	_, err := ctx.Write(bytesOut)
	if err != nil {
		logger.Error(fmt.Sprintf("Fail to send response"))
	}
}

// IndexHandler main server handler
func (w *WebServer) IndexHandler(ctx *fasthttp.RequestCtx) {
	indexResp := IndexResponse{
		Service: "router",
	}
	bytesOut, _ := json.Marshal(&indexResp)

	ctx.Response.Header.Add("Content-Type", "application/json")
	_, err := ctx.Write(bytesOut)
	if err != nil {
		logger.Error("IndexHandler: Fail to send response")
	}
}

// CatHandler get categories handler
func (w *WebServer) CatHandler(ctx *fasthttp.RequestCtx) {
	var err error

	client := post.NewPostClient(w.configs.Settings().Post.Ip,
		w.configs.Settings().Post.Port)

	dataResp := DataResponse{
		Service: "post",
		Result:  true,
	}

	dataResp.Data, err = client.Categories()
	if err != nil {
		w.ErrorResponse(ctx, "Fail to get categories data from service")
		return
	}

	bytesOut, _ := json.Marshal(&dataResp)
	ctx.Response.Header.Add("Content-Type", "application/json")

	_, err = ctx.Write(bytesOut)
	if err != nil {
		logger.Error("CatHandler: Fail to send response")
	}
}

// ArtHandler get posts handler
func (w *WebServer) ArtHandler(ctx *fasthttp.RequestCtx) {
	var err error

	category, err := strconv.Atoi(ctx.UserValue("cat").(string))
	if err != nil {
		w.ErrorResponse(ctx, "Fail to get categories data from service")
		return
	}

	client := post.NewPostClient(w.configs.Settings().Post.Ip,
		w.configs.Settings().Post.Port)

	dataResp := DataResponse{
		Service: "post",
		Result:  true,
	}

	dataResp.Data, err = client.Posts(category)
	if err != nil {
		w.ErrorResponse(ctx, "Fail to get posts data from service")
		return
	}

	bytesOut, _ := json.Marshal(&dataResp)
	ctx.Response.Header.Add("Content-Type", "application/json")

	_, err = ctx.Write(bytesOut)
	if err != nil {
		logger.Error("PostHandler: Fail to send response")
	}
}
