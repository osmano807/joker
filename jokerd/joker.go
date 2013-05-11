package joker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	. "github.com/osmano807/joker/interfaces"
	"github.com/osmano807/joker/plugins"
)

var NewSquidFormat bool = false

func Main() {
	runPluginsInit()
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
		if err != nil {
			break
		}
		fields := strings.Fields(line)
		if ln := len(fields); ln == 0 {
			// TODO: check if we have any goroutine running and wait
			// to exit (or force exit?)
			log.Println("Normal exit from squid")
			return
		} else if ln != 4 && ln != 5 {
			log.Println("Linha com erro:", line)
			continue
		}

		log.Println("Fields are:", fields)
		il := parse(fields)
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

		log.Println("Exiting handling...")
	}
	if il.ChannelId == NON_CONCURRENT { // Sequential
		fn()
	} else {
		go fn()
	}
}

func printOutput(il *InputLine, ol *OutputLine) {
	if ol.ChannelId != NON_CONCURRENT {
		fmt.Printf("%v ", ol.ChannelId)
	}
	if NewSquidFormat {

		switch ol.Result {
		case NO_CHANGE:
			fmt.Println("ERR") // Squid misleading return code
		case NEW_STOREID:
			fmt.Printf("OK store-id=%v\n", ol.StoreId)
		}
	} else {
		switch ol.Result {
		case NO_CHANGE:
			fmt.Println(il.URL.String())
		case NEW_STOREID:
			fmt.Println(ol.StoreId)

		}
	}
}

func parse(s []string) (il *InputLine) {
	il = &InputLine{}
	start := 1
	if ChannelId, err := strconv.ParseUint(s[0], 10, 0); err != nil {
		il.ChannelId = NON_CONCURRENT
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

	return

}
