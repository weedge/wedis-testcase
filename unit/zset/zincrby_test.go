package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zincrby Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key := "zincrby"
		Expect(c.Do(ctx, "zMCLEAR", key).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZIncrBy(ctx, key, 10, "m1").Val()).To(Equal(float64(11)))
		Expect(c.ZIncrBy(ctx, key, 10, "m2").Val()).To(Equal(float64(12)))
		Expect(c.ZIncrBy(ctx, key, 10, "m3").Val()).To(Equal(float64(13)))
		Expect(c.ZIncrBy(ctx, key, 10, "m4").Val()).To(Equal(float64(10)))
	})
})
