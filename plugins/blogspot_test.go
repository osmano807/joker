// Copyright (c) 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"
import "testing"

func TestBlogspotHandle(t *testing.T) {
	var tests = pluginTestTable{
		{"http://3.bp.blogspot.com/-9bbrD6x0Ea8/UjMf2c5s_sI/AAAAAAAAM5k/Y5ZL417h6BU/s150/logo.png", "GET", NEW_STOREID, "http://cdn.bp.blogspot.com." + JOKER_SUFFIX + "/-9bbrD6x0Ea8/UjMf2c5s_sI/AAAAAAAAM5k/Y5ZL417h6BU/s150/logo.png"},
		{"http://3.bp.blogspot.com/-9bbrD6x0Ea8/UjMf2c5s_sI/AAAAAAAAM5k/Y5ZL417h6BU/s150/logo.png?s=4920&skal", "GET", NEW_STOREID, "http://cdn.bp.blogspot.com." + JOKER_SUFFIX + "/-9bbrD6x0Ea8/UjMf2c5s_sI/AAAAAAAAM5k/Y5ZL417h6BU/s150/logo.png"},
		{"http://img1.blogblog.com/img/icon18_email.gif", "GET", NEW_STOREID, "http://cdn.img.blogblog.com." + JOKER_SUFFIX + "/img/icon18_email.gif"},
	}

	var myp = &Blogspot{}

	testPluginHandle(myp, tests, t)
}
