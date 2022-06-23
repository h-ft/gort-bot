package commands

import (
	"regexp"
	"strings"
)

func isYouTubeLink(link string) bool {
	if strings.Contains(link, "youtu") || strings.ContainsAny(link, "\"?&/<%=") {
		matchers := []*regexp.Regexp{
			regexp.MustCompile(`(?:v|embed|watch\?v)(?:=|/)([^"&?/=%]{11})`),
			regexp.MustCompile(`(?:=|/)([^"&?/=%]{11})`),
			regexp.MustCompile(`([^"&?/=%]{11})`),
		}

		for _, re := range matchers {
			if isMatch := re.MatchString(link); isMatch {
				return true
			}
		}
	}

	return false
}
