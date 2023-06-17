package string

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("setex Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, val := "setex", "val"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.SetEx(ctx, key, val, 100*time.Second).Val()).To(Equal("OK"))
		Expect(c.TTL(ctx, key).Val()).To(Equal(100 * time.Second))
	})
	It("noexists", func() {
		key, val := "setexnoexists", "val"
		Expect(c.SetEx(ctx, key, val, 100*time.Second).Val()).To(Equal("OK"))
		Expect(c.TTL(ctx, key).Val()).To(Equal(100 * time.Second))
	})
})
