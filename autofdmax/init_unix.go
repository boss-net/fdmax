//go:build linux || openbsd || netbsd

package autofdmax

import (
	"github.com/boss-net/fdmax"
)

func init() {
	_ = fdmax.Set(fdmax.UnixMax)
}
