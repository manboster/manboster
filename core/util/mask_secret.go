package util

// MaskSecret hid your secret key when you configure.
func MaskSecret(secret string) string {
	if len(secret) <= 12 {
		return "*******"
	}
	return secret[:8] + "***" + secret[len(secret)-4:]
}
