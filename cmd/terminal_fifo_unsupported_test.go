//go:build !(darwin || dragonfly || freebsd || linux || netbsd || openbsd)

package cmd

import (
	"os"
	"testing"
)

func openTestFIFO(t *testing.T) *os.File {
	t.Helper()
	t.Skip("Unix FIFO is not available on this platform")
	return nil
}
