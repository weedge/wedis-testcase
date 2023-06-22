package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zrange Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zrange"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "v1"},
			{Score: 2, Member: "v2"},
			{Score: 3, Member: "v3"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"v1", "v2", "v3",
		}))
		Expect(c.ZRangeWithScores(ctx, k, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 1, Member: "v1"},
			{Score: 2, Member: "v2"},
			{Score: 3, Member: "v3"},
		}))
	})
	It("nokey", func() {
		k := "nokeyzrange"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{}))
	})
})
