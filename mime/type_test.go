package mime_test

import (
	"github.com/speedyhoon/cnst/mime"
	"testing"
)

func TestUnique(t *testing.T) {
	// Using constants as the keys to prevent duplicate constants from being defined.
	names := map[string]string{
		mime.A3GP:  "audio/3gpp",
		mime.V3GP:  "video/3gpp",
		mime.A3GP2: "audio/3gpp2",
		mime.V3GP2: "video/3gpp2",

		mime.AAC:    "audio/aac",
		mime.ABW:    "application/x-abiword",
		mime.ARC:    "application/x-freearc",
		mime.AVI:    "video/x-msvideo",
		mime.AZW:    "application/vnd.amazon.ebook",
		mime.BIN:    "application/octet-stream",
		mime.BMP:    "image/bmp",
		mime.BZ:     "application/x-bzip",
		mime.BZ2:    "application/x-bzip2",
		mime.CSH:    "application/x-csh",
		mime.CSS:    "text/css; charset=utf-8",
		mime.CSV:    "text/csv",
		mime.DOC:    "application/msword",
		mime.DOCX:   "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		mime.EOT:    "application/vnd.ms-fontobject",
		mime.EPUB:   "application/epub+zip",
		mime.FLAC:   "audio/flac",
		mime.GZ:     "application/gzip",
		mime.GIF:    "image/gif",
		mime.HTML:   "text/html; charset=utf-8",
		mime.ICO:    "image/vnd.microsoft.icon",
		mime.ICS:    "text/calendar",
		mime.JAR:    "application/java-archive",
		mime.JPG:    "image/jpeg",
		mime.JS:     "text/javascript; charset=utf-8",
		mime.JSON:   "application/json",
		mime.JSONLD: "application/ld+json",
		mime.MD:     "text/markdown",
		mime.MID:    "audio/midi",
		mime.MIDI:   "audio/x-midi",
		mime.MP3:    "audio/mpeg",
		mime.MP4a:   "audio/mp4",
		mime.MP4v:   "video/mp4",
		mime.MPEG:   "video/mpeg",
		mime.MPKG:   "application/vnd.apple.installer+xml",
		mime.ODP:    "application/vnd.oasis.opendocument.presentation",
		mime.ODS:    "application/vnd.oasis.opendocument.spreadsheet",
		mime.ODT:    "application/vnd.oasis.opendocument.text",
		mime.OGA:    "audio/ogg",
		mime.OGV:    "video/ogg",
		mime.OGX:    "application/ogg",
		mime.OPUS:   "audio/opus",
		mime.OTF:    "font/otf",
		mime.PNG:    "image/png",
		mime.PDF:    "application/pdf",
		mime.PHP:    "application/x-httpd-php",
		mime.PPT:    "application/vnd.ms-powerpoint",
		mime.PPTX:   "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		mime.RAR:    "application/vnd.rar",
		mime.RTF:    "application/rtf",
		mime.SH:     "application/x-sh",
		mime.SVG:    "image/svg+xml",
		mime.SWF:    "application/x-shockwave-flash",
		mime.TAR:    "application/x-tar",
		mime.TIF:    "image/tiff",
		mime.TS:     "video/mp2t",
		mime.TTF:    "font/ttf",
		mime.TXT:    "text/plain",
		mime.VSD:    "application/vnd.visio",
		mime.WAV:    "audio/wav",
		mime.WEBMa:  "audio/webm",
		mime.WEBMv:  "video/webm",
		mime.WEBP:   "image/webp",
		mime.WOFF:   "font/woff",
		mime.WOFF2:  "font/woff2",
		mime.XHTML:  "application/xhtml+xml",
		mime.XLS:    "application/vnd.ms-excel",
		mime.XLSX:   "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		mime.XML:    "text/xml",
		mime.XMLA:   "application/xml",
		mime.XUL:    "application/vnd.mozilla.xul+xml",
		mime.ZIP:    "application/zip",
		mime.ZIP7:   "application/x-7z-compressed",
	}

	t.Run("expected quantity", func(t *testing.T) {
		const expectedQty = 77
		if qty := len(names); qty != expectedQty {
			t.Errorf("have %d MIME types, expected %d\n", qty, expectedQty)
		}
	})

	for value, expected := range names {
		t.Run(expected, func(t *testing.T) {
			if value != expected {
				t.Errorf("expected mime type `%s`, got: `%s`", expected, value)
			}
		})
	}
}
