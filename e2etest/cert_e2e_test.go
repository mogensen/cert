package e2etest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/genkiroid/cert"
)

func TestNewCertE2E(t *testing.T) {
	cert.SkipVerify = false

	tests := []struct {
		hostport       string
		wantErrMessage string
	}{
		// Expired
		{hostport: "expired.badssl.com", wantErrMessage: "certificate has expired or is not yet valid"},

		// Wrong host
		{hostport: "wrong.host.badssl.com", wantErrMessage: "x509: certificate is valid for *.badssl.com, badssl.com, not wrong.host.badssl.com"},

		// Bad root certificates
		{hostport: "untrusted-root.badssl.com", wantErrMessage: "x509: certificate signed by unknown authority"},
		{hostport: "self-signed.badssl.com", wantErrMessage: "x509: certificate signed by unknown authority"},

		// Revoked certificate
		{hostport: "revoked.badssl.com", wantErrMessage: "Certificate have been marked as revoked"},

		// Cipher suites not allowed
		{hostport: "dh480.badssl.com", wantErrMessage: "tls: handshake failure"},
		{hostport: "dh512.badssl.com", wantErrMessage: "tls: handshake failure"},
		{hostport: "null.badssl.com", wantErrMessage: "tls: handshake failure"},
		{hostport: "rc4-md5.badssl.com", wantErrMessage: "tls: handshake failure"},
		{hostport: "rc4.badssl.com", wantErrMessage: "tls: handshake failure"},

		// Not supported
		// {hostport: "pinning-test.badssl.com", wantErrMessage: "The server uses key pinning (HPKP) but no trusted certificate chain could be constructed that matches the pinset."},
		// Certificate Transparency
		// {hostport: "no-sct.badssl.com", wantErrMessage: "The server does not send a Signed Certificate Timestamp (SCT) for this domain"},

	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s has bad certificate", tt.hostport), func(t *testing.T) {
			got := cert.NewCert(tt.hostport)
			if got.Error == "" {
				t.Errorf("NewCert() has no error, want %v", tt.wantErrMessage)
			} else {
				if !strings.Contains(got.Error, tt.wantErrMessage) {
					t.Errorf("NewCert() has %s, want %v", got.Error, tt.wantErrMessage)
				}
			}
		})
	}
}
