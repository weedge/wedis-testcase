package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zremrangebyscore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zremrangebyscore"
		arrZ := []redis.Z{
			{Score: 0, Member: "a"},
			{Score: 1, Member: "b"},
			{Score: 2, Member: "c"},
			{Score: 3, Member: "foo"},
			{Score: 4, Member: "zap"},
			{Score: 5, Member: "zip"},
			{Score: 6, Member: "ALPHA"},
			{Score: 7, Member: "alpha"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"a", "b", "c", "foo", "zap", "zip", "ALPHA", "alpha",
		}))
		Expect(c.ZRemRangeByScore(ctx, k, "1", "3").Val()).To(Equal(int64(3)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"a", "zap", "zip", "ALPHA", "alpha",
		}))
		Expect(c.ZRemRangeByScore(ctx, k, "1", "3").Val()).To(Equal(int64(0)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"a", "zap", "zip", "ALPHA", "alpha",
		}))
		Expect(c.ZRemRangeByScore(ctx, k, "4", "6").Val()).To(Equal(int64(3)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"a", "alpha",
		}))
		Expect(c.ZRemRangeByScore(ctx, k, "-inf", "+inf").Val()).To(Equal(int64(2)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{}))
	})

	It("nokey", func() {
		k := "zremrangebyscorenokey"
		Expect(c.ZRemRangeByScore(ctx, k, "-inf", "+inf").Val()).To(Equal(int64(0)))
	})
})
