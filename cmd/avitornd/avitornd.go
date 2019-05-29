/*
Copyright 2019 Vladislav Dmitriyev.
*/

// Package avitornd provides command-line interface to start REST API
// server for generating random values by specified type and length params.
package main

import (
	"flag"
	"os"

	"github.com/eeonevision/avito-pro-test/internal/api"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
}

func main() {
	var host string
	var port string

	flag.StringVar(&host, "host", "0.0.0.0", "the host address for starting api")
	flag.StringVar(&port, "port", "8888", "the port number for starting api")
	flag.Parse()

	api.Serve(host+":"+port, logrus.New())
}
