package OsPlatform

import (
	"TranslateServer/internal/OsPlatform/api"
	"TranslateServer/internal/OsPlatform/instance"
)

func GetOsInstance() OsPlatformApi.OsInterface {
	return OsInstance.GetOsInstance()
}
