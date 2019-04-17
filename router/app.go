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
	"github.com/LittleBuster/blog/router/server"
	"github.com/LittleBuster/blog/router/utils"
	"github.com/LittleBuster/blog/router/utils/configs"

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
		return
	}

	srvCfg := a.configs.Settings().Server

	// Starting web server
	logger.Infof("Starting web server...")
	err = a.server.Start(srvCfg.Ip, srvCfg.Port, srvCfg.Api)
	if err != nil {
		logger.Error("Fail to start web server")
	}
}

// Free free memory before exit
func (a *App) Free() {
	a.log.Free()
}
