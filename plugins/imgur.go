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
	newpath := removeExtension(il.URL.Path)
	ol.StoreId = "http://" + "Joker/" + il.URL.Host + newpath

	return
}
