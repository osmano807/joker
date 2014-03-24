// Copyright (c) 2013, 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

// The order in which the joker will evaluate the plugins
var PLUGINS_LIST []Plugin = []Plugin{
	&Imgur{},
	&Blogspot{},
	&Globo{},
	&Gravatar{}}
