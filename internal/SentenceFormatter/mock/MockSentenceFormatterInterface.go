package MockSentenceFormatter

import "github.com/stretchr/testify/mock"

type MockSentenceFormatterInterface struct {
	mock.Mock
}

func (f *MockSentenceFormatterInterface) PrepareInput(text string) string {
	args := f.Called(text)
	return args.String(0)
}

func (f *MockSentenceFormatterInterface) CleanOutput(text string) string {
	args := f.Called(text)
	return args.String(0)
}
