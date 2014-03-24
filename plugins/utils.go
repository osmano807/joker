// Copyright (c) 2013 The Joker Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import "strings"

func removeExtension(s string) string {
	if idx := strings.LastIndex(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
