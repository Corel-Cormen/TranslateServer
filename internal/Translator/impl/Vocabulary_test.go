package TranslatorImpl

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"TranslateServer/internal/OsPlatform/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestVocabulary_GetIdAndGetProperties(t *testing.T) {
	v := Vocabulary{
		Id:      "v1",
		Decoder: "decoder",
		Model:   "model",
		Vocab:   "vocabulary",
	}

	require.Equal(t, "v1", v.GetId())

	props := v.GetProperties()
	require.Equal(t, "decoder", props.Decoder)
	require.Equal(t, "model", props.Model)
	require.Equal(t, "vocabulary", props.Vocab)
}

func TestVocabulary_Translate_Success(t *testing.T) {
	inBuf := &MockOsPlatformApi.NopWriteCloser{Buffer: new(bytes.Buffer)}
	outBuf := bytes.NewBufferString("translated\n")

	v := &Vocabulary{}
	require.NoError(t, v.RegisterInput(inBuf))
	require.NoError(t, v.RegisterOutput(io.NopCloser(outBuf)))

	result, err := v.Translate("hello")
	require.NoError(t, err)
	require.Equal(t, "translated", result)

	require.Equal(t, "hello\n", inBuf.String())
}

func TestVocabulary_Translate_ErrorOnWrite(t *testing.T) {
	mockIn := new(MockOsPlatformApi.MockWriteCloser)
	mockIn.Buf = &bytes.Buffer{}
	mockIn.On("Write", mock.Anything).Return(0, errors.New("write error"))

	v := Vocabulary{}
	require.NoError(t, v.RegisterInput(mockIn))

	_, err := v.Translate("text")
	require.Error(t, err)
	require.EqualError(t, err, "write error")

	mockIn.AssertExpectations(t)
}

func TestVocabulary_Unregister(t *testing.T) {
	mockIn := new(MockOsPlatformApi.MockWriteCloser)
	mockIn.On("Close").Return(nil)

	v := Vocabulary{}
	_ = v.RegisterInput(mockIn)

	err := v.Unregister()
	require.NoError(t, err)

	mockIn.AssertExpectations(t)
}
