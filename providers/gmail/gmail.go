package gmail

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("gmail", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for gmail.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://mail.google.com/mail/u/0/#search/%s", url.QueryEscape(q))
}
