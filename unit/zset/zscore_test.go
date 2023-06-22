package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zscore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zscore"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "a"},
			{Score: 3, Member: "b"},
			{Score: 2, Member: "c"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZScore(ctx, k, "a").Val()).To(Equal(float64(1)))
		Expect(c.ZScore(ctx, k, "b").Val()).To(Equal(float64(3)))
		Expect(c.ZScore(ctx, k, "c").Val()).To(Equal(float64(2)))
		Expect(c.ZScore(ctx, k, "ccc").Val()).To(Equal(float64(0)))
	})
})
