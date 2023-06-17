package string

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("expire Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "keyexists"
		Expect(c.Set(ctx, k, "m1", 0).Val()).To(Equal("OK"))
		Expect(c.Expire(ctx, k, 100*time.Second).Val()).To(BeTrue())
		Expect(c.TTL(ctx, k).Val()).To(Equal(100 * time.Second))
		time.Sleep(1 * time.Second)
		Expect(c.TTL(ctx, k).Val()).To(Equal(99 * time.Second))
	})

	It("no", func() {
		Expect(c.Expire(ctx, "nokeyex", 100*time.Second).Val()).To(BeFalse())
	})
})
