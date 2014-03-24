// Copyright (c) 2013, 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// Per http://wiki.squid-cache.org/Features/StoreID
// the more common use id just squid.internal
const JOKER_SUFFIX = "squid.internal"
