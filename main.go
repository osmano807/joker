package main

import "github.com/osmano807/joker/jokerd"
import "flag"

func main() {
	var newSquidFormat = flag.Bool("new-format", false, "Defines if joker should output the new Squid 3.HEAD format")
	flag.Parse()
	joker.NewSquidFormat = *newSquidFormat
	joker.Main()
}
