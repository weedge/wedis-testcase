package srv

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("select Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.Do(ctx, "SELECT", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, "select", "test", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "SELECT", 1).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, "select").Err().Error()).To(ContainSubstring("redis: nil"))
		Expect(c.Do(ctx, "SELECT", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, "select").Val()).To(Equal("test"))
	})
})
