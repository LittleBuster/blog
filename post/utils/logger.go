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

package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/google/logger"
)

type Log interface {
	Init(path string)
	Free()
}

type Logger struct {
	logFile *os.File
	log     *logger.Logger
}

// NewLogger make new struct
func NewLogger() Log {
	return &Logger{}
}

// Init logger initialization
func (l *Logger) Init(path string) {
	var err error
	date := time.Now().Format("2006-01-02")

	// Configuring log mod
	l.logFile, err = os.OpenFile(fmt.Sprintf("%s%s.log", path, date), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Fail to open log file!")
		return
	}

	l.log = logger.Init("BlogLogger", true, true, l.logFile)
}

// Free unload logger mod
func (l *Logger) Free() {
	if l.log != nil {
		l.log.Close()
	}
	if l.logFile != nil {
		l.logFile.Close()
	}
}
