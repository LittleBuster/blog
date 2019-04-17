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

package main

import (
	"github.com/LittleBuster/blog/post/server"
	"github.com/LittleBuster/blog/post/utils"
	"github.com/LittleBuster/blog/post/utils/configs"

	"github.com/google/logger"
)

type Application interface {
	Start()
	Free()
}

type App struct {
	log     utils.Log
	server  server.Server
	configs configs.Configurable
}

// NewApplication make new struct
func NewApplication(log utils.Log, srv server.Server,
	cfg configs.Configurable) Application {
	return &App{
		log:     log,
		server:  srv,
		configs: cfg,
	}
}

// Start application main function
func (a *App) Start() {
	// Init logger
	a.log.Init(utils.LogPath)
	logger.Info("Log was init")

	// Loading configs
	err := a.configs.LoadFromFile(utils.ConfigsPath)
	if err != nil {
		logger.Error("Fail to load configs")
		return;
	}

	// Starting web server
	logger.Infof("Starting gRPC server...")
	err = a.server.Start(a.configs.Settings().Server.Ip,
		a.configs.Settings().Server.Port)
	if err != nil {
		logger.Error("Fail to start gRPC server")
	}
}

// Free free memory before exit
func (a *App) Free() {
	a.log.Free()
}
