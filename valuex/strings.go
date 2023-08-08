package valuex

import (
	"strings"
	"unicode"
)

// HasPrefixIn
// Determines whether str has one of the prefix in the list.
func HasPrefixIn(str string, prefixes []string) bool {
	for _, v := range prefixes {
		if strings.HasPrefix(str, v) {
			return true
		}
	}

	return false
}

// CamelToUnderline
// Converts src string from Camel-Format (e.g SameStringInCamelFormat) to
// Underline-Format (e.g some_string_in_underline_format).
func CamelToUnderline(src string) string {
	sb := strings.Builder{}

	for i, r := range src {
		if unicode.IsUpper(r) {
			if i != 0 {
				sb.WriteByte('_')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// FirstToLower
// Converts the first rune to lower-case.
func FirstToLower(src string) string {
	if len(src) == 0 {
		return src
	}

	return strings.ToLower(src[:1]) + src[1:]
}

// FirstToUpper
// Converts the first rune to upper-case.
func FirstToUpper(src string) string {
	if len(src) == 0 {
		return src
	}

	return strings.ToUpper(src[:1]) + src[1:]
}

// EmptyStrPtr
// Determines the string pointer is nil or contains no value.
func EmptyStrPtr(v *string) bool {
	return v == nil || len(*v) == 0
}

// NotEmptyStrPtr
// Determines the string pointer is not nil and contains some value.
func NotEmptyStrPtr(v *string) bool {
	return !EmptyStrPtr(v)
}

func FirstNonDigit(str string) int {
	return strings.IndexFunc(str, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
}
