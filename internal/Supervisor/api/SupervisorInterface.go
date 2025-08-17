package SupervisorApi

import "io"

type Channel struct {
	In  io.WriteCloser
	Out io.ReadCloser
}

var (
	WORKING   int = 0
	NOT_FOUND int = 1
	DEFECT    int = 2
)

type ChannelStatus struct {
	Id     string
	Status int
}

type SupervisorInterface interface {
	InitVocabTaskChannel(id string, decoder string, model string, vocab string) (Channel, error)
	GetMetric() []ChannelStatus
}
