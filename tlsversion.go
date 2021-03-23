package cert

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type tlsVers struct {
	version uint16
	name    string
}

var tlsVersions = []tlsVers{
	{version: tls.VersionSSL30, name: "SSLv3"},
	{version: tls.VersionTLS10, name: "TLS 1.0"},
	{version: tls.VersionTLS11, name: "TLS 1.1"},
	{version: tls.VersionTLS12, name: "TLS 1.2"},
	{version: tls.VersionTLS13, name: "TLS 1.3"},
}

func findMinTLSVersion(host, port string) (uint16, string, error) {

	for _, tlsVersion := range tlsVersions {
		if tryTLS(host, port, tlsVersion.version) {
			return tlsVersion.version, tlsVersion.name, nil
		}
	}
	return 0, "Unknown", fmt.Errorf("No valid TLS version found")
}

func tryTLS(host, port string, tlsVersion uint16) bool {

	d := &net.Dialer{
		Timeout: time.Duration(TimeoutSeconds) * time.Second,
	}

	connOldTLS, err := tls.DialWithDialer(d, "tcp", host+":"+port, &tls.Config{
		InsecureSkipVerify: SkipVerify,
		MaxVersion:         tlsVersion,
	})

	if err != nil {
		return false
	}

	defer connOldTLS.Close()
	return true
}
