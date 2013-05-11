package plugins

import . "github.com/osmano807/joker/interfaces"

type Imgur struct {
	name string
}

func (p *Imgur) Name() string {
	return p.name
}

func (p *Imgur) Init() {
	p.name = "Imgur"
}

func (p *Imgur) Handle(il *InputLine) (ol *OutputLine) {
	ol = &OutputLine{}

	ol.ChannelId = il.ChannelId

	if il.URL.Host != "i.imgur.com" {
		ol.Result = NO_CHANGE
		return
	}

	ol.Result = NEW_STOREID
	il.URL.Path = removeExtension(il.URL.Path)
	il.URL.Host = JOKER_PREFIX + "/" + il.URL.Host
	il.URL.RawQuery = ""
	ol.StoreId = il.URL.String()

	return
}
