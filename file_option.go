package gguf_parser

import "net/url"

type (
	_GGUFReadOptions struct {
		Debug             bool
		SkipLargeMetadata bool

		// Local.
		MMap bool

		// Remote.
		ProxyURL            *url.URL
		SkipProxy           bool
		SkipTLSVerification bool
		BufferSize          int
	}
	GGUFReadOption func(o *_GGUFReadOptions)
)

// UseDebug uses debug mode to read the file.
func UseDebug() GGUFReadOption {
	return func(o *_GGUFReadOptions) {
		o.Debug = true
	}
}

// SkipLargeMetadata skips reading large GGUFMetadataKV items,
// which are not necessary for most cases.
func SkipLargeMetadata() GGUFReadOption {
	return func(o *_GGUFReadOptions) {
		o.SkipLargeMetadata = true
	}
}

// UseMMap uses mmap to read the local file.
func UseMMap() GGUFReadOption {
	return func(o *_GGUFReadOptions) {
		o.MMap = true
	}
}

// UseProxy uses the given url as a proxy when reading from a remote URL.
func UseProxy(url *url.URL) GGUFReadOption {
	return func(o *_GGUFReadOptions) {
		o.ProxyURL = url
	}
}

// SkipProxy skips the proxy when reading from a remote URL.
func SkipProxy() GGUFReadOption {
	return func(o *_GGUFReadOptions) {
		o.SkipProxy = true
	}
}

// SkipTLSVerification skips the TLS verification when reading from a remote URL.
func SkipTLSVerification() GGUFReadOption {
	return func(o *_GGUFReadOptions) {
		o.SkipTLSVerification = true
	}
}

// UseBufferSize sets the buffer size when reading from a remote URL.
func UseBufferSize(size int) GGUFReadOption {
	const minSize = 32 * 1024
	if size < minSize {
		size = minSize
	}
	return func(o *_GGUFReadOptions) {
		o.BufferSize = size
	}
}
