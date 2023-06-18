package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zrangebyscore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zrangebyscore"
		arrZ := []redis.Z{
			{Score: 1, Member: "a"},
			{Score: 3, Member: "b"},
			{Score: 2, Member: "c"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
		}).Val()).To(Equal([]string{"a", "c", "b"}))
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
			Offset: 0,
			Count:  2,
		}).Val()).To(Equal([]string{"a", "c"}))
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
			Offset: 1,
			Count:  2,
		}).Val()).To(Equal([]string{"c", "b"}))
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "(1", Max: "+inf",
		}).Val()).To(Equal([]string{"c", "b"}))
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "(1", Max: "+inf",
			Offset: 1,
			Count:  2,
		}).Val()).To(Equal([]string{"b"}))
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "(1", Max: "(3",
			Offset: 0,
			Count:  4,
		}).Val()).To(Equal([]string{"c"}))
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "(1", Max: "3",
			Offset: 0,
			Count:  4,
		}).Val()).To(Equal([]string{"c", "b"}))
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "1", Max: "3",
			Offset: 0,
			Count:  4,
		}).Val()).To(Equal([]string{"a", "c", "b"}))
	})

	It("nokey", func() {
		k := "zrangebyscorenokey"
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
		}).Val()).To(Equal([]string{}))
	})

	It("ERR value is not an integer or out of range", func() {
		k := "zrangebyscoreerror"
		Expect(c.ZRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-", Max: "+",
		}).Err().Error()).To(ContainSubstring("ERR value is not an integer or out of range"))
	})
})
