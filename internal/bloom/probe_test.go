package bloom

import (
	"testing"
)

// TestProbeOptimalKRounding verifies that OptimalK uses rounding (not
// truncation) so that k=round(6.64)=7 for n=1000,p=0.01 (m=9586).
func TestProbeOptimalKRounding(t *testing.T) {
	// m=9586, n=1000: k_exact = (9586/1000)*ln2 ≈ 6.643
	// Correct: round(6.643) = 7
	k := OptimalK(9586, 1000)
	if k != 7 {
		t.Errorf("OptimalK(9586,1000) = %d, want 7 (should round, not floor)", k)
	}
}
