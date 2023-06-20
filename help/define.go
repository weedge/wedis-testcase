package help

import (
	"time"
)

const (
	DefaultContainerGracePeriod = 30 * time.Second

	EnvSrvBinPath  = "srvBinPath"
	EnvDataDir     = "dataDir"
	EnvKeepDataDir = "keepDataDir"
	EnvSrvPort     = "srvPort"

	EnvConfTestCase = "confTestCase"
)
