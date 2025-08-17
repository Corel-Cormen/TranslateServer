package MockSupervisorApi

import (
	"TranslateServer/internal/Supervisor/api"
	"github.com/stretchr/testify/mock"
)

type MockSupervisorApi struct {
	mock.Mock
}

func (m *MockSupervisorApi) InitVocabTaskChannel(id string, decoder string, model string, vocab string) (SupervisorApi.Channel, error) {
	args := m.Called(id, decoder, model, vocab)
	return args.Get(0).(SupervisorApi.Channel), args.Error(1)
}

func (m *MockSupervisorApi) GetMetric() []SupervisorApi.ChannelStatus {
	args := m.Called()
	return args.Get(0).([]SupervisorApi.ChannelStatus)
}
