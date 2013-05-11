package interfaces

import "net/url"

type InputLine struct {
	ChannelId int
	URL       *url.URL
	Method    string
}

const (
	NON_CONCURRENT = -1
	NEW_STOREID    = iota
	NO_CHANGE
)

type OutputLine struct {
	ChannelId int
	Result    uint
	StoreId   string
}

const JOKER_PREFIX = "Joker.squid.internal"
