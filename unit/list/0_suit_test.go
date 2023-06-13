package list

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
	"github.com/weedge/wedis-testcase/help"
)

var (
	gSrv *help.WedisSrv

	ctx context.Context
	c   redis.UniversalClient
)

func TestStringCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LIST CMD Suite")
}

var _ = BeforeSuite(func() {
	gSrv = help.StartSrv(map[string]string{})
	ctx = context.Background()
	c = gSrv.NewRedisClient()
})

var _ = AfterSuite(func() {
	Expect(c != nil).Should(BeTrue())
	Expect(c.Close()).NotTo(HaveOccurred())
	Expect(gSrv != nil).Should(BeTrue())
	gSrv.Close()
})