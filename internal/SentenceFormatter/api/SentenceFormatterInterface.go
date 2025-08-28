package SentenceFormatterApi

type SentenceFormatterInterface interface {
	PrepareInput(text string) string
	CleanOutput(text string) string
}
