package srv

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("echo Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.Echo(ctx, "hi").Val()).To(Equal("hi"))
		Expect(c.Echo(ctx, " ").Val()).To(Equal(" "))
		Expect(c.Echo(ctx, "").Val()).To(Equal(""))
	})
})
