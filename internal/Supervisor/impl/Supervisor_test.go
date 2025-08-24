package SupervisorImpl

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/mock"
	"io"
	"testing"

	"TranslateServer/internal/OsPlatform/api"
	"TranslateServer/internal/OsPlatform/mock"
	"TranslateServer/internal/Supervisor/api"
	"github.com/stretchr/testify/require"
)

func TestSupervisor_InitVocabTaskChannel_Success(t *testing.T) {
	mockOs := new(MockOsPlatformApi.MockOsInterface)
	supervisor := NewSupervisor(mockOs)

	decoder := "decoder.bin"
	model := "model.bin"
	vocab := "vocab.bin"
	mockOs.On("FileExist", decoder).Return(true)
	mockOs.On("FileExist", model).Return(true)
	mockOs.On("FileExist", vocab).Return(true)

	inCh := &MockOsPlatformApi.NopWriteCloser{}
	outCh := io.NopCloser(bytes.NewBuffer(nil))
	errCh := io.NopCloser(bytes.NewBufferString("channel log"))
	processProp := OsPlatformApi.ProcessProp{Pid: 100, In: inCh, Out: outCh, Err: errCh}
	mockOs.On("AsyncCommand", decoder, "-m", model, "-v", vocab, vocab).Return(processProp, nil)

	ch, err := supervisor.InitVocabTaskChannel("task", decoder, model, vocab)

	require.NoError(t, err)
	require.Equal(t, inCh, ch.In)
	require.Equal(t, outCh, ch.Out)

	mockOs.AssertExpectations(t)
}

func TestSupervisor_InitVocabTaskChannel_AsyncCommandFail(t *testing.T) {
	mockOs := new(MockOsPlatformApi.MockOsInterface)
	supervisor := NewSupervisor(mockOs)

	decoder := "decoder.bin"
	model := "model.bin"
	vocab := "vocab.bin"
	mockOs.On("FileExist", decoder).Return(true)
	mockOs.On("FileExist", model).Return(true)
	mockOs.On("FileExist", vocab).Return(true)

	mockOs.On("AsyncCommand", decoder, "-m", model, "-v", vocab, vocab).Return(
		OsPlatformApi.ProcessProp{}, errors.New("exec error"))

	ch, err := supervisor.InitVocabTaskChannel("task", decoder, model, vocab)

	require.Error(t, err)
	require.Empty(t, ch.In)
	require.Empty(t, ch.Out)

	mockOs.AssertExpectations(t)
}

func TestSupervisor_InitVocabTaskChannel_FileMissing(t *testing.T) {
	decoder := "decoder.bin"
	model := "model.bin"
	vocab := "vocab.bin"

	cases := []struct {
		name       string
		filesExist map[string]bool
	}{
		{
			name: "decoder missing",
			filesExist: map[string]bool{
				decoder: false,
				model:   true,
				vocab:   true,
			},
		},
		{
			name: "model missing",
			filesExist: map[string]bool{
				decoder: true,
				model:   false,
				vocab:   true,
			},
		},
		{
			name: "vocab missing",
			filesExist: map[string]bool{
				decoder: true,
				model:   true,
				vocab:   false,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mockOs := new(MockOsPlatformApi.MockOsInterface)
			supervisor := NewSupervisor(mockOs)

			mockOs.On("FileExist", mock.Anything).
				Run(func(args mock.Arguments) {
					path := args.String(0)
					if tc.filesExist[path] {
						mockOs.ExpectedCalls[0].ReturnArguments = mock.Arguments{true}
					} else {
						mockOs.ExpectedCalls[0].ReturnArguments = mock.Arguments{false}
					}
				})

			ch, err := supervisor.InitVocabTaskChannel("task", decoder, model, vocab)

			require.Error(t, err)
			require.Nil(t, ch.In)
			require.Nil(t, ch.Out)

			mockOs.AssertExpectations(t)
		})
	}
}

func TestGetMetric(t *testing.T) {
	mockOs := new(MockOsPlatformApi.MockOsInterface)
	supervisor := &Supervisor{
		osInterface: mockOs,
		tasks: []VocabTaskChannel{
			{id: "t1", decoderName: "decoder1", processProp: OsPlatformApi.ProcessProp{Pid: 1}},
			{id: "t2", decoderName: "decoder2", processProp: OsPlatformApi.ProcessProp{Pid: 2}},
			{id: "t3", decoderName: "decoder3", processProp: OsPlatformApi.ProcessProp{Pid: 3}},
			{id: "t4", decoderName: "decoder4", processProp: OsPlatformApi.ProcessProp{Pid: 4}},
		},
	}

	proc1 := new(MockOsPlatformApi.MockProcessInterface)
	proc1.On("Signal", 0).Return(nil)
	proc2 := new(MockOsPlatformApi.MockProcessInterface)
	proc2.On("Signal", 0).Return(errors.New("signal error"))
	proc3 := new(MockOsPlatformApi.MockProcessInterface)
	proc3.On("Signal", 0).Return(nil)
	proc4 := new(MockOsPlatformApi.MockProcessInterface)
	proc4.On("Signal", 0).Return(nil)

	mockOs.On("GetProcess", 1).Return(proc1, nil)
	mockOs.On("GetProcess", 2).Return(proc2, nil)
	mockOs.On("GetProcess", 3).Return(proc3, nil)
	mockOs.On("GetProcess", 4).Return(proc4, nil)

	mockOs.On("ReadFile", "/proc/1/comm").Return([]byte("decoder1\n"), nil)
	mockOs.On("ReadFile", "/proc/3/comm").Return([]byte("other_process\n"), nil)
	mockOs.On("ReadFile", "/proc/4/comm").Return([]byte(""), errors.New("read error"))

	metrics := supervisor.GetMetric()

	require.Len(t, metrics, 4)
	require.Equal(t, SupervisorApi.WORKING, metrics[0].Status)
	require.Equal(t, SupervisorApi.DEFECT, metrics[1].Status)
	require.Equal(t, SupervisorApi.NOT_FOUND, metrics[2].Status)
	require.Equal(t, SupervisorApi.NOT_FOUND, metrics[3].Status)

	mockOs.AssertExpectations(t)
	proc1.AssertExpectations(t)
	proc2.AssertExpectations(t)
	proc3.AssertExpectations(t)
	proc4.AssertExpectations(t)
}
