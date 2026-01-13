package cnst

// Accept-Encoding Directives https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding#directives
const (
	Brotli = "br"

	// Compress Lempel-Ziv-Welch (LZW) algorithm.
	Compress = "compress"

	// Deflate zlib structure, with the Deflate compression algorithm.
	Deflate = "deflate"

	// Gzip Lempel-Ziv coding (LZ77), with a 32-bit CRC.
	Gzip = "gzip"

	Zstandard = "zstd"
)

// CacheControl Directives https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control#directives
const (
	NoCache = "no-cache"
	NoStore = "no-store"
)

// X-Frame-Options Directives https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options#directives
const (
	Deny       = "deny"
	SameOrigin = "SAMEORIGIN"
)
