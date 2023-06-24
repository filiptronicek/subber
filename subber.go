package subber

import (
	"net/url"
	"strings"
)

type SubClaimClaim struct {
	Type  string
	Value string
}

type Subber struct {
	// Delimiter to use between the type and value of the sub claim
	Delimiter string
	// Encode the value of the sub claim
	Encode bool
}

func NewSubber(delimiter string, encode bool) *Subber {
	return &Subber{Delimiter: delimiter, Encode: encode}
}

// PrepareSubClaim generates the sub claim using the provided sub claims
func (s *Subber) PrepareSubClaim(subClaims ...SubClaimClaim) (string, error) {
	var parts []string

	for _, claim := range subClaims {

		encodedValue := claim.Value
		if s.Encode {
			encodedValue = url.QueryEscape(claim.Value)
		}

		parts = append(parts, claim.Type+s.Delimiter+encodedValue)
	}

	return strings.Join(parts, s.Delimiter), nil
}
