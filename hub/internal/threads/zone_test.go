package threads

import (
	"os"
	"testing"

	"life/hub/internal/store"
)

// A turn's prompt, the board block and the calendar rows are all worded in the
// hub's own zone (store.Day — Eastern unless main pinned another), while these
// tests build their fixtures in time.Local. Those two only agreed on a machine
// already set to the household's zone, so a `today` written as
// time.Now().Format("2006-01-02") could be a day ahead of the board's own idea
// of today. Pin the zone the way main does, so `make check` is green wherever
// the developer lives.
func TestMain(m *testing.M) {
	store.UseEastern()
	os.Exit(m.Run())
}
