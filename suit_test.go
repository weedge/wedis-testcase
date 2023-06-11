package wedistestcase

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/weedge/wedis-testcase/help"
)

var gSrv *help.WedisSrv

func TestRESPCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "redis RESP Cmd Suite")
}

var _ = BeforeSuite(func() {
	gSrv = help.StartSrv(map[string]string{})
})

var _ = AfterSuite(func() {
	Expect(gSrv).NotTo(HaveOccurred())
	gSrv.Close()
})

var _ = Describe("init", func() {
	It("ok", func() {
		println("ok")
	})
})
