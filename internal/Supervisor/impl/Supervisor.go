package SupervisorImpl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"TranslateServer/internal/OsPlatform/api"
	"TranslateServer/internal/Supervisor/api"
)

type VocabTaskChannel struct {
	id          string
	decoderName string
	processProp OsPlatformApi.ProcessProp
}

type Supervisor struct {
	osInterface OsPlatformApi.OsInterface
	tasks       []VocabTaskChannel
}

func NewSupervisor(osInterface OsPlatformApi.OsInterface) SupervisorApi.SupervisorInterface {
	return &Supervisor{osInterface: osInterface, tasks: []VocabTaskChannel{}}
}

func (super *Supervisor) verifyInputFiles(decoder string, model string, vocab string) bool {
	return super.osInterface.FileExist(decoder) &&
		super.osInterface.FileExist(model) &&
		super.osInterface.FileExist(vocab)
}

func logMonitoring(id string, stram io.ReadCloser) {
	scanner := bufio.NewScanner(stram)
	for scanner.Scan() {
		fmt.Println("[", id, "]", scanner.Text())
	}
}

func (super *Supervisor) InitVocabTaskChannel(id string, decoder string, model string, vocab string) (SupervisorApi.Channel, error) {
	channel := SupervisorApi.Channel{}
	err := error(nil)

	if super.verifyInputFiles(decoder, model, vocab) {
		processProp, cmdErr := super.osInterface.AsyncCommand(
			decoder,
			"-m", model,
			"-v", vocab, vocab,
		)

		if cmdErr == nil {
			super.tasks = append(super.tasks, VocabTaskChannel{id: id, decoderName: decoder, processProp: processProp})
			go logMonitoring(id, processProp.Err)
			channel.In = processProp.In
			channel.Out = processProp.Out
		} else {
			err = cmdErr
		}
	} else {
		err = fmt.Errorf(id + "configuration files not found")
	}

	return channel, err
}

func (super *Supervisor) GetMetric() []SupervisorApi.ChannelStatus {
	var result []SupervisorApi.ChannelStatus

	for _, task := range super.tasks {
		var channelStatus SupervisorApi.ChannelStatus
		channelStatus.Id = task.id

		proc, _ := super.osInterface.GetProcess(task.processProp.Pid)
		if err := proc.Signal(0); err != nil {
			fmt.Println("[", task.id, "] error sending signal: ", err)
			channelStatus.Status = SupervisorApi.DEFECT
		} else {
			data, err := super.osInterface.ReadFile(fmt.Sprintf("/proc/%d/comm", task.processProp.Pid))
			if err != nil {
				fmt.Println("[", task.id, "] error reading file: ", err)
				channelStatus.Status = SupervisorApi.NOT_FOUND
			} else {
				procName := strings.TrimSpace(string(data))
				if strings.Contains(task.decoderName, procName) {
					fmt.Println("[", task.id, "] process is detected")
					channelStatus.Status = SupervisorApi.WORKING
				} else {
					fmt.Println("[", task.id, "] process is not detected")
					channelStatus.Status = SupervisorApi.NOT_FOUND
				}
			}
		}

		result = append(result, channelStatus)
	}

	return result
}
