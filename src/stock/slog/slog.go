/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file slog.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/18 10:20:49
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package slog

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

var LogLevel = map[string]int{
	"info":    1,
	"warning": 2,
	"failed":  3,
}

type Logger struct {
	logFile string
	level   int
}

var Slog = func() *Logger {
	return &Logger{"", 2}
}()

func New(level, file string) error {
	err := Level(level)
	if err != nil {
		return err
	}
	logFile, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	Slog.logFile = file
	logFile.Close()
	return nil
}

func Level(l string) error {
	if level, ok := LogLevel[l]; ok {
		Slog.level = level
		return nil
	}
	return errors.New("error level value[info/warning/error]")
}

func Warning(v ...interface{}) {
	if Slog.level > LogLevel["warning"] {
		return
	}
	write("WARNING", fmt.Sprint(v...))
}

func Failed(v ...interface{}) {
	if Slog.level > LogLevel["failed"] {
		return
	}
	write("FAILED", fmt.Sprint(v...))
}

func write(prefix, content string) {
	if Slog.logFile == "" {
		log.Println("Not file", prefix, content)
		return
	}

	logFile, err := os.OpenFile(Slog.logFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("log file open failed", err)
		return
	}

	logger := log.New(logFile, "", 0)
	logger.Println(time.Now().Format("2006/01/02 15:04:05") + " [" + prefix + "] " + content)
}
