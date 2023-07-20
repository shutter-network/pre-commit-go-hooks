//go:build tools
// +build tools

package tools

import (
	_ "github.com/daixiang0/gci"
	_ "mvdan.cc/gofumpt"
	_ "mvdan.cc/sh/v3/cmd/shfmt"
)
