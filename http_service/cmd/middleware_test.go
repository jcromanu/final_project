package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	testCases := []struct {
		testName      string
		bearerToken   string
		expectedError error
		expectedValue bool
	}{
		{
			testName:      "authenticate valid token",
			bearerToken:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imp1YW5jYXJsb3Mucm9tYW5AZ2xvYmFudC5jb20ifQ.SUXTBhJQOhOdPc-wTbGuQ2AUP2QXD0kyRC4j0vZ6n34",
			expectedValue: true,
			expectedError: nil,
		},
		/*{
			testName:      "authentication invalid token",
			bearerToken:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imp1YW5jYXJsb3Mucm9tYW5AZ2xvYmFudC5jb20ifQ.SUXTBhJQOhOdPc-wTbGuQ2AUP2QXD0kyRC4j0vZ6n33",
			expectedValue: false,
			expectedError: &jwt.ValidationError{},
		},*/
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			res, err := ValidateToken(tc.bearerToken)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedValue, res)
		})
	}
}
