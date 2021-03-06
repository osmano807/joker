// Copyright (c) 2013, 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package joker

import (
	"bufio"
	"bytes"
	. "github.com/osmano807/joker/interfaces"
	"github.com/osmano807/joker/plugins"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

var NewSquidFormat bool = false
var exitWaitGroup sync.WaitGroup
var outputStream *log.Logger // Hack because fmt is not thread safe

func Main() {
	runPluginsInit()
	outputStream = log.New(os.Stdout, "", 0)
	readerLoop(os.Stdin)
}

// run some initialization code for each plugin
func runPluginsInit() {
	for _, p := range plugins.PLUGINS_LIST {
		p.Init()
	}
}

// Executes the main reading loop of squid requests
func readerLoop(input io.Reader) {
	rd := bufio.NewReader(input)
	for {
		line, err := rd.ReadString('\n')
		if err != nil && err != io.EOF { // EOF is handled on the next if
			log.Fatalln("Erro inesperado", err)
		}
		fields := strings.Fields(line)
		if ln := len(fields); ln == 0 {
			log.Println("Esperando goroutines")
			exitWaitGroup.Wait()
			log.Println("Normal exit from squid")
			return
		} else if ln < 4 {
			log.Println("Linha com erro:", line)
			continue
		}

		log.Println("Fields are:", fields)
		il, err := parse(fields)
		if err != nil {
			log.Println("Erro inesperado", err)
			continue
		}
		log.Println("Parsed:", il, il.URL.Host)
		handleInput(il)
	}
}

func handleInput(il *InputLine) {
	fn := func() {
		log.Println("Handling...")

		var ol *OutputLine

		for _, myp := range plugins.PLUGINS_LIST {
			ol = myp.Handle(il)
			if ol.Result != NO_CHANGE {
				log.Println("Match found by plugin", myp.Name())
				break
			}
		}

		printOutput(il, ol)

		if il.ChannelId != NON_CONCURRENT {
			log.Println("Done goroutine")
			exitWaitGroup.Done()
		}

		log.Println("Exiting handling...")
	}
	if il.ChannelId == NON_CONCURRENT { // Sequential
		fn()
	} else {
		log.Println("Starting goroutine")
		exitWaitGroup.Add(1)
		go fn()
	}
}

func printOutput(il *InputLine, ol *OutputLine) {
	var buffer bytes.Buffer

	if ol.ChannelId != NON_CONCURRENT {
		buffer.WriteString(strconv.Itoa(ol.ChannelId))
		buffer.WriteString(" ")
	}

	if NewSquidFormat {
		// OK	Success. A new storage ID is presented for this URL.
		// ERR	Success. No change for this URL.
		// BH	Failure. The helper encountered a problem.

		switch ol.Result {
		case NO_CHANGE:
			buffer.WriteString("ERR") // Squid misleading return code
		case NEW_STOREID:
			buffer.WriteString("OK store-id=")
			buffer.WriteString(ol.StoreId)

		}
	} else {
		switch ol.Result {
		case NO_CHANGE:
			buffer.WriteString(il.URL.String())
		case NEW_STOREID:
			buffer.WriteString(ol.StoreId)

		}
	}

	log.Println("Result:", buffer.String())

	outputStream.Println(buffer.String())
}

func parse(s []string) (il *InputLine, oerr error) {
	il = &InputLine{}
	oerr = nil
	start := 1
	if ChannelId, err := strconv.ParseUint(s[0], 10, 0); err != nil {
		il.ChannelId = NON_CONCURRENT
		start = 0
	} else {
		il.ChannelId = int(ChannelId)
	}

	if URL, err := url.Parse(s[start]); err != nil {
		log.Println("Error!")
		oerr = err
		return
	} else {
		il.URL = URL
	}

	log.Println("start:", start)

	il.Method = s[start+3]

	return

}
