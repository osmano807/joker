// Copyright (c) 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "net/url"
import "testing"

type pluginTestTable []struct {
	in     string
	method string
	rst    uint
	out    string
}

func testPluginHandle(myp Plugin, tests pluginTestTable, t *testing.T) {

	for i, tt := range tests {
		var murl *url.URL

		if murl2, err := url.Parse(tt.in); err != nil {
			t.Error(i, "Failed to parse url \"", tt.in, err)
		} else {
			murl = murl2
		}

		var il = &InputLine{ChannelId: 0, URL: murl, Method: tt.method}
		var il2 = *il

		myp.Init()

		var ol = myp.Handle(il)

		if il2 != *il {
			t.Error(i, "Plugin", myp.Name(), "rewrites InputLine", il, il2)
		}

		// Maybe I shouldn't delegate this function to the plugin
		if ol.ChannelId != il.ChannelId {
			t.Error(i, "Plugin", myp.Name(), "don't return correct channel id", il.ChannelId, ol.ChannelId)
		}

		if ol.Result != tt.rst {
			t.Error(i, "Plugin", myp.Name(), "don't return expected result", tt.rst, ol.Result)
		}

		if ol.StoreId != tt.out {
			t.Error(i, "Plugin", myp.Name(), "don't return expected store_id", tt.out, ol.StoreId)
		}
	}

}
