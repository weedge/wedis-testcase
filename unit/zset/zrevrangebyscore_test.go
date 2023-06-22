package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zrevrangebyscore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zrevrangebyscore"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "a"},
			{Score: 3, Member: "b"},
			{Score: 2, Member: "c"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
		}).Val()).To(Equal([]string{"b", "c", "a"}))
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
			Offset: 0,
			Count:  2,
		}).Val()).To(Equal([]string{"b", "c"}))
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
			Offset: 1,
			Count:  2,
		}).Val()).To(Equal([]string{"c", "a"}))
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "(1", Max: "+inf",
		}).Val()).To(Equal([]string{"b", "c"}))
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "(1", Max: "+inf",
			Offset: 1,
			Count:  2,
		}).Val()).To(Equal([]string{"c"}))
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "(1", Max: "(3",
			Offset: 0,
			Count:  4,
		}).Val()).To(Equal([]string{"c"}))
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "(1", Max: "3",
			Offset: 0,
			Count:  4,
		}).Val()).To(Equal([]string{"b", "c"}))
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "1", Max: "3",
			Offset: 0,
			Count:  4,
		}).Val()).To(Equal([]string{"b", "c", "a"}))
	})

	It("nokey", func() {
		k := "zrevrangebyscorenokey"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
		}).Val()).To(Equal([]string{}))
	})

	It("ERR value is not an integer or out of range", func() {
		k := "zrevrangebyscoreerror"
		Expect(c.ZRevRangeByScore(ctx, k, &redis.ZRangeBy{
			Min: "-", Max: "+",
		}).Err().Error()).To(ContainSubstring("ERR value is not an integer or out of range"))
	})
})
