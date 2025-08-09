package OsInstance

import (
	"TranslateServer/internal/OsPlatform/api"
	"TranslateServer/internal/OsPlatform/impl"
	"sync"
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
