package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zkeyexists Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zkeyexists"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.Do(ctx, "zKEYEXISTS", k).Val()).To(Equal(int64(1)))
	})
	It("no", func() {
		Expect(c.Do(ctx, "zKEYEXISTS", "zk11s").Val()).To(Equal(int64(0)))
	})
})
