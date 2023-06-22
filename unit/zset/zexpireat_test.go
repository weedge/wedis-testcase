package zset

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zexpireat Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})
	It("ok", func() {
		t := time.Now().UTC().Unix() + 3
		k := "zexpireat"
		Expect(c.Do(ctx, "zMCLEAR", k).Err()).NotTo(HaveOccurred())
		arrZ := []redis.Z{
			{Score: 1, Member: "m1"},
		}
		Expect(c.ZAdd(ctx, k, arrZ...).Val()).To(Equal(int64(len(arrZ))))
		Expect(c.Do(ctx, "zEXPIREAT", k, t).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "zTTL", k).Val()).To(Equal(int64(3)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "zTTL", k).Val()).To(Equal(int64(2)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "zTTL", k).Val()).To(Equal(int64(1)))
		time.Sleep(2 * time.Second) //1 s intervals check ttl
		Expect(c.Do(ctx, "zTTL", k).Val()).To(Equal(int64(-2)))
		Expect(c.SMembers(ctx, k).Val()).To(Equal([]string{}))
	})

	It("ERR invalid expire value", func() {
		k := "zexpireatinvad"
		Expect(c.Do(ctx, "zEXPIREAT", k, "100").Err().Error()).To(ContainSubstring("ERR invalid expire value"))
	})

	It("ERR value is not an integer or out of range", func() {
		Expect(c.Do(ctx, "zEXPIREAT", "z110", "1a0").Err().Error()).To(ContainSubstring("ERR value is not an integer or out of range"))
	})
})
