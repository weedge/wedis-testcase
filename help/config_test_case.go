package help

// add some config test case
var mapConfigTestCases = map[string]map[string]string{
	"default": {},
	"xdis-tikv": {
		"server.storagerName":    "xdis-tikv",
		"tikvStoreCfg.prefixKey": "wedis-testcase-tikv",
	},
}

func GetConfigTestCase(name string) map[string]string {
	data, ok := mapConfigTestCases[name]
	if ok {
		return data
	}

	return map[string]string{}
}
