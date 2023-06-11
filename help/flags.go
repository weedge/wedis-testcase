package help

import "flag"

var srvBinPath = flag.String("srvBinPath", "", "directory including srv build files")

var dataDir = flag.String("dataDir", "", "directory including local kv store files")

var keepDataDir = flag.Bool("keepDataDir", true, "whether to keep data dire")
