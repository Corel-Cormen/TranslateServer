package OsInstance

import (
	"sync"

	"TranslateServer/internal/OsPlatform/api"
	"TranslateServer/internal/OsPlatform/impl"
)

var (
	osInstance     OsPlatformApi.OsInterface
	onceOsInstance sync.Once
)

func GetOsInstance() OsPlatformApi.OsInterface {
	onceOsInstance.Do(func() {
		osInstance = &OsPlatformImpl.OsFacade{}
	})
	return osInstance
}
