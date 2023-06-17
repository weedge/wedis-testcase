package string

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ttl Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k, val := "ttlkey", "val"
		Expect(c.Set(ctx, k, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Expire(ctx, k, 100*time.Second).Val()).To(BeTrue())
		Expect(c.TTL(ctx, k).Val()).To(Equal(100 * time.Second))
		Expect(c.Persist(ctx, k).Val()).To(BeTrue())
		Expect(c.TTL(ctx, k).Val()).To(Equal(time.Duration(-1)))
	})

	It("no key", func() {
		Expect(c.TTL(ctx, "nokeyttl").Val()).To(Equal(time.Duration(-2)))
	})
})
