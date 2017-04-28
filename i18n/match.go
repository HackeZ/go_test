package i18n

import (
	"golang.org/x/text/language"
)

// LangMatch ...
func LangMatch(speak string) {
	preferred := []language.Tag{
		language.Chinese, // The first language is used as fallback.
		language.Spanish,
		language.Norwegian,
		language.MustParse("en-AU"),
	}

	match := language.NewMatcher(preferred)
	match.Match(preferred...)
}
