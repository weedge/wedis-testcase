package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zrank Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zrank"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "a"},
			{Score: 3, Member: "b"},
			{Score: 2, Member: "c"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRank(ctx, k, "b").Val()).To(Equal(int64(2)))
		Expect(c.ZRank(ctx, k, "a").Val()).To(Equal(int64(0)))
		Expect(c.ZRank(ctx, k, "c").Val()).To(Equal(int64(1)))
		Expect(c.ZRank(ctx, k, "aaa").Err().Error()).To(Equal("redis: nil"))
	})

	It("ERR invalid zset member size", func() {
		k := "zrankerr"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.ZRank(ctx, k, "").Err().Error()).To(Equal("ERR invalid zset member size"))
	})
	It("nokey", func() {
		k := "zranknokey"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.ZRank(ctx, k, "a").Err().Error()).To(Equal("redis: nil"))
	})
})
