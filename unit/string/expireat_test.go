package string

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("expireat Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		t := time.Now().UTC().Add(3 * time.Second)
		k := "ssss1"
		Expect(c.Set(ctx, k, "m1", 0).Val()).To(Equal("OK"))
		Expect(c.ExpireAt(ctx, k, t).Val()).To(BeTrue())
		Expect(c.TTL(ctx, k).Val()).To(Equal(3 * time.Second))
		time.Sleep(1 * time.Second)
		Expect(c.TTL(ctx, k).Val()).To(Equal(2 * time.Second))
		time.Sleep(1 * time.Second)
		Expect(c.TTL(ctx, k).Val()).To(Equal(1 * time.Second))
		time.Sleep(2 * time.Second) //1 s intervals check ttl
		//Expect(c.Do(ctx, "TTL", k).Val()).To(Equal(int64(-2)))
		Expect(c.TTL(ctx, k).Val()).To(Equal(time.Duration(-2)))
		Expect(c.Get(ctx, k).Val()).To(Equal(""))
	})

	It("ERR invalid expire value", func() {
		Expect(c.ExpireAt(ctx, "s110", time.Unix(100, 0)).Err().Error()).To(ContainSubstring("ERR invalid expire value"))
	})

	It("ERR value is not an integer or out of range", func() {
		Expect(c.Do(ctx, "EXPIREAT", "s110", "1a0").Err().Error()).To(ContainSubstring("ERR value is not an integer or out of range"))
	})
})
