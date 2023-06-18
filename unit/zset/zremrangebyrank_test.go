package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zremrangebyrank Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zremrangebyrank"
		arrZ := []redis.Z{
			{Score: 0, Member: "a"},
			{Score: 0, Member: "b"},
			{Score: 0, Member: "c"},
			{Score: 0, Member: "foo"},
			{Score: 0, Member: "zap"},
			{Score: 0, Member: "zip"},
			{Score: 0, Member: "ALPHA"},
			{Score: 0, Member: "alpha"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"ALPHA", "a", "alpha", "b", "c", "foo", "zap", "zip",
		}))
		Expect(c.ZRemRangeByRank(ctx, k, 0, 1).Val()).To(Equal(int64(2)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"alpha", "b", "c", "foo", "zap", "zip",
		}))
		Expect(c.ZRemRangeByRank(ctx, k, 0, 1).Val()).To(Equal(int64(2)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"c", "foo", "zap", "zip",
		}))
		Expect(c.ZRemRangeByRank(ctx, k, 0, 1).Val()).To(Equal(int64(2)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"zap", "zip",
		}))
		Expect(c.ZRemRangeByRank(ctx, k, 0, 1).Val()).To(Equal(int64(2)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{}))
	})

	It("nokey", func() {
		k := "zremrangebyranknokey"
		Expect(c.ZRemRangeByRank(ctx, k, 0, -1).Val()).To(Equal(int64(0)))
	})
})
