package main

import (
	"flag"
	goroutine "workshop2022/1_goroutine"
	case2 "workshop2022/9_case2"
)

func main() {
	var folder string
	flag.StringVar(&folder, "folder", "goroutine", "workshop")
	flag.Parse()

	if folder == "goroutine" {
		goroutine.Init()
	} else if folder == "channel" {

	} else if folder == "channelbuff" {

	} else if folder == "waitgroup" {

	} else if folder == "mutex" {

	} else if folder == "pubsub" {

	} else if folder == "worker" {

	} else if folder == "case1" {

	} else if folder == "case2" {
		case2.Init()
	}
}
