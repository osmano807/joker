package plugins

import . "github.com/osmano807/joker/interfaces"
import "net/url"
import "strings"

type Gravatar struct {
	name string
}

func (p *Gravatar) Name() string {
	return p.name
}

func (p *Gravatar) Init() {
	p.name = "Gravatar"
}

func (p *Gravatar) Handle(il *InputLine) (ol *OutputLine) {
	ol = &OutputLine{}

	ol.ChannelId = il.ChannelId

	if !(strings.HasSuffix(il.URL.Host, "gravatar.com") && strings.HasPrefix(il.URL.Path, "/avatar/")) {
		ol.Result = NO_CHANGE
		return
	}

	ol.Result = NEW_STOREID
	// Copy the URL so I don't modify the original
	var oURL url.URL
	oURL = *il.URL

	oURL.Host = "cdn.gravatar.com." + JOKER_SUFFIX

	// For now, preserve the query string
	// it's needed to server different images based on the
	// image rate (G, PG, etc)
	// BUT, I need to pay attention to a "d" field
	// that serves probably to track the request url
	// maybe that can be removed

	ol.StoreId = oURL.String()

	return
}
