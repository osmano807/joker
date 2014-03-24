// Copyright (c) 2013, 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "net/url"

// Plugin Imgur is used to help the caching of imgur.com images
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

	// There's also i.stack.imgur.com, but this host is a
	// custom imgur installation to stackoverflow, and has
	// some query string to resize the images, like:
	// http://i.stack.imgur.com/YYo1x.jpg?s=128&g=1&g&s=32
	// Squid should cache it as is.

	if il.URL.Host != "i.imgur.com" {
		ol.Result = NO_CHANGE
		return
	}

	ol.Result = NEW_STOREID
	// Copy the URL so I don't modify the original
	var oURL url.URL
	oURL = *il.URL
	// The extension is ignored by the imgur server, so I remove it
	// because any will work
	oURL.Path = removeExtension(oURL.Path)
	oURL.Host = oURL.Host + "." + JOKER_SUFFIX
	oURL.RawQuery = ""
	ol.StoreId = oURL.String()

	return
}
