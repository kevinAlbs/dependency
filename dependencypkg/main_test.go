package dependencypkg

import (
	"testing"
)

func TestGetVersion(t *testing.T) {

	if GetVersion() == "" {
		t.Error("expected non-empty GetVersion")
	}
}
