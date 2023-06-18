package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zremrangebylex Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zremrangebylex"
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
		Expect(c.ZRemRangeByLex(ctx, k, "[alpha", "[omega").Val()).To(Equal(int64(4)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{
			"ALPHA", "a", "zap", "zip",
		}))
	})

	It("nokey", func() {
		k := "zremrangebylexnokey"
		Expect(c.ZRemRangeByLex(ctx, k, "a", "b").Val()).To(Equal(int64(0)))
	})
})
