package list

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lexpireat Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		t := time.Now().UTC().Unix() + 3
		k := "lexpireatl1"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.RPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "LEXPIREAT", k, t).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(3)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(2)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(1)))
		time.Sleep(2 * time.Second) //1 s intervals check ttl
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(-2)))
		Expect(c.LRange(ctx, k, 0, -1).Val()).To(Equal([]string{}))
	})

	It("ERR invalid expire value", func() {
		Expect(c.Do(ctx, "LEXPIREAT", "lexpireatk110", "100").Err().Error()).To(ContainSubstring("ERR invalid expire value"))
	})

	It("ERR value is not an integer or out of range", func() {
		Expect(c.Do(ctx, "LEXPIREAT", "lexpireatk110", "1a0").Err().Error()).To(ContainSubstring("ERR value is not an integer or out of range"))
	})
})
