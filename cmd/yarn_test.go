package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYarnDoctor_fetchURLFromRepository(t *testing.T) {
	tests := []struct {
		name     string
		gem_name string
	}{
		{
			name:     "source_code_uri exists",
			gem_name: "react",
		},
	}
	expects := []struct {
		name string
		url  string
	}{
		{
			name: "source_code_uri exists",
			url:  "git+https://github.com/facebook/react",
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := YarnDoctor{}
			r, _ := s.fetchURLFromRepository(tt.gem_name)
			expect := expects[i]
			assert.Equal(t, true, strings.HasPrefix(r, expect.url))
		})
	}
}
