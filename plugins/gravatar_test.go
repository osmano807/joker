// Copyright (c) 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "testing"

func TestGravatarHandle(t *testing.T) {
	var tests = pluginTestTable{
		{"http://0.gravatar.com/avatar/ce1a09e75dad55baeec4dce1857b348f?s=70&d=retro&r=G", "GET", NEW_STOREID, "http://cdn.gravatar.com." + JOKER_SUFFIX + "/avatar/ce1a09e75dad55baeec4dce1857b348f?s=70&d=retro&r=G"},
	}

	var myp = &Gravatar{}

	testPluginHandle(myp, tests, t)
}
