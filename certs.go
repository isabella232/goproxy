package goproxy

import (
	"crypto/tls"
	"crypto/x509"
)

func init() {
	if goproxyCaErr != nil {
		panic("Error parsing builtin CA " + goproxyCaErr.Error())
	}
	var err error
	if GoproxyCa.Leaf, err = x509.ParseCertificate(GoproxyCa.Certificate[0]); err != nil {
		panic("Error parsing builtin CA " + err.Error())
	}
}

var tlsClientSkipVerify = &tls.Config{
	InsecureSkipVerify: true,

	// This is Go's default list of cipher suites (as of go 1.8.3),
	// with the following differences:
	//
	// - 3DES-based cipher suites have been removed. This cipher is
	//   vulnerable to the Sweet32 attack and is sometimes reported by
	//   security scanners. (This is arguably a false positive since
	//   it will never be selected: Any TLS1.2 implementation MUST
	//   include at least one cipher higher in the priority list, but
	//   there's also no reason to keep it around)
	// - AES is always prioritized over ChaCha20. Go makes this decision
	//   by default based on the presence or absence of hardware AES
	//   acceleration.
	//   TODO(bdarnell): do the same detection here. See
	//   https://github.com/golang/go/issues/21167
	//
	// Note that some TLS cipher suite guidance (such as Mozilla's[1])
	// recommend replacing the CBC_SHA suites below with CBC_SHA384 or
	// CBC_SHA256 variants. We do not do this because Go does not
	// currerntly implement the CBC_SHA384 suites, and its CBC_SHA256
	// implementation is vulnerable to the Lucky13 attack and is disabled
	// by default.[2]
	//
	// [1]: https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility
	// [2]: https://github.com/golang/go/commit/48d8edb5b21db190f717e035b4d9ab61a077f9d7
	PreferServerCipherSuites: true,
	CipherSuites: []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	},

	MinVersion: tls.VersionTLS12,
}

var defaultTLSConfig = &tls.Config{
	InsecureSkipVerify: true,

	// This is Go's default list of cipher suites (as of go 1.8.3),
	// with the following differences:
	//
	// - 3DES-based cipher suites have been removed. This cipher is
	//   vulnerable to the Sweet32 attack and is sometimes reported by
	//   security scanners. (This is arguably a false positive since
	//   it will never be selected: Any TLS1.2 implementation MUST
	//   include at least one cipher higher in the priority list, but
	//   there's also no reason to keep it around)
	// - AES is always prioritized over ChaCha20. Go makes this decision
	//   by default based on the presence or absence of hardware AES
	//   acceleration.
	//   TODO(bdarnell): do the same detection here. See
	//   https://github.com/golang/go/issues/21167
	//
	// Note that some TLS cipher suite guidance (such as Mozilla's[1])
	// recommend replacing the CBC_SHA suites below with CBC_SHA384 or
	// CBC_SHA256 variants. We do not do this because Go does not
	// currerntly implement the CBC_SHA384 suites, and its CBC_SHA256
	// implementation is vulnerable to the Lucky13 attack and is disabled
	// by default.[2]
	//
	// [1]: https://wiki.mozilla.org/Security/Server_Side_TLS#Modern_compatibility
	// [2]: https://github.com/golang/go/commit/48d8edb5b21db190f717e035b4d9ab61a077f9d7
	PreferServerCipherSuites: true,
	CipherSuites: []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	},

	MinVersion: tls.VersionTLS12,
}

var CA_CERT = []byte(`-----BEGIN CERTIFICATE-----
MIIF8jCCA9qgAwIBAgIUAp68XvvuMwaTCeQQjGxHEZhJcPgwDQYJKoZIhvcNAQEN
BQAwgZAxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJDQTEWMBQGA1UEBxMNU2FuIEZy
YW5jaXNjbzEUMBIGA1UEChMLQ29yZU9TLCBJbmMxIjAgBgNVBAsTGWdpdGh1Yi5j
b20vY29yZW9zL2dvcHJveHkxIjAgBgNVBAMTGWdpdGh1Yi5jb20vY29yZW9zL2dv
cHJveHkwHhcNMTcwMjIyMjI0NjAwWhcNMjIwMjIxMjI0NjAwWjCBkDELMAkGA1UE
BhMCVVMxCzAJBgNVBAgTAkNBMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRQwEgYD
VQQKEwtDb3JlT1MsIEluYzEiMCAGA1UECxMZZ2l0aHViLmNvbS9jb3Jlb3MvZ29w
cm94eTEiMCAGA1UEAxMZZ2l0aHViLmNvbS9jb3Jlb3MvZ29wcm94eTCCAiIwDQYJ
KoZIhvcNAQEBBQADggIPADCCAgoCggIBANghh+Y4gUYyIY1YzAgHBuLjTt13z5ED
tbjksaK8kUMofaYCnQRrepTORB3xpxn9cIXpmmPND/c3pUZz+sSbidZF+Rfkz+G4
oo5I04X7R2iFrw9jECcavr7qGqOt+vhep9iqVaioSWTKZoXY9FOxTpEUfdKyLE8K
7rd6PfdjkbZ2ZqoKdRHARBE6QlTJv+elIRM6kzshx0oSdZfaUvLdmXFqHdfKz8gb
FABaa8Aca6r76wHRnhK3+RQ2p+Wfz45/ZMxqQVIMr2mbiCCL4lSUPkUrid6L5DoA
4A4pLqY+Y0cX2qIF+lmgUwUPvunqM7dA9toHjwBgeB1yF+jVOhFOizB/dXSCsBA2
D47raP3REIxE6N05pEpGZqUPatcdyakP1Dan9aVO2W3o7P/LGPBdB5AiwVxRiVPx
dNrz+UcE3dCPBrc9bxJLjyD7PtgPTrZ+hzR/kBfBW/+jvBRKIB5NCNSuQEUN99it
P1fIencEz1ghaWhCAU5tQutnVBu8d0YlgjfnaD4vCJYEprifAiIpgZH3HaiHfG14
UXFbdU0MmuIrie2PMAGmtnjhWHG0vc8f2THRklq/CtbfNWtPzUfYv+rpaOFw7bj+
km0n3kVuGzWZhXPixe63TATR5lK9p21CZPXS+2mW/8uswBiJ4aT5lUZmCf/i1xyn
Gzgv9so56vjFAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8EBTAD
AQH/MB0GA1UdDgQWBBSYLcfHmyPv3jmdTyBs/lu5PZLxPTANBgkqhkiG9w0BAQ0F
AAOCAgEAppb+3MhGcL/IjqyRIWCc587eVMMtuHwGlczpMfXnf8TU0MPrvOdTPrqg
a+AEc1W8O6IGrowmgbZVgfr7Pw0BL4VcSdAEbP7QcbPmRAudF+/xdm5vylJmB339
4Yq5v1G2Ya5AN5PAHx3WVOA7s/caz30DGrsAoNhezlrk8WJRRxrQUauNveXyxNya
urBrqjqNVUWAlv2fRHuNXsxMvG32MWQ1BoPa9Ix+dUWYeeWkxS5E9oNo99TRRTHB
WX4SpsIynmhE+pFEzu2HVP2k5GE5+i78F9/N0cgI5QMSzgYxSebyPEflL1ziqWoq
4xio69OM0Cq+GCxvOZmvKaNlOjwJa6NWa8IyumNFclTLNzMlBaDkVE9Tb5VEBphO
ODzSxFMuAfXba4UlRjG9+BLg8ayzu2t9amB+B5/3/1+gkx1TE26jZNwleU6MnpFv
umPqPa9AW2ZPobCTGjV3MwOIKFyO/cPt4fI6BpKF7jYhUYzsht/BZqQZU/rWY6h8
GSEH8LcJk+Paaw03HzZWSbMp2oXNTz/BpQjuAdWAkeS3yhx5gvHIysWHvPRT017b
Tx+Va9+T4DElT0HtmR/w27bXNL9Urj38ETSlbYoilL7ZKltj8djXmqkI1ZF8T1zk
tj1zk0XYkestBgfrk9XYwvvL7DLyRPPZTAiVq/o5xn5zWp/y40M=
-----END CERTIFICATE-----`)

var CA_KEY = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

var GoproxyCa, goproxyCaErr = tls.X509KeyPair(CA_CERT, CA_KEY)
