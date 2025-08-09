package OsPlatform

import (
	"TranslateServer/internal/OsPlatform/api"
	"TranslateServer/internal/OsPlatform/impl"
)

func GetOsInstance() OsPlatformApi.OsInterface {
	return &OsPlatformImpl.OsFacade{}
}
