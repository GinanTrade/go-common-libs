package lib

import (
	"html"
	"strings"
)

// cleanString removes double quotes (") and single quotes (') from a string.
func CleanString(input string) string {
	// Replace double quotes (") with an empty string
	cleaned := strings.ReplaceAll(input, "\"", "")

	// Replace single quotes (') with an empty string
	cleaned = strings.ReplaceAll(cleaned, "'", "")

	return cleaned
}

func CleanHTMlString(input string) string {
	cleaned := html.EscapeString(input)
	return cleaned

}
