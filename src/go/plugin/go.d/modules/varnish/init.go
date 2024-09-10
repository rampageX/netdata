// SPDX-License-Identifier: GPL-3.0-or-later

package varnish

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/netdata/netdata/go/plugins/pkg/executable"
)

func (v *Varnish) initVarnishstatBinary() (varnishstatBinary, error) {
	ndsudoPath := filepath.Join(executable.Directory, "ndsudo")

	if _, err := os.Stat(ndsudoPath); err != nil {
		return nil, fmt.Errorf("ndsudo executable not found: %v", err)

	}

	varnishstat := newVarnishstatBinary(ndsudoPath, v.Config, v.Logger)

	return varnishstat, nil
}