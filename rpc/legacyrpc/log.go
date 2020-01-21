// Copyright (c) 2020 The DiviProject developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package legacyrpc

import "github.com/DiviProject/divilog"

var log = divilog.Disabled

// UseLogger sets the package-wide logger.  Any calls to this function must be
// made before a server is created and used (it is not concurrent safe).
func UseLogger(logger divilog.Logger) {
	log = logger
}
