package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zrevrank Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zrevrank"
		arrZ := []redis.Z{
			{Score: 1, Member: "a"},
			{Score: 3, Member: "b"},
			{Score: 2, Member: "c"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRevRank(ctx, k, "b").Val()).To(Equal(int64(0)))
		Expect(c.ZRevRank(ctx, k, "a").Val()).To(Equal(int64(2)))
		Expect(c.ZRevRank(ctx, k, "c").Val()).To(Equal(int64(1)))
		Expect(c.ZRevRank(ctx, k, "aaa").Err().Error()).To(Equal("redis: nil"))
	})

	It("ERR invalid zset member size", func() {
		k := "zrevrankerr"
		Expect(c.ZRevRank(ctx, k, "").Err().Error()).To(Equal("ERR invalid zset member size"))
	})
	It("nokey", func() {
		k := "zrevranknokey"
		Expect(c.ZRevRank(ctx, k, "a").Err().Error()).To(Equal("redis: nil"))
	})
})
