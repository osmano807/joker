// Copyright (c) 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "testing"

func TestGloboHandle(t *testing.T) {
	var tests = pluginTestTable{
		{"http://s2.glbimg.com/FGMXtx3WbbX9VHnTADSr0xM31_Y=/s.glbimg.com/og/rg/f/original/2013/11/18/vslogo60.jpg?s=543&jdjs", "GET", NEW_STOREID, "http://cdn.glbimg.com." + JOKER_SUFFIX + "/FGMXtx3WbbX9VHnTADSr0xM31_Y=/s.glbimg.com/og/rg/f/original/2013/11/18/vslogo60.jpg"},
		{"http://s01.video.glbimg.com/x360/3174912.jpg", "GET", NEW_STOREID, "http://cdn.video.glbimg.com." + JOKER_SUFFIX + "/x360/3174912.jpg"},
	}

	var myp = &Globo{}

	testPluginHandle(myp, tests, t)
}
