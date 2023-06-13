import requests
import os

wedisSrvUrl = "http://127.0.0.1:8110"


def getCmds(cmdType: str = None) -> object:
    r = requests.get("{}/{}".format(wedisSrvUrl, "cmds"))
    if r.status_code != 200:
        return None
    res = r.json()
    if cmdType is None:
        return res
    return res[cmdType]


if __name__ == '__main__':
    cmds = getCmds()
    for cmdType in cmds:
        cmdTypeDir = "./{}/{}".format("unit", cmdType)
        os.makedirs(cmdTypeDir, exist_ok=True)
        for cmd in cmds[cmdType]:
            cmdTestFile = cmdTypeDir+"/{}_test.go".format(cmd)
            if os.path.exists(cmdTestFile):
                print(cmdTestFile, "exists")
                continue

            tf = open("./scripts/tpl/ginkgo_cmd_test.tpl", 'r')
            testStr = tf.read().replace("{{CMD_TYPE}}", cmdType).replace(
                "{{CMD}}", cmd)
            tf.close()

            f = open(cmdTestFile, 'w')
            f.write(testStr)
            f.close()
            print(cmdTestFile, "create ok")
    print("gen success, have fun :)")
