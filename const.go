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

	//Common MIME Types developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types

	Mime3GPP   = "audio/3gpp"
	Mime3GPPV  = "video/3gpp"
	Mime3GPP2  = "audio/3gpp2"
	Mime3GPP2V = "video/3gpp2"
	Mime3GP2   = "audio/3gp2"
	Mime3GP2V  = "video/3gp2"
	Mime7Z     = "application/x-7z-compressed"
	MimeAAC    = "audio/aac"
	MimeABW    = "application/x-abiword"
	MimeARC    = "application/x-freearc"
	MimeAVI    = "video/x-msvideo"
	MimeAZW    = "application/vnd.amazon.ebook"
	MimeBIN    = "application/octet-stream"
	MimeBMP    = "image/bmp"
	MimeBZ     = "application/x-bzip"
	MimeBZ2    = "application/x-bzip2"
	MimeCSH    = "application/x-csh"
	MimeCSS    = "text/css"
	MimeCSV    = "text/csv"
	MimeDOC    = "application/msword"
	MimeDOCX   = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	MimeEOT    = "application/vnd.ms-fontobject"
	MimeEPUB   = "application/epub+zip"
	MimeFLAC   = "audio/flac"
	MimeGZ     = "application/gzip"
	MimeGIF    = "image/gif"
	MimeHTM    = "text/html"
	MimeHTML   = MimeHTM
	MimeICO    = "image/vnd.microsoft.icon"
	MimeICS    = "text/calendar"
	MimeJAR    = "application/java-archive"
	MimeJPG    = "image/jpeg"
	MimeJS     = "text/javascript"
	MimeJSON   = "application/json"
	MimeJSONLD = "application/ld+json"
	MimeMID    = "audio/midi"
	MimeMIDI   = "audio/x-midi"
	MimeMP3    = "audio/mpeg"
	MimeMP4    = "audio/mp4"
	MimeMP4V   = "video/mp4"
	MimeMPEG   = "video/mpeg"
	MimeMPKG   = "application/vnd.apple.installer+xml"
	MimeODP    = "application/vnd.oasis.opendocument.presentation"
	MimeODS    = "application/vnd.oasis.opendocument.spreadsheet"
	MimeODT    = "application/vnd.oasis.opendocument.text"
	MimeOGA    = "audio/ogg"
	MimeOGV    = "video/ogg"
	MimeOGX    = "application/ogg"
	MimeOPUS   = "audio/opus"
	MimeOTF    = "font/otf"
	MimePNG    = "image/png"
	MimePDF    = "application/pdf"
	MimePHP    = "application/x-httpd-php"
	MimePPT    = "application/vnd.ms-powerpoint"
	MimePPTX   = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	MimeRAR    = "application/vnd.rar"
	MimeRTF    = "application/rtf"
	MimeSH     = "application/x-sh"
	MimeSVG    = "image/svg+xml"
	MimeSWF    = "application/x-shockwave-flash"
	MimeTAR    = "application/x-tar"
	MimeTIF    = "image/tiff"
	MimeTS     = "video/mp2t"
	MimeTTF    = "font/ttf"
	MimeTXT    = "text/plain"
	MimeVSD    = "application/vnd.visio"
	MimeWAV    = "audio/wav"
	MimeWEBA   = "audio/webm"
	MimeWEBM   = "video/webm"
	MimeWEBP   = "image/webp"
	MimeWOFF   = "font/woff"
	MimeWOFF2  = "font/woff2"
	MimeXHTML  = "application/xhtml+xml"
	MimeXLS    = "application/vnd.ms-excel"
	MimeXLSX   = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	MimeXML    = "text/xml"        //if readable from casual users (RFC 3023, section 3)
	MimeXMLA   = "application/xml" //if not readable from casual users (RFC 3023, section 3)
	MimeXUL    = "application/vnd.mozilla.xul+xml"
	MimeZIP    = "application/zip"
)
