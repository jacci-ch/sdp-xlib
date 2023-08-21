// Copyright 2023 to now() The SDP Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package valuex

import (
	"strings"
	"unicode"
)

// HasPrefixIn - detects whether str has one of the prefix in
// the given list. This function returns false if no prefix matches.
func HasPrefixIn(str string, prefixes []string) bool {
	for _, v := range prefixes {
		if strings.HasPrefix(str, v) {
			return true
		}
	}

	return false
}

// CamelToUnderline - converts src string from Camel-Format (e.g StringInCamelFormat)
// to Underline-Format (e.g string_in_underline_format).
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

// FirstToLower - converts the first rune to lower-case.
func FirstToLower(src string) string {
	if len(src) == 0 {
		return src
	}

	return strings.ToLower(src[:1]) + src[1:]
}

// FirstToUpper - converts the first rune to upper-case.
func FirstToUpper(src string) string {
	if len(src) == 0 {
		return src
	}

	return strings.ToUpper(src[:1]) + src[1:]
}

// EmptyStr - detects the string pointer is not nil and contains some value.
func EmptyStr(v *string) bool {
	return v == nil || len(*v) == 0
}

// FirstNonDigit - retrieves  the index of first non-digit (not a digit value) rune.
func FirstNonDigit(str string) int {
	return strings.IndexFunc(str, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
}
