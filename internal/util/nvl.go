// Copyright 2023 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

func Nvl(v string, fallback string) string {
	if len(v) == 0 {
		return fallback
	}
	return v
}
