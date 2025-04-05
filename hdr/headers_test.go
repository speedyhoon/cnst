package hdr_test

import (
	"github.com/speedyhoon/cnst/hdr"
	"testing"
)

func TestUnique(t *testing.T) {
	// Using constants as the keys to prevent duplicate constants from being defined.
	names := map[string]string{
		hdr.Accept:                         "accept",
		hdr.AcceptCH:                       "accept-ch",
		hdr.AcceptEncoding:                 "accept-encoding",
		hdr.AcceptLanguage:                 "accept-language",
		hdr.AcceptPatch:                    "accept-patch",
		hdr.AcceptPost:                     "accept-post",
		hdr.AcceptRanges:                   "accept-ranges",
		hdr.AccessControlAllowCredentials:  "access-control-allow-credentials",
		hdr.AccessControlAllowHeaders:      "access-control-allow-headers",
		hdr.AccessControlAllowMethods:      "access-control-allow-methods",
		hdr.AccessControlAllowOrigin:       "access-control-allow-origin",
		hdr.AccessControlExposeHeaders:     "access-control-expose-headers",
		hdr.AccessControlMaxAge:            "access-control-max-age",
		hdr.AccessControlRequestHeaders:    "access-control-request-headers",
		hdr.AccessControlRequestMethod:     "access-control-request-method",
		hdr.Age:                            "age",
		hdr.Allow:                          "allow",
		hdr.AltSvc:                         "alt-svc",
		hdr.AltUsed:                        "alt-used",
		hdr.Authorization:                  "authorization",
		hdr.CacheControl:                   "cache-control",
		hdr.ClearSiteData:                  "clear-site-data",
		hdr.Connection:                     "connection",
		hdr.ContentDigest:                  "content-digest",
		hdr.ContentDisposition:             "content-disposition",
		hdr.ContentEncoding:                "content-encoding",
		hdr.ContentLanguage:                "content-language",
		hdr.ContentLength:                  "content-length",
		hdr.ContentLocation:                "content-location",
		hdr.ContentRange:                   "content-range",
		hdr.ContentType:                    "content-type",
		hdr.Cookie:                         "cookie",
		hdr.CrossOriginEmbedderPolicy:      "cross-origin-embedder-policy",
		hdr.CrossOriginOpenerPolicy:        "cross-origin-opener-policy",
		hdr.CrossOriginResourcePolicy:      "cross-origin-resource-policy",
		hdr.CSP:                            "content-security-policy",
		hdr.CSPReportOnly:                  "content-security-policy-report-only",
		hdr.Date:                           "date",
		hdr.DeviceMemory:                   "device-memory",
		hdr.ETag:                           "etag",
		hdr.Expect:                         "expect",
		hdr.Expires:                        "expires",
		hdr.Forwarded:                      "forwarded",
		hdr.From:                           "from",
		hdr.Host:                           "host",
		hdr.IfMatch:                        "if-match",
		hdr.IfModifiedSince:                "if-modified-since",
		hdr.IfNoneMatch:                    "if-none-match",
		hdr.IfRange:                        "if-range",
		hdr.IfUnmodifiedSince:              "if-unmodified-since",
		hdr.KeepAlive:                      "keep-alive",
		hdr.LastModified:                   "last-modified",
		hdr.Link:                           "link",
		hdr.Location:                       "location",
		hdr.MaxForwards:                    "max-forwards",
		hdr.Origin:                         "origin",
		hdr.OriginAgentCluster:             "origin-agent-cluster",
		hdr.Prefer:                         "prefer",
		hdr.PreferenceApplied:              "preference-applied",
		hdr.Priority:                       "priority",
		hdr.ProxyAuthenticate:              "proxy-authenticate",
		hdr.ProxyAuthorization:             "proxy-authorization",
		hdr.Range:                          "range",
		hdr.Referer:                        "referer",
		hdr.ReferrerPolicy:                 "referrer-policy",
		hdr.Refresh:                        "refresh",
		hdr.ReprDigest:                     "repr-digest",
		hdr.RetryAfter:                     "retry-after",
		hdr.SecFetchDest:                   "sec-fetch-dest",
		hdr.SecFetchMode:                   "sec-fetch-mode",
		hdr.SecFetchSite:                   "sec-fetch-site",
		hdr.SecFetchUser:                   "sec-fetch-user",
		hdr.SecPurpose:                     "sec-purpose",
		hdr.SecWebSocketAccept:             "sec-websocket-accept",
		hdr.SecWebSocketExtensions:         "sec-websocket-extensions",
		hdr.SecWebSocketKey:                "sec-websocket-key",
		hdr.SecWebSocketProtocol:           "sec-websocket-protocol",
		hdr.SecWebSocketVersion:            "sec-websocket-version",
		hdr.Server:                         "server",
		hdr.ServerTiming:                   "server-timing",
		hdr.ServiceWorker:                  "service-worker",
		hdr.ServiceWorkerAllowed:           "service-worker-allowed",
		hdr.ServiceWorkerNavigationPreload: "service-worker-navigation-preload",
		hdr.SetCookie:                      "set-cookie",
		hdr.SourceMap:                      "sourcemap",
		hdr.StrictTransportSecurity:        "strict-transport-security",
		hdr.TE:                             "te",
		hdr.TimingAllowOrigin:              "timing-allow-origin",
		hdr.Trailer:                        "trailer",
		hdr.TransferEncoding:               "transfer-encoding",
		hdr.Upgrade:                        "upgrade",
		hdr.UpgradeInsecureRequests:        "upgrade-insecure-requests",
		hdr.UserAgent:                      "user-agent",
		hdr.Vary:                           "vary",
		hdr.Via:                            "via",
		hdr.WantContentDigest:              "want-content-digest",
		hdr.WantReprDigest:                 "want-repr-digest",
		hdr.WWWAuthenticate:                "www-authenticate",
		hdr.XContentTypeOptions:            "x-content-type-options",
		hdr.XFrameOptions:                  "x-frame-options",
	}

	t.Run("expected quantity", func(t *testing.T) {
		const expectedQty = 100
		if qty := len(names); qty != expectedQty {
			t.Errorf("have %d headers, expected %d\n", qty, expectedQty)
		}
	})

	for value, expected := range names {
		t.Run(expected, func(t *testing.T) {
			if value != expected {
				t.Errorf("expected header `%s`, got: `%s`", expected, value)
			}
		})
	}
}
