package hash

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hexpireat Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		t := time.Now().UTC().Unix() + 3
		k, f, v := "key", "field", "val"
		Expect(c.HSet(ctx, k, f, v).Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "HEXPIREAT", k, t).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "HTTL", k).Val()).To(Equal(int64(3)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "HTTL", k).Val()).To(Equal(int64(2)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "HTTL", k).Val()).To(Equal(int64(1)))
		time.Sleep(2 * time.Second) //1 s intervals check ttl
		Expect(c.Do(ctx, "HTTL", k).Val()).To(Equal(int64(-2)))
		time.Sleep(1 * time.Second)
		Expect(c.HGet(ctx, k, f).Err().Error()).To(ContainSubstring("redis: nil"))
	})

	It("ERR invalid expire value", func() {
		Expect(c.Do(ctx, "HEXPIREAT", "k110", "100").Err().Error()).To(ContainSubstring("ERR invalid expire value"))
	})

	It("ERR value is not an integer or out of range", func() {
		Expect(c.Do(ctx, "HEXPIREAT", "k110", "1a0").Err().Error()).To(ContainSubstring("ERR value is not an integer or out of range"))
	})
})
