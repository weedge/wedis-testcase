package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lpush Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.LPush(ctx, "ll1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.RPop(ctx, "ll1").Val()).To(Equal("v1"))
		Expect(c.RPop(ctx, "ll1").Val()).To(Equal("v2"))
		Expect(c.RPop(ctx, "ll1").Val()).To(Equal("v3"))
		Expect(c.RPop(ctx, "ll1").Val()).To(Equal(""))
	})
})
