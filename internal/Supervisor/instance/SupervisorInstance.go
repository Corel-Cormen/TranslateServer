package SupervisorInstance

import (
	"sync"

	"TranslateServer/internal/OsPlatform"
	"TranslateServer/internal/Supervisor/api"
	"TranslateServer/internal/Supervisor/impl"
)

var (
	supervisorInstance     SupervisorApi.SupervisorInterface
	onceSupervisorInstance sync.Once
)

func GetSupervisorInstance() SupervisorApi.SupervisorInterface {
	onceSupervisorInstance.Do(func() {
		supervisorInstance = SupervisorImpl.NewSupervisor(OsPlatform.GetOsInstance())
	})
	return supervisorInstance
}
