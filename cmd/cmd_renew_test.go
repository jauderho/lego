package cmd

import (
	"crypto/x509"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	testCases := []struct {
		desc        string
		prevDomains []string
		nextDomains []string
		expected    []string
	}{
		{
			desc:        "all empty",
			prevDomains: []string{},
			nextDomains: []string{},
			expected:    []string{},
		},
		{
			desc:        "next empty",
			prevDomains: []string{"a", "b", "c"},
			nextDomains: []string{},
			expected:    []string{"a", "b", "c"},
		},
		{
			desc:        "prev empty",
			prevDomains: []string{},
			nextDomains: []string{"a", "b", "c"},
			expected:    []string{"a", "b", "c"},
		},
		{
			desc:        "merge append",
			prevDomains: []string{"a", "b", "c"},
			nextDomains: []string{"a", "c", "d"},
			expected:    []string{"a", "b", "c", "d"},
		},
		{
			desc:        "merge same",
			prevDomains: []string{"a", "b", "c"},
			nextDomains: []string{"a", "b", "c"},
			expected:    []string{"a", "b", "c"},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			actual := merge(test.prevDomains, test.nextDomains)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_needRenewal(t *testing.T) {
	testCases := []struct {
		desc     string
		x509Cert *x509.Certificate
		days     int
		expected bool
	}{
		{
			desc: "30 days, NotAfter now",
			x509Cert: &x509.Certificate{
				NotAfter: time.Now(),
			},
			days:     30,
			expected: true,
		},
		{
			desc: "30 days, NotAfter 31 days",
			x509Cert: &x509.Certificate{
				NotAfter: time.Now().Add(31*24*time.Hour + 1*time.Second),
			},
			days:     30,
			expected: false,
		},
		{
			desc: "30 days, NotAfter 30 days",
			x509Cert: &x509.Certificate{
				NotAfter: time.Now().Add(30 * 24 * time.Hour),
			},
			days:     30,
			expected: true,
		},
		{
			desc: "0 days, NotAfter 30 days: only the day of the expiration",
			x509Cert: &x509.Certificate{
				NotAfter: time.Now().Add(30 * 24 * time.Hour),
			},
			days:     0,
			expected: false,
		},
		{
			desc: "-1 days, NotAfter 30 days: always renew",
			x509Cert: &x509.Certificate{
				NotAfter: time.Now().Add(30 * 24 * time.Hour),
			},
			days:     -1,
			expected: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			actual := needRenewal(test.x509Cert, "foo.com", test.days, false)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_needRenewalDynamic(t *testing.T) {
	testCases := []struct {
		desc                string
		now                 time.Time
		notBefore, notAfter time.Time
		expected            assert.BoolAssertionFunc
	}{
		{
			desc:      "higher than 1/3 of the certificate lifetime left (lifetime > 10 days)",
			now:       time.Date(2025, 1, 19, 1, 1, 1, 1, time.UTC),
			notBefore: time.Date(2025, 1, 1, 1, 1, 1, 1, time.UTC),
			notAfter:  time.Date(2025, 1, 30, 1, 1, 1, 1, time.UTC),
			expected:  assert.False,
		},
		{
			desc:      "lower than 1/3 of the certificate lifetime left(lifetime > 10 days)",
			now:       time.Date(2025, 1, 21, 1, 1, 1, 1, time.UTC),
			notBefore: time.Date(2025, 1, 1, 1, 1, 1, 1, time.UTC),
			notAfter:  time.Date(2025, 1, 30, 1, 1, 1, 1, time.UTC),
			expected:  assert.True,
		},
		{
			desc:      "higher than 1/2 of the certificate lifetime left (lifetime < 10 days)",
			now:       time.Date(2025, 1, 4, 1, 1, 1, 1, time.UTC),
			notBefore: time.Date(2025, 1, 1, 1, 1, 1, 1, time.UTC),
			notAfter:  time.Date(2025, 1, 10, 1, 1, 1, 1, time.UTC),
			expected:  assert.False,
		},
		{
			desc:      "lower than 1/2 of the certificate lifetime left (lifetime < 10 days)",
			now:       time.Date(2025, 1, 6, 1, 1, 1, 1, time.UTC),
			notBefore: time.Date(2025, 1, 1, 1, 1, 1, 1, time.UTC),
			notAfter:  time.Date(2025, 1, 10, 1, 1, 1, 1, time.UTC),
			expected:  assert.True,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			x509Cert := &x509.Certificate{
				NotBefore: test.notBefore,
				NotAfter:  test.notAfter,
			}

			ok := needRenewalDynamic(x509Cert, "example.com", test.now)

			test.expected(t, ok)
		})
	}
}
