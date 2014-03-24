// Copyright (c) 2013, 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/osmano807/joker/jokerd"
	"io/ioutil"
	"log"
	"os"
)

var (
	// VERSION is initialised by the linker during compilation if the appropriate flag is specified:
	// e.g. go build -ldflags "-X main.VERSION 0.1.2-abcd" goxc.go
	// thanks to minux for this advice
	// So, goxc does this automatically during 'go build'
	VERSION     = "1.0.0"
	BUILD_DATE  = ""
	SOURCE_DATE = "2014-03-23T20:57:00-03:00"
)

func main() {
	var usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Usage = usage

	var newSquidFormat = flag.Bool("new-format", false, "Defines if joker should output the new Squid 3.HEAD format")
	var showVersion = flag.Bool("version", false, "Shows joker version")
	var enableDebug = flag.Bool("debug", false, "Enable debug messages")
	flag.Parse()

	if *showVersion {
		fmt.Fprintf(os.Stderr, " joker version: %s\n", VERSION)
		fmt.Fprintf(os.Stderr, " build date: %s\n", BUILD_DATE)
		fmt.Fprintf(os.Stderr, " source date: %s\n", SOURCE_DATE)
		os.Exit(0)
	}

	if !*enableDebug {
		log.SetOutput(ioutil.Discard)
	}

	joker.NewSquidFormat = *newSquidFormat
	joker.Main()
}
