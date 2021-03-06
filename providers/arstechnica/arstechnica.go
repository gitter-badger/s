package arstechnica

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("arstechnica", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for arstechnica.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://arstechnica.com/search/?ie=UTF-8&q=%s", url.QueryEscape(q))
}
