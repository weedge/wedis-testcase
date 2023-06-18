package zset

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zexpire Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect("").To(Equal(""))
	})
	It("ok", func() {
		key := "zexpire"
		arrZ := []redis.Z{
			{Score: 1, Member: "m1"},
		}
		Expect(c.ZAdd(ctx, key, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.Do(ctx, "ZEXPIRE", key, "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "ZTTL", key).Val()).To(Equal(int64(100)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "ZTTL", key).Val()).To(Equal(int64(99)))
	})

	It("no", func() {
		Expect(c.Do(ctx, "ZEXPIRE", "Z110", "100").Val()).To(Equal(int64(0)))
	})
})
