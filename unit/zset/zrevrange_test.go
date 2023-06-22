package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zrevrange Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zrevrange"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 0, Member: "a"},
			{Score: 1, Member: "b"},
			{Score: 2, Member: "c"},
			{Score: 3, Member: "d"},
			{Score: 4, Member: "e"},
			{Score: 5, Member: "f"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZRevRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"f", "e", "d", "c", "b", "a",
		}))
		Expect(c.ZRevRangeWithScores(ctx, k, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 5, Member: "f"},
			{Score: 4, Member: "e"},
			{Score: 3, Member: "d"},
			{Score: 2, Member: "c"},
			{Score: 1, Member: "b"},
			{Score: 0, Member: "a"},
		}))
	})

	It("nokey", func() {
		k := "nokeyzrange"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.ZRevRange(ctx, k, 0, -1).Val()).To(Equal([]string{}))
	})
})
