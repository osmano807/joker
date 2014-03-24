package plugins

import . "github.com/osmano807/joker/interfaces"
import "net/url"
import "strings"

type Blogspot struct {
	name string
}

func (p *Blogspot) Name() string {
	return p.name
}

func (p *Blogspot) Init() {
	p.name = "Blogspot"
}

func (p *Blogspot) Handle(il *InputLine) (ol *OutputLine) {
	ol = &OutputLine{}

	ol.ChannelId = il.ChannelId

	if !strings.HasSuffix(il.URL.Host, "bp.blogspot.com") &&
		!(strings.HasSuffix(il.URL.Host, "blogblog.com") && strings.HasPrefix(il.URL.Host, "img")) {
		ol.Result = NO_CHANGE
		return
	}

	ol.Result = NEW_STOREID
	// Copy the URL so I don't modify the original
	var oURL url.URL
	oURL = *il.URL

	if strings.HasSuffix(il.URL.Host, "bp.blogspot.com") {
		oURL.Host = "cdn.bp.blogspot.com."
	} else {
		oURL.Host = "cdn.img.blogblog.com."
	}
	oURL.Host = oURL.Host + JOKER_SUFFIX
	oURL.RawQuery = ""

	ol.StoreId = oURL.String()

	return
}
