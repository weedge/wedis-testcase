package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zadd Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key := "zaddkey"
		Expect(c.Do(ctx, "zMCLEAR", key).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key, arrZ...).Val()).To(Equal(int64(len(arrZ))))
	})

	It("ERR wrong number of arguments", func() {
		key := "zaddkey"
		Expect(c.ZAdd(ctx, key).Err().Error()).To(ContainSubstring("ERR wrong number of arguments"))
	})

	It("ERR invalid zset member size", func() {
		key := "zaddkey"
		arrZ := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key, arrZ...).Err().Error()).To(ContainSubstring("ERR invalid zset member size"))
	})
})
