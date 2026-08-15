package reverseproxy

import (
	"net/http"
	"net/url"
)

// Ensure the request is properly prepared for the reverse proxy
// to prevent connection pool leakage.
func prepareForwardAuthRequest(req *http.Request, upstream *url.URL) {
	req.URL.Scheme = upstream.Scheme
	req.URL.Host = upstream.Host
	req.URL.Path = upstream.Path
	req.URL.RawQuery = upstream.RawQuery
	// Explicitly set the Host header to match the upstream host
	// This is critical for Go's http.Transport to select the correct connection pool.
	req.Host = upstream.Host
}