// Package dnsmadeeasy implements a DNS provider for solving the DNS-01 challenge using DNS Made Easy.
package dnsmadeeasy

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/platform/config/env"
	"github.com/go-acme/lego/v4/providers/dns/dnsmadeeasy/internal"
)

// Environment variables names.
const (
	envNamespace = "DNSMADEEASY_"

	EnvAPIKey    = envNamespace + "API_KEY"
	EnvAPISecret = envNamespace + "API_SECRET"
	EnvSandbox   = envNamespace + "SANDBOX"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
	EnvHTTPTimeout        = envNamespace + "HTTP_TIMEOUT"
)

var _ challenge.ProviderTimeout = (*DNSProvider)(nil)

// Config is used to configure the creation of the DNSProvider.
type Config struct {
	BaseURL            string
	APIKey             string
	APISecret          string
	Sandbox            bool
	HTTPClient         *http.Client
	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
}

// NewDefaultConfig returns a default configuration for the DNSProvider.
func NewDefaultConfig() *Config {
	tr := &http.Transport{}

	defaultTransport, ok := http.DefaultTransport.(*http.Transport)
	if ok {
		tr = defaultTransport.Clone()
	}

	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Config{
		TTL:                env.GetOrDefaultInt(EnvTTL, dns01.DefaultTTL),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, dns01.DefaultPropagationTimeout),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, dns01.DefaultPollingInterval),
		HTTPClient: &http.Client{
			Timeout:   env.GetOrDefaultSecond(EnvHTTPTimeout, 10*time.Second),
			Transport: tr,
		},
	}
}

// DNSProvider implements the challenge.Provider interface.
type DNSProvider struct {
	config *Config
	client *internal.Client
}

// NewDNSProvider returns a DNSProvider instance configured for DNSMadeEasy DNS.
// Credentials must be passed in the environment variables:
// DNSMADEEASY_API_KEY and DNSMADEEASY_API_SECRET.
func NewDNSProvider() (*DNSProvider, error) {
	values, err := env.Get(EnvAPIKey, EnvAPISecret)
	if err != nil {
		return nil, fmt.Errorf("dnsmadeeasy: %w", err)
	}

	config := NewDefaultConfig()
	config.Sandbox = env.GetOrDefaultBool(EnvSandbox, false)
	config.APIKey = values[EnvAPIKey]
	config.APISecret = values[EnvAPISecret]

	return NewDNSProviderConfig(config)
}

// NewDNSProviderConfig return a DNSProvider instance configured for DNS Made Easy.
func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	if config == nil {
		return nil, errors.New("dnsmadeeasy: the configuration of the DNS provider is nil")
	}

	var baseURL string
	if config.Sandbox {
		baseURL = internal.DefaultSandboxBaseURL
	} else {
		if config.BaseURL == "" {
			baseURL = internal.DefaultProdBaseURL
		} else {
			baseURL = config.BaseURL
		}
	}

	client, err := internal.NewClient(config.APIKey, config.APISecret)
	if err != nil {
		return nil, fmt.Errorf("dnsmadeeasy: %w", err)
	}

	client.HTTPClient = config.HTTPClient
	client.BaseURL, err = url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &DNSProvider{
		client: client,
		config: config,
	}, nil
}

// Present creates a TXT record using the specified parameters.
func (d *DNSProvider) Present(domainName, token, keyAuth string) error {
	info := dns01.GetChallengeInfo(domainName, keyAuth)

	authZone, err := dns01.FindZoneByFqdn(info.EffectiveFQDN)
	if err != nil {
		return fmt.Errorf("dnsmadeeasy: could not find zone for domain %q: %w", domainName, err)
	}

	ctx := context.Background()

	// fetch the domain details
	domain, err := d.client.GetDomain(ctx, authZone)
	if err != nil {
		return fmt.Errorf("dnsmadeeasy: unable to get domain for zone %s: %w", authZone, err)
	}

	// create the TXT record
	name := strings.Replace(info.EffectiveFQDN, "."+authZone, "", 1)
	record := &internal.Record{Type: "TXT", Name: name, Value: info.Value, TTL: d.config.TTL}

	err = d.client.CreateRecord(ctx, domain, record)
	if err != nil {
		return fmt.Errorf("dnsmadeeasy: unable to create record for %s: %w", name, err)
	}
	return nil
}

// CleanUp removes the TXT records matching the specified parameters.
func (d *DNSProvider) CleanUp(domainName, token, keyAuth string) error {
	info := dns01.GetChallengeInfo(domainName, keyAuth)

	authZone, err := dns01.FindZoneByFqdn(info.EffectiveFQDN)
	if err != nil {
		return fmt.Errorf("dnsmadeeasy: could not find zone for domain %q: %w", domainName, err)
	}

	ctx := context.Background()

	// fetch the domain details
	domain, err := d.client.GetDomain(ctx, authZone)
	if err != nil {
		return fmt.Errorf("dnsmadeeasy: unable to get domain for zone %s: %w", authZone, err)
	}

	// find matching records
	name := strings.Replace(info.EffectiveFQDN, "."+authZone, "", 1)
	records, err := d.client.GetRecords(ctx, domain, name, "TXT")
	if err != nil {
		return fmt.Errorf("dnsmadeeasy: unable to get records for domain %s: %w", domain.Name, err)
	}

	// delete records
	var lastError error
	for _, record := range *records {
		err = d.client.DeleteRecord(ctx, record)
		if err != nil {
			lastError = fmt.Errorf("dnsmadeeasy: unable to delete record [id=%d, name=%s]: %w", record.ID, record.Name, err)
		}
	}

	return lastError
}

// Timeout returns the timeout and interval to use when checking for DNS propagation.
// Adjusting here to cope with spikes in propagation times.
func (d *DNSProvider) Timeout() (timeout, interval time.Duration) {
	return d.config.PropagationTimeout, d.config.PollingInterval
}
