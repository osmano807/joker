// Copyright (c) 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "testing"

func TestGoogleHandle(t *testing.T) {
	var tests = pluginTestTable{
		{"http://lh4.googleusercontent.com/-H7wrBlfJhMA/AAAAAAAAAAI/AAAAAAAAASQ/0HZnUHdTPgg/s512-c/photo.jpg?s=543&jdjs", "GET", NEW_STOREID, "http://cdn.lh.googleusercontent.com." + JOKER_SUFFIX + "/-H7wrBlfJhMA/AAAAAAAAAAI/AAAAAAAAASQ/0HZnUHdTPgg/s512-c/photo.jpg?s=543&jdjs"},
	}

	var myp = &Google{}

	testPluginHandle(myp, tests, t)
}
