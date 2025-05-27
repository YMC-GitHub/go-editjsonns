// Package jsonns provides functionality for parsing and manipulating JSON namespace expressions
package jsonns

import (
	"regexp"
	"strings"
)

// NSHasKeyArrayIndex checks if a string contains array index notation (e.g., "[0]")
func NSHasKeyArrayIndex(s string) bool {
	re := regexp.MustCompile(`\[\d+\]`)
	return re.MatchString(s)
}

// NSGetKeyArrayIndex extracts the array index from a string (e.g., "names[0]" -> "0")
func NSGetKeyArrayIndex(s string) *string {
	re := regexp.MustCompile(`\[(\d+)\]`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 1 {
		return &matches[1]
	}
	return nil
}

// NSGetKeyArrKey extracts the key part from an array expression (e.g., "names[0]" -> "names")
func NSGetKeyArrKey(s string) string {
	re := regexp.MustCompile(`\[\d+\]`)
	return re.ReplaceAllString(s, "")
}

// KeyArrObj represents a key-array object with key and index
type KeyArrObj struct {
	Key   string  `json:"key"`
	Index *string `json:"index"`
}

// NSKeyarrObjify converts a string into a KeyArrObj
func NSKeyarrObjify(s string) KeyArrObj {
	return KeyArrObj{
		Key:   NSGetKeyArrKey(s),
		Index: NSGetKeyArrayIndex(s),
	}
}

// NSPureName removes brackets from a string (e.g., "[0]" -> "0")
func NSPureName(s string) string {
	return strings.Trim(s, "[]")
}

// NSGetMatch finds all array expressions in a string
func NSGetMatch(s string, reg *regexp.Regexp, preset string) []string {
	if reg == nil {
		reg = regexp.MustCompile(`\[[^\s\[\]]+\]`)
	}

	presets := strings.Split(preset, ",")
	for _, p := range presets {
		p = strings.TrimSpace(p)
		switch p {
		case "only-number":
			reg = regexp.MustCompile(`\[\d+\]`)
		case "allow-string":
			reg = regexp.MustCompile(`\[[^\s]+\]`)
		}
	}

	return reg.FindAllString(s, -1)
}

// NSStd standardizes a namespace string into its path components
//
// Parameters:
//
//	s - input path string (e.g. "a.b.c", "users[0].name")
//	sep - separator character (defaults to ".")
//	reg - optional custom regex for array matching
//
// Returns:
//
//	[]string - array of standardized path components
//
// Examples:
//
//	NSStd("a.b.c", ".", nil) → ["a", "b", "c"]
//	NSStd("users[0].address.city", ".", nil) → ["users", "[0]", "address", "city"]
//	NSStd("data/items/0/id", "/", nil) → ["data", "items", "0", "id"]
func NSStd(s string, sep string, reg *regexp.Regexp) []string {
	if sep == "" {
		sep = "."
	}
	if reg == nil {
		reg = regexp.MustCompile(`\[[^\s\[\]]+\]`)
	}

	tmp := NSStdDotTypeKey(s, sep)
	var result []string

	for _, item := range tmp {
		arrTypeify := NSStdArrTypeKey(item, reg)
		result = append(result, arrTypeify...)
	}

	return result
}

// NSStdArrTypeKey breaks down array notation expressions
func NSStdArrTypeKey(s string, reg *regexp.Regexp) []string {
	if reg == nil {
		reg = regexp.MustCompile(`\[[^\s\[\]]+\]`)
	}

	matches := NSGetMatch(s, reg, "")
	head := reg.ReplaceAllString(s, "")

	if len(matches) > 0 {
		if head != "" {
			return append([]string{head}, matches...)
		}
		return matches
	}
	return []string{head}
}

// NSStdDotTypeKey splits a string by separator
func NSStdDotTypeKey(s string, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	parts := strings.Split(s, sep)
	var result []string
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}
