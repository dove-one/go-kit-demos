package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"

	conf "helloKit/conf"
	grpc "helloKit/server/grpc"
	http "helloKit/server/http"

	_ "helloKit/router/grpc"
	_ "helloKit/router/http"
)

func init() {
	// flags
	flag.StringVar(&conf.HttpAddr, "http-addr", conf.GetEnv("HttpAddr", "0.0.0.0:8088"), "http服务地址")
	flag.StringVar(&conf.GrpcAddr, "grpc-addr", conf.GetEnv("GrpcAddr", "0.0.0.0:5000"), "grpc服务地址")

	// log
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2005-01-02 15:04:05",
	})

	log.WithFields(log.Fields{
		"http-addr": conf.HttpAddr,
		"grpc-addr": conf.GrpcAddr,
	}).Info("run flags:")
}

func main() {
	flag.Parse()

	errc := make(chan error)

	// http server
	{
		log.WithField("http-addr", conf.HttpAddr).Info("http server is running...")
		go http.Run(conf.HttpAddr, errc)
	}

	// grpc server
	{
		log.WithField("grpc-addr", conf.GrpcAddr).Info("grpc server is running...")
		go grpc.Run(conf.GrpcAddr, errc)
	}

	log.WithField("error", <-errc).Info("Exit")
}
