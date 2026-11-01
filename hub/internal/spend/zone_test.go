package spend

import (
	"os"
	"testing"

	"life/hub/internal/store"
)

// The hub's wall clock is hub.json's time_zone, and main pins it at start-up
// (store.UseZone). This package reads it as time.Local at call time
// (eastern()), so the reset and fill times it words are the household's, not
// the laptop's — which means the suite only agreed with the hub on a machine
// already set to the household's zone (Eastern when none is named). Pin it the
// way main does, so `make check` is green wherever the developer lives.
func TestMain(m *testing.M) {
	store.UseEastern()
	os.Exit(m.Run())
}
