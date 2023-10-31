package cmd

import (
	"net/http"

	parser_io "github.com/aquasecurity/go-dep-parser/pkg/io"
	"github.com/aquasecurity/go-dep-parser/pkg/php/composer"
	"github.com/aquasecurity/go-dep-parser/pkg/types"
)

type ComposerDoctor struct {
	HTTPClient http.Client
}

func NewComposerDoctor() *ComposerDoctor {
	client := &http.Client{}
	return &ComposerDoctor{HTTPClient: *client}
}

func (d *ComposerDoctor) Libraries(r parser_io.ReadSeekerAt) []types.Library {
	p := composer.NewParser()
	libs, _, _ := p.Parse(r)
	return libs
}

func (d *ComposerDoctor) SourceCodeURL(lib types.Library) (string, error) {
	packagist := Packagist{lib: lib}
	url, err := packagist.fetchURLFromRegistry(d.HTTPClient)
	return url, err
}