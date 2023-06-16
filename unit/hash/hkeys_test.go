package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hkeys Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.HSet(ctx, "k1", "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, "k1", "f2", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HKeys(ctx, "k1").Val()).To(Equal([]string{"f1", "f2"}))
	})

	It("echo other", func() {
		Expect(c.HSet(ctx, "k11", "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, "k22", "f2", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HKeys(ctx, "k11").Val()).To(Equal([]string{"f1"}))
		Expect(c.HKeys(ctx, "k22").Val()).To(Equal([]string{"f2"}))
	})
})
