package joker

import (
	"bufio"
	"fmt"
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
	var prefix string = ""
	if ol.ChannelId != NON_CONCURRENT {
		prefix = fmt.Sprintf("%d ", ol.ChannelId)
	}

	if NewSquidFormat {
		switch ol.Result {
		case NO_CHANGE:
			outputStream.Println(prefix + "ERR") // Squid misleading return code
		case NEW_STOREID:
			outputStream.Printf(prefix+"OK store-id=%v\n", ol.StoreId)
		}
	} else {
		switch ol.Result {
		case NO_CHANGE:
			outputStream.Println(prefix + il.URL.String())
		case NEW_STOREID:
			outputStream.Println(prefix + ol.StoreId)

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
