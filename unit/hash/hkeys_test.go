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
		Expect(c.HSet(ctx, "hkeysk1", "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, "hkeysk1", "f2", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HKeys(ctx, "hkeysk1").Val()).To(Equal([]string{"f1", "f2"}))
	})

	It("echo other", func() {
		Expect(c.HSet(ctx, "hkeysk11", "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, "hkeysk22", "f2", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HKeys(ctx, "hkeysk11").Val()).To(Equal([]string{"f1"}))
		Expect(c.HKeys(ctx, "hkeysk22").Val()).To(Equal([]string{"f2"}))
	})
})
