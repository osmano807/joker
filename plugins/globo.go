// Copyright (c) 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "net/url"
import "strings"

type Globo struct {
	name string
}

func (p *Globo) Name() string {
	return p.name
}

func (p *Globo) Init() {
	p.name = "Globo"
}

func (p *Globo) Handle(il *InputLine) (ol *OutputLine) {
	ol = &OutputLine{}

	ol.ChannelId = il.ChannelId

	if !strings.HasSuffix(il.URL.Host, "glbimg.com") {
		ol.Result = NO_CHANGE
		return
	}

	ol.Result = NEW_STOREID
	// Copy the URL so I don't modify the original
	var oURL url.URL
	oURL = *il.URL

	if strings.HasSuffix(il.URL.Host, "video.glbimg.com") {
		oURL.Host = "cdn.video.glbimg.com."
	} else {
		oURL.Host = "cdn.glbimg.com."
	}
	oURL.Host = oURL.Host + JOKER_SUFFIX
	oURL.RawQuery = ""

	ol.StoreId = oURL.String()

	return
}
