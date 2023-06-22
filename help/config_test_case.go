package help

import "fmt"

// add some config test case
var mapConfigTestCases = map[string]map[string]string{
	"default": {},
	"xdis-tikv": {
		"server.storagerName":                   "xdis-tikv",
		"tikvStoreCfg.prefixKey":                "wedis-testcase-tikv",
		"tikvStoreCfg.ttlCheckInterval":         "1",
		"tikvStoreCfg.tikvClientOpts.useTxnApi": "1",
	},
}

func GetConfigTestCase(name string) (data map[string]string) {
	defer func() {
		fmt.Printf("name:%s testCaseConf: %+v\n", name, data)
	}()

	data, ok := mapConfigTestCases[name]
	if ok {
		return
	}

	return map[string]string{}
}
