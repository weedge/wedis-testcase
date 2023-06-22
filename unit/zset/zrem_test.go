package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zrem Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zrem"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "a"},
			{Score: 3, Member: "b"},
			{Score: 2, Member: "c"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRem(ctx, k, "a").Val()).To(Equal(int64(1)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{"c", "b"}))
		Expect(c.ZRem(ctx, k, "a").Val()).To(Equal(int64(0)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{"c", "b"}))
		Expect(c.ZRem(ctx, k, "c").Val()).To(Equal(int64(1)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{"b"}))
		Expect(c.ZRem(ctx, k, "cc").Val()).To(Equal(int64(0)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{"b"}))
		Expect(c.ZRem(ctx, k, "b").Val()).To(Equal(int64(1)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{}))
	})

	It("nokey", func() {
		k := "zremnokey"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.ZRem(ctx, k, "a").Val()).To(Equal(int64(0)))
	})
})
