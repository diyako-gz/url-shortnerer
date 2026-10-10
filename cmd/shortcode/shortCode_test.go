package shortcode

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	createdCode := CreateShortCode()

	if len(createdCode) != 6 {
		t.Errorf("expected code length 6, got %d (%q)", len(createdCode), createdCode)
	}
	for _, ch := range createdCode {
		if !strings.ContainsRune(alphabet, ch) {
			t.Errorf("invalid character %q in code %q", ch, createdCode)
			return
		}
	}

}
