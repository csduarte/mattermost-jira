package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
)

func initLog(path string, verbose bool) *logrus.Logger {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Failed to open/create log -", err.Error())
		os.Exit(1)
	}

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	log := logrus.New()
	fileHook := lfshook.NewHook(lfshook.PathMap{
		logrus.InfoLevel: absPath,
	})
	fileHook.SetFormatter(&logrus.JSONFormatter{})
	log.Hooks.Add(fileHook)
	log.Infof("Logs: %s", absPath)
	return log
}
