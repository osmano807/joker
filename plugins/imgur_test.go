// Copyright (c) 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "testing"

func TestImgurHandle(t *testing.T) {
	var tests = pluginTestTable{
		{"http://i.imgur.com/cPO4ZXg.png", "GET", NEW_STOREID, "http://i.imgur.com." + JOKER_SUFFIX + "/cPO4ZXg"},
		{"http://i.imgur.com/cPO4ZXg.png?s=128&g=1&s=32", "GET", NEW_STOREID, "http://i.imgur.com." + JOKER_SUFFIX + "/cPO4ZXg"},
		{"http://i.stack.imgur.com/YYo1x.jpg?s=128&g=1&g&s=32", "GET", NO_CHANGE, ""},
	}

	var myp = &Imgur{}

	testPluginHandle(myp, tests, t)
}
