package zset

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zttl Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zttl"
		arrZ := []redis.Z{
			{Score: 1, Member: "a"},
			{Score: 3, Member: "b"},
			{Score: 2, Member: "c"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.Do(ctx, "ZEXPIRE", k, "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "ZTTL", k).Val()).To(Equal(int64(100)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "ZTTL", k).Val()).To(Equal(int64(99)))
		Expect(c.Do(ctx, "ZPERSIST", k).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "ZTTL", k).Val()).To(Equal(int64(-1)))
	})

	It("no key", func() {
		Expect(c.Do(ctx, "ZTTL", "nozttlkey").Val()).To(Equal(int64(-2)))
	})
})
