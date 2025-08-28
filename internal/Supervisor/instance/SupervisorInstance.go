package SupervisorInstance

import (
	"sync"

	"TranslateServer/internal/OsPlatform/instance"
	"TranslateServer/internal/Supervisor/api"
	"TranslateServer/internal/Supervisor/impl"
)

var (
	supervisorInstance     SupervisorApi.SupervisorInterface
	onceSupervisorInstance sync.Once
)

func GetSupervisorInstance() SupervisorApi.SupervisorInterface {
	onceSupervisorInstance.Do(func() {
		supervisorInstance = SupervisorImpl.NewSupervisor(OsInstance.GetOsInstance())
	})
	return supervisorInstance
}
