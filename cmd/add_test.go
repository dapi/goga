package cmd

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestReplaceGithubDirectLink(t *testing.T) {
  tests := []struct{
    a      string
    b      string
  }{
    { a: "https://github.com/dapi/elements/blob/master/spinner.js", b: "https://raw.githubusercontent.com/dapi/elements/master/spinner.js" },
  }

  for _, test := range tests {
    result:= replaceGithubDirectLink(test.a)
    assert.Equal(t, result, test.b, "Should convert direct github links")
  }
}
