//go:build darwin

package autofdmax

import "github.com/boss-net/fdmax"

func init() {
	_ = fdmax.Set(fdmax.OSXMax)
}
