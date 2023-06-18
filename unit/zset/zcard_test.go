package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zcard Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key := "zcard"
		arrZ := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZCard(ctx, key).Val()).To(Equal(int64(3)))
	})

	It("nokey", func() {
		key := "zcardnokey"
		Expect(c.ZCard(ctx, key).Val()).To(Equal(int64(0)))
	})
})
