package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zlexcount Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zkeyexists"
		arrZ := []redis.Z{
			{Score: 0, Member: "a"},
			{Score: 0, Member: "b"},
			{Score: 0, Member: "c"},
			{Score: 0, Member: "d"},
			{Score: 0, Member: "e"},
			{Score: 0, Member: "f"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))

		Expect(c.ZLexCount(ctx, k, "-", "+").Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZLexCount(ctx, k, "-", "-").Val()).To(Equal(int64(0)))
		Expect(c.ZLexCount(ctx, k, "+", "+").Val()).To(Equal(int64(0)))

		Expect(c.ZLexCount(ctx, k, "[b", "[f").Val()).To(Equal(int64(5)))
		Expect(c.ZLexCount(ctx, k, "[b", "(f").Val()).To(Equal(int64(4)))
		Expect(c.ZLexCount(ctx, k, "(b", "[f").Val()).To(Equal(int64(4)))
		Expect(c.ZLexCount(ctx, k, "(b", "(f").Val()).To(Equal(int64(3)))
	})
})
