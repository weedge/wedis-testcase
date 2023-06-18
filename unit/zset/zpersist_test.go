package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zpersist Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zpersist"
		arrZ := []redis.Z{
			{Score: 1, Member: "v1"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.Do(ctx, "zEXPIRE", k, "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "zTTL", k).Val()).To(Equal(int64(100)))
		Expect(c.Do(ctx, "zPERSIST", k).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "zTTL", k).Val()).To(Equal(int64(-1)))
	})
	It("no exists key", func() {
		Expect(c.Do(ctx, "zPERSIST", "noexistzkey").Val()).To(Equal(int64(0)))
	})
})
