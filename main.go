package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"io"
	"strconv"
	"strings"
)

type InputLine struct {
	ChannelId int
	URL       *url.URL
	Method    string
}

const (
	NEW_STOREID = iota
	NO_CHANGE
)

type OutputLine struct {
	ChannelId int
	result    uint
	store_id  string
}

func main() {
	readerLoop(os.Stdin)
}

func readerLoop(input io.Reader) {
	rd := bufio.NewReader(input)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Fields(line)
		if ln := len(fields); ln == 0 {
			log.Println("Normal exit from squid")
			return
		} else if ln != 4 && ln != 5 {
			log.Println("Linha com erro:", line)
			continue
		}

		il := new(InputLine)
		log.Println("Fields are:", fields)
		il.Parse(fields)
		log.Println("Parsed:", il, il.URL.Host)
		handleInput(il)
	}
}

func handleInput(il *InputLine) {
	fn := func () {
		log.Println("Handling...")

		ol := new(OutputLine)
		ol.ChannelId = il.ChannelId
		ol.result = NO_CHANGE

		printOutput(ol)

		log.Println("Exiting handling...")
	}
	if il.ChannelId == -1 { // Sequential
		fn()
	} else {
		go fn()
	}
}

func printOutput(ol *OutputLine) {
	if(ol.ChannelId != -1) {
		fmt.Printf("%v ", ol.ChannelId)
	}
	switch ol.result {
	case NO_CHANGE:
		fmt.Println("ERR") // Squid misleading return code
	case NEW_STOREID:
		fmt.Printf("OK store-id=%v", ol.store_id)
	}
}

func (il *InputLine) Parse(s []string) {
	start := 1
	if ChannelId, err := strconv.ParseUint(s[0], 10, 0); err != nil {
		il.ChannelId = -1
		start = 0
	} else {
		il.ChannelId = int(ChannelId)
	}

	if URL, err := url.Parse(s[start]); err != nil {
		log.Println("Error!")
		return
	} else {
		il.URL = URL
	}

	log.Println("start:", start)

	il.Method = s[start+3]

}
