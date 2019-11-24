package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRepoFromUrl(t *testing.T) {
	tests := []struct {
		a string
		b string
	}{
		{a: "https://github.com/goga/elements/blob/master/spinner.js", b: "git@github.com:goga/elements.git"},
	}

	for _, test := range tests {
		result := GetRepoFromUrl(test.a)
		assert.Equal(t, result, test.b, "Should convert github file link to repo url")
	}
}
