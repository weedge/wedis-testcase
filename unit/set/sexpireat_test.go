package set

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sexpireat Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		t := time.Now().UTC().Unix() + 3
		k := "sexpireatssss1"
		Expect(c.Do(ctx, "SMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k, "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "sEXPIREAT", k, t).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "sTTL", k).Val()).To(Equal(int64(3)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "sTTL", k).Val()).To(Equal(int64(2)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "sTTL", k).Val()).To(Equal(int64(1)))
		time.Sleep(2 * time.Second) //1 s intervals check ttl
		Expect(c.Do(ctx, "sTTL", k).Val()).To(Equal(int64(-2)))
		Expect(c.SMembers(ctx, k).Val()).To(Equal([]string{}))
	})

	It("ERR invalid expire value", func() {
		Expect(c.Do(ctx, "sEXPIREAT", "sexpireats110", "100").Err().Error()).To(ContainSubstring("ERR invalid expire value"))
	})

	It("ERR value is not an integer or out of range", func() {
		Expect(c.Do(ctx, "sEXPIREAT", "sexpireats110", "1a0").Err().Error()).To(ContainSubstring("ERR value is not an integer or out of range"))
	})
})
