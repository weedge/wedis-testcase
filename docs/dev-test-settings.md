.vscode/settings.json config test sample

default config test case
```json
{
    "files.associations": {
        "*.go": "go",
    },
    "go.testEnvVars": {
        "srvBinPath": "/Users/wuyong/go/src/github.com/weedge/wedis/bin/wedis",
        "dataDir": "/tmp/wedis-testcase",
        //"keepDataDir": "1", //len(keepDataDir)>0 true
        "srvPort": "6669",
        //"confTestCase": "default",
    },
    "ginkgotestexplorer.testEnvVars": {
        "srvBinPath": "/Users/wuyong/go/src/github.com/weedge/wedis/bin/wedis",
        "dataDir": "/tmp/wedis-testcase",
        //"keepDataDir": "1", //len(keepDataDir)>0 true
        "srvPort": "6669",
        //"confTestCase": "default",
    }
}
```

xdis-tikv config test case
```json
{
    "files.associations": {
        "*.go": "go",
    },
    "go.testEnvVars": {
        "srvBinPath": "/Users/wuyong/go/src/github.com/weedge/wedis/bin/wedis",
        "dataDir": "/tmp/wedis-testcase",
        //"keepDataDir": "1", //len(keepDataDir)>0 true
        "srvPort": "6669",
        "confTestCase": "xdis-tikv",
    },
    "ginkgotestexplorer.testEnvVars": {
        "srvBinPath": "/Users/wuyong/go/src/github.com/weedge/wedis/bin/wedis",
        "dataDir": "/tmp/wedis-testcase",
        //"keepDataDir": "1", //len(keepDataDir)>0 true
        "srvPort": "6669",
        "confTestCase": "xdis-tikv",
    }
}
```