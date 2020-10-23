package main

import (
	"flag"
	"fmt"
	"github.com/ammartynenko/plugwww/packages"
	"os"
)

const (
	info = "usage: %s\n"
)

func main() {

	var (
		configPath string
	)
	flag.StringVar(&configPath, "config", "", "файл конфигурации с полным путем к нему")
	flag.Parse()
	flag.Usage = func() {
		fmt.Printf(info, os.Args[0])
		flag.PrintDefaults()
	}
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	s := packages.NewServer(configPath)
	s.Mux.Get("/", s.HandlerRoot)
	s.Run()

}
