package MockSupervisorApi

import (
	"TranslateServer/internal/Supervisor/api"
	"github.com/stretchr/testify/mock"
)

type MockSupervisor struct {
	mock.Mock
}

func (m *MockSupervisor) InitVocabTaskChannel(id string, decoder string, model string, vocab string) (SupervisorApi.Channel, error) {
	args := m.Called(id, decoder, model, vocab)
	return args.Get(0).(SupervisorApi.Channel), args.Error(1)
}

func (m *MockSupervisor) GetMetric() []SupervisorApi.ChannelStatus {
	args := m.Called()
	return args.Get(0).([]SupervisorApi.ChannelStatus)
}
