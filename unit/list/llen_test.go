package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("llen Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.RPush(ctx, "llenl1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.LLen(ctx, "llenl1").Val()).To(Equal(int64(3)))
	})

	It("no", func() {
		Expect(c.LLen(ctx, "llenl111").Val()).To(Equal(int64(0)))
	})
})
