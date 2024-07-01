package httputilz

import regexp "github.com/wasilibs/go-re2"

var (
	normalizeSpacesRegex = regexp.MustCompile(`\s+`)
)

func NormalizeSpaces(data string) string {
	return normalizeSpacesRegex.ReplaceAllString(data, " ")
}
