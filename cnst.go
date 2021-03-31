package cnst

const (
	//Accept-Encoding Directives developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding#directives

	Brotli   = "br"
	Compress = "compress" //Lempel-Ziv-Welch (LZW) algorithm
	Deflate  = "deflate"  //zlib structure, with the deflate compression algorithm
	Gzip     = "gzip"     //Lempel-Ziv coding (LZ77), with a 32-bit CRC

	//CacheControl Directives developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control#directives

	NoCache = "no-cache"
	NoStore = "no-store"

	//X-Frame-Options Directives developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options#directives

	Deny = "deny"
)
