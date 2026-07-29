//go:build integration

package integration

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	terminateSharedContainer()
	os.Exit(code)
}
