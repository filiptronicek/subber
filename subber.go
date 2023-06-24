package subber

import (
	"net/url"
	"strings"
)

// Claim represents a single field inside of the claim
type Claim struct {
	Key   string
	Value string
}

type Subber struct {
	// Delimiter to use between the type and value of the sub claim
	Delimiter string
	// Encode the value of the sub claim
	Encode bool
}

// NewSubber creates a new Subber object
func NewSubber(delimiter string, encode bool) *Subber {
	return &Subber{Delimiter: delimiter, Encode: encode}
}

// PrepareClaim generates the sub claim using the provided sub claims
func (s *Subber) PrepareClaim(claims ...Claim) (string, error) {
	var parts []string

	for _, claim := range claims {

		encodedValue := claim.Value
		if s.Encode {
			encodedValue = url.QueryEscape(claim.Value)
		}

		parts = append(parts, claim.Key+s.Delimiter+encodedValue)
	}

	return strings.Join(parts, s.Delimiter), nil
}
