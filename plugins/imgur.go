// Copyright (c) 2013, 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "net/url"

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
	// Copy the URL so I don't modify the original
	var oURL url.URL
	oURL = *il.URL
	oURL.Path = removeExtension(oURL.Path)
	oURL.Host = oURL.Host + "." + JOKER_SUFFIX
	oURL.RawQuery = ""
	ol.StoreId = oURL.String()

	return
}
