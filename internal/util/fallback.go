package util

func FallbackName(nameJa, nameEn string) string {
	if nameJa != "" {
		return nameJa
	}
	return nameEn
}