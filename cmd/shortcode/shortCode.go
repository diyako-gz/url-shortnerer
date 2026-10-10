package shortcode

import (
	"crypto/rand"
)

func CreateShortCode() string {
	code := rand.Text()
	shortedCode := code[:6]

	return shortedCode
}
