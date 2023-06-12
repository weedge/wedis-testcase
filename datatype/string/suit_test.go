package string

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/weedge/wedis-testcase/help"
)

var gSrv *help.WedisSrv

func TestStringCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Type String CMD Suite")
}

var _ = BeforeSuite(func() {
	gSrv = help.StartSrv(map[string]string{})
})

var _ = AfterSuite(func() {
	Expect(gSrv != nil).Should(BeTrue())
	gSrv.Close()
})
