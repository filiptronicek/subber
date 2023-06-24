package subber

import (
	"testing"
)

func TestPrepareClaim(t *testing.T) {
	subber := NewSubber(":", true)

	// Define a set of sub claims
	subClaims := []Claim{
		{Type: "repo", Value: "https://github.com/filiptronicek/configcat-pre-evaluate"},
		{Type: "org_id", Value: "35ec3933-83ab-4324-ad36-d4a65d678245"},
		{Type: "user_id", Value: "ceab0063-f688-4aaf-b053-8a0d158981f8"},
		{Type: "org", Value: "google"},
		{Type: "special_ability", Value: "throwing 🍎s"},
	}

	// Prepare the sub claim
	subClaim, err := subber.PrepareClaim(subClaims...)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := "repo:https%3A%2F%2Fgithub.com%2Ffiliptronicek%2Fconfigcat-pre-evaluate:org_id:35ec3933-83ab-4324-ad36-d4a65d678245:user_id:ceab0063-f688-4aaf-b053-8a0d158981f8:org:google:special_ability:throwing+%F0%9F%8D%8Es"
	if subClaim != expected {
		t.Errorf("Expected %s, but got %s", expected, subClaim)
	}
}
