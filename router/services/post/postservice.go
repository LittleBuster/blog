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

package post

import (
	"errors"
	"fmt"
	"time"

	pb "github.com/LittleBuster/blog/router/services/post/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Category struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Category int    `json:"category"`
	Created  string `json:"created"`
}

type PostService interface {
	Categories() ([]Category, error)
	Posts(category int) ([]Post, error)
}

type PostClient struct {
	ip   string
	port int
}

// NewPostService make new struct
func NewPostClient(ip string, port int) PostService {
	return &PostClient{
		ip:   ip,
		port: port,
	}
}

// Categories get all categories
func (p *PostClient) Categories() ([]Category, error) {
	var cats []Category

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", p.ip, p.port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewPostsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.GetCategories(ctx, &pb.CategoryRequest{})
	if err != nil {
		return nil, err
	}

	if resp.Result != true {
		return nil, errors.New("Fail to get data from database")
	}

	for _, category := range resp.Categories {
		cat := Category{
			Id:    int(category.Id),
			Title: category.Title,
		}
		cats = append(cats, cat)
	}

	return cats, nil
}

// Posts get all posts
func (p *PostClient) Posts(category int) ([]Post, error) {
	var posts []Post

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", p.ip, p.port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewPostsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.GetPosts(ctx, &pb.PostRequest{Category: int32(category)})
	if err != nil {
		return nil, err
	}

	if resp.Result != true {
		return nil, errors.New("Fail to get data from database")
	}

	for _, post := range resp.Posts {
		ps := Post{
			Id:       int(post.Id),
			Title:    post.Title,
			Text:     post.Text,
			Category: int(post.Category),
			Created:  post.Created,
		}
		posts = append(posts, ps)
	}

	return posts, nil
}
