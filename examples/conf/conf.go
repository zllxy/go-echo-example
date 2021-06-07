package main

import (
	"fmt"
	"go-echo-example/pkg/conf"
)

const path = "./config/config.yaml"

func main() {
	notify := make(chan bool, 1)
	v := conf.NewLoadConf(notify)
	err := v.Load(path)
	if err != nil {
		fmt.Println(err)
	}
	c := &conf.Conf{}
	c, err = v.Parse(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.Logger.LoggerFilePath)
	v.Watch()
	if <-notify {
		c, err = v.Parse(c)
		if err != nil {
			return
		}
		fmt.Println(c.Logger.LoggerFilePath)
	}
}
