// Copyright (c) 2013, 2014 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import . "github.com/osmano807/joker/interfaces"

type Plugin interface {
	Name() string
	Init()
	Handle(*InputLine) *OutputLine
}
