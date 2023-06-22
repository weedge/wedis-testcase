package zset

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zmclear Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "zmclear"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: float64(time.Now().Unix()), Member: "v1"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.Do(ctx, "zMCLEAR", k).Val()).To(Equal(int64(1)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{}))
	})

	It("empty", func() {
		k := "zmclearEmpty"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.Do(ctx, "zMCLEAR", k).Val()).To(Equal(int64(0)))
		Expect(c.ZRange(ctx, k, 0, -1).Val()).To(Equal([]string{}))
	})

	It("zmclear onkey", func() {
		Expect(c.Do(ctx, "zMCLEAR", "zzs11Nokey").Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "zMCLEAR", "zzs11Nokey").Val()).To(Equal(int64(0)))
	})
})
