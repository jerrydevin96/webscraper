package main

import (
	"testing"
)

func Test_execution(t *testing.T) {
	fetchPageDetails("http://go-colly.org/docs/examples/login/")
}
