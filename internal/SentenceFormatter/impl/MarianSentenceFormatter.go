package SentenceFormatterImpl

import (
	"regexp"
	"strings"
)

type MarianSentenceFormatter struct{}

func (f *MarianSentenceFormatter) PrepareInput(text string) string {
	text = strings.TrimSpace(text)

	re := regexp.MustCompile(`([.,!?:;])`)
	text = re.ReplaceAllString(text, " $1")

	reSpace := regexp.MustCompile(`\s+`)
	text = reSpace.ReplaceAllString(text, " ")

	if !regexp.MustCompile(`[.!?]$`).MatchString(text) {
		text += " ."
	}

	return text
}

func (f *MarianSentenceFormatter) CleanOutput(text string) string {
	text = strings.ReplaceAll(text, " ", "")

	reChar := regexp.MustCompile(`([.,!?:;])`)
	text = reChar.ReplaceAllString(text, "$1 ")

	text = strings.ReplaceAll(text, "▁", " ")

	reSpace := regexp.MustCompile(`\s+`)
	text = reSpace.ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)

	return text
}
