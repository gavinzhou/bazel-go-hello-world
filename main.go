package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
)

func main() {
	fmt.Println("Hello, world!")
	flag.Parse()
	defer glog.Flush()
	glog.Info("Hello")
}
