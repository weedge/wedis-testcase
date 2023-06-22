package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zcount Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key := "zcount"
		Expect(c.Do(ctx, "zMCLEAR", key).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.ZCount(ctx, key, "1", "3").Val()).To(Equal(int64(3)))
		Expect(c.ZCount(ctx, key, "v1", "v2").Err().Error()).To(ContainSubstring("ERR value is not an integer or out of range"))
	})
})
