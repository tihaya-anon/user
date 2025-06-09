package test

import (
	"MVC_DI/security"
	"testing"
)

func Test_JWT(t *testing.T) {
	// Case 1: Generate JWT
	// Given
	type Claims struct {
		Name string
		Age  int
	}
	// When
	token, err := security.GenerateJWT(Claims{Name: "John", Age: 30})
	// Then
	if err != nil {
		t.Errorf("case `Generate JWT` failed: %v", err)
		return
	}
	// Case 2: Parse JWT Success
	// Given
	// token above
	// When
	_, err = security.ParseJWT[Claims](token)
	// Then
	if err != nil {
		t.Errorf("case `Parse JWT Success` failed: %v", err)
		return
	}
	// Case 3: Parse JWT Failed
	// Given
	token = "abcd" + token
	// When
	_, err = security.ParseJWT[Claims](token)
	// Then
	if err == nil {
		t.Errorf("case `Parse JWT Failed` failed: %v", err)
		return
	}
}
