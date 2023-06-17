package srv

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("flushall Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.Do(ctx, "SELECT", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, "f1", "t1", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, "f2", "t2", 0).Err()).NotTo(HaveOccurred())

		Expect(c.Do(ctx, "SELECT", 1).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, "ff1", "t1", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, "ff2", "t2", 0).Err()).NotTo(HaveOccurred())

		Expect(c.Do(ctx, "SELECT", 0).Err()).NotTo(HaveOccurred())
		Expect(c.FlushAll(ctx).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, "f1").Err().Error()).To(ContainSubstring("redis: nil"))
		Expect(c.Get(ctx, "f2").Err().Error()).To(ContainSubstring("redis: nil"))

		Expect(c.Do(ctx, "SELECT", 1).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, "ff1").Err().Error()).To(ContainSubstring("redis: nil"))
		Expect(c.Get(ctx, "ff2").Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
