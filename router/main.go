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
	"fmt"

	"github.com/LittleBuster/blog/router/server"
	"github.com/LittleBuster/blog/router/utils"
	"github.com/LittleBuster/blog/router/utils/configs"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	container.Provide(utils.NewLogger)
	container.Provide(server.NewWebServer)
	container.Provide(configs.NewConfigs)
	container.Provide(NewApplication)

	err := container.Invoke(func(app Application) {
		app.Start()
		app.Free()
	})

	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
	}
}
