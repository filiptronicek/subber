# `subber`
A non-standard way of inlining abitrary fields into the OIDC JWT `sub` claim.

## Why?

When you want to use OIDC to control access to resources with services like AWS IAM, you need to use the `sub` claim in the JWT to identify the user. This is some primary user identification (like an email address), but sometimes you want to store more information in this single claim. This library allows you to easily do that.

## Usage

```go
package main

import (
	"log"
	"fmt"

	"github.com/yourname/subber"
)

func main() {
	// Import the subber package
	// Replace "github.com/yourname/subber" with the actual path of your subber package

	// Create a new Subber with your desired delimiter and specify if the claim value should be URL-encoded
	subber := subber.NewSubber("-", true)

	// Create a set of ClaimClaim objects
	claims := []subber.ClaimClaim{
		{Type: "type1", Value: "value1"},
		{Type: "type2", Value: "value2 with/special&characters"},
	}

	// Use the PrepareClaim method of the Subber object to generate the claim string
	claim, err := subber.PrepareClaim(claims...)
	if err != nil {
		log.Fatalf("Error preparing claim: %v", err)
	}

	fmt.Println(claim)
	// If you've specified the delimiter as "-" and encoding as true, this will output:
	// type1-value1-type2-value2+with%2Fspecial%26characters
}
```

## Testing

To run tests, navigate to the package directory and run the following command:

```bash
go test
```
