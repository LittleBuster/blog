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
	"fmt"
	"net"

	"github.com/LittleBuster/blog/post/db/connector"
	pb "github.com/LittleBuster/blog/post/server/proto"
	"github.com/LittleBuster/blog/post/utils/configs"

	"github.com/google/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server interface {
	Start(ip string, port int) error
}

type RpcServer struct {
	configs configs.Configurable
}

// NewWebServer make new struct
func NewRpcServer(cfg configs.Configurable) Server {
	return &RpcServer{
		configs: cfg,
	}
}

// GetCategories get all categories
func (r *RpcServer) GetCategories(context.Context, *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	resp := &pb.CategoryResponse{}

	dbCfg := r.configs.Settings().Db
	db := connector.NewDatabase(dbCfg.Ip, dbCfg.User, dbCfg.Passwd, dbCfg.Base)

	err := db.Connect()
	if err != nil {
		logger.Error("Fail to connect to database database")
		resp.Result = false
		return nil, err
	}
	defer db.Close()

	cats, err := db.Categories()
	if err != nil {
		logger.Error("Fail to get categories from database")
		resp.Result = false
		return nil, err
	}

	for _, cat := range cats {
		category := &pb.Category{
			Id:    int32(cat.Id),
			Title: cat.Title,
		}
		resp.Categories = append(resp.Categories, category)
	}

	if len(cats) == 0 {
		resp.Result = false
	} else {
		resp.Result = true
	}

	return resp, nil
}

// GetPosts get all posts data
func (r *RpcServer) GetPosts(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {
	resp := &pb.PostResponse{}

	dbCfg := r.configs.Settings().Db
	db := connector.NewDatabase(dbCfg.Ip, dbCfg.User, dbCfg.Passwd, dbCfg.Base)

	err := db.Connect()
	if err != nil {
		logger.Error("Fail to connect to database database")
		resp.Result = false
		return nil, err
	}
	defer db.Close()

	posts, err := db.Posts(int(req.Category))
	if err != nil {
		logger.Error("Fail to connect to database database")
		resp.Result = false
		return nil, err
	}

	for _, ps := range posts {
		post := &pb.Post{
			Id:       int32(ps.Id),
			Title:    ps.Title,
			Text:     ps.Text,
			Category: int32(ps.Category),
			Created:  ps.Created,
		}
		resp.Posts = append(resp.Posts, post)
	}

	if len(posts) == 0 {
		resp.Result = false
	} else {
		resp.Result = true
	}

	return resp, nil
}

// Start start gRPC server
func (r *RpcServer) Start(ip string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterPostsServer(s, r)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
