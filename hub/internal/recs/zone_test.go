package recs

import (
	"os"
	"testing"

	"life/hub/internal/store"
)

// Every fixture in this package builds its wall times in time.Local, and the
// store days them in the hub's zone (store.Day — Eastern unless main pinned
// another). Those two only agree on a machine already set to the household's
// zone. Pin it the way main does, so `make check` is green wherever the
// developer lives.
func TestMain(m *testing.M) {
	store.UseEastern()
	os.Exit(m.Run())
}
