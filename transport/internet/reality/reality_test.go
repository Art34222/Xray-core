package reality_test

import (
	"testing"

	"github.com/xtls/xray-core/transport/internet/reality"
	"github.com/xtls/xray-core/transport/internet/tls"
)

func TestCheckFingerprint(t *testing.T) {
	// `random` is chosen from tls.ModernFingerprints, all of them must be accepted by REALITY servers
	for name := range tls.ModernFingerprints {
		if err := reality.CheckFingerprint(name); err != nil {
			t.Error(err)
		}
	}
	for _, name := range []string{"", "chrome", "firefox", "safari", "random", "randomized", "randomizednoalpn"} {
		if err := reality.CheckFingerprint(name); err != nil {
			t.Error(err)
		}
	}
	for _, name := range []string{"edge", "ios", "android", "360", "qq", "hellochrome_120", "hellofirefox_120", "helloedge_106"} {
		if err := reality.CheckFingerprint(name); err == nil {
			t.Error(`"fingerprint": "`, name, `" should be rejected`)
		}
	}
}
