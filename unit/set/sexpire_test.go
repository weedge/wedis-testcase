package set

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sexpire Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})
	It("ok", func() {
		Expect(c.Do(ctx, "SMCLEAR", "sexpires1").Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, "sexpires1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "SEXPIRE", "sexpires1", "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "STTL", "sexpires1").Val()).To(Equal(int64(100)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "STTL", "sexpires1").Val()).To(Equal(int64(99)))
	})

	It("no", func() {
		Expect(c.Do(ctx, "SMCLEAR", "sexpireSnokey").Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "SEXPIRE", "sexpireSnokey", "100").Val()).To(Equal(int64(0)))
	})
})
