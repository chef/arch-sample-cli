package main

// these use stretchr style test assertions (better TDD than go test in hello_test.go)
// when running this, also use gocov
// go install github.com/axw/gocov/gocov@latest
// gocov test > hello_testreport.json
// gocov test | gocov report > hello_gcovreport

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestmatchInt2String calls matchInt with a non-integer string, should FAIL
func TestMatchInt2String(t *testing.T) {
	assert.Equal(t, matchInt("hello"), false, "they should be equal")
}

// TestMatchInt2StringInt calls matchInt with an integer string, should PASS
func TestMatchInt2StringInt(t *testing.T) {
	assert.Equal(t, matchInt("123"), true, "they should be equal")
}

// TestMatchInt2Empty calls matchInt with an empty string, should PASS
func TestMatchInt2Empty(t *testing.T) {
	assert.Equal(t, matchInt(""), false, "they should be equal")
}
