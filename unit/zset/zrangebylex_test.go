package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zrangebylex Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zrangebylex"
		arrZ := []redis.Z{
			{Score: 0, Member: "a"},
			{Score: 0, Member: "b"},
			{Score: 0, Member: "c"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRangeByLex(ctx, k, &redis.ZRangeBy{
			Min: "-",
			Max: "+",
		}).Val()).To(Equal([]string{"a", "b", "c"}))
		Expect(c.ZRangeByLex(ctx, k, &redis.ZRangeBy{
			Min:    "-",
			Max:    "+",
			Offset: 0,
			Count:  2,
		}).Val()).To(Equal([]string{"a", "b"}))
		Expect(c.ZRangeByLex(ctx, k, &redis.ZRangeBy{
			Min:    "-",
			Max:    "+",
			Offset: 1,
			Count:  2,
		}).Val()).To(Equal([]string{"b", "c"}))
		Expect(c.ZRangeByLex(ctx, k, &redis.ZRangeBy{
			Min: "(a", Max: "(c",
			Offset: 0,
			Count:  2,
		}).Val()).To(Equal([]string{"b"}))
		Expect(c.ZRangeByLex(ctx, k, &redis.ZRangeBy{
			Min: "[a", Max: "(c",
			Offset: 0,
			Count:  3,
		}).Val()).To(Equal([]string{"a", "b"}))
		Expect(c.ZRangeByLex(ctx, k, &redis.ZRangeBy{
			Min: "(a", Max: "[c",
			Offset: 0,
			Count:  3,
		}).Val()).To(Equal([]string{"b", "c"}))
		Expect(c.ZRangeByLex(ctx, k, &redis.ZRangeBy{
			Min: "[a", Max: "[c",
			Offset: 0,
			Count:  4,
		}).Val()).To(Equal([]string{"a", "b", "c"}))
	})

	It("nokey", func() {
		k := "zrangebylexnokey"
		Expect(c.ZRangeByLex(ctx, k, &redis.ZRangeBy{
			Min: "-", Max: "+",
		}).Val()).To(Equal([]string{}))
	})

	It("ERR wrong number of arguments", func() {
		k := "zrangebylexerror"
		Expect(c.ZRangeByLex(ctx, k, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
		}).Err().Error()).To(ContainSubstring("ERR wrong number of arguments"))
	})
})
