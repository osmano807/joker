// Copyright (c) 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "net/url"
import "strings"

type Google struct {
	name string
}

func (p *Google) Name() string {
	return p.name
}

func (p *Google) Init() {
	p.name = "Google"
}

func (p *Google) Handle(il *InputLine) (ol *OutputLine) {
	ol = &OutputLine{}

	ol.ChannelId = il.ChannelId

	if !(strings.HasSuffix(il.URL.Host, "googleusercontent.com") && strings.HasPrefix(il.URL.Host, "lh")) {
		ol.Result = NO_CHANGE
		return
	}

	ol.Result = NEW_STOREID
	// Copy the URL so I don't modify the original
	var oURL url.URL
	oURL = *il.URL

	oURL.Host = "cdn.lh.googleusercontent.com." + JOKER_SUFFIX

	ol.StoreId = oURL.String()

	return
}
