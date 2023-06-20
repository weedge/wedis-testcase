package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("scard Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "scardskeyk"
		Expect(c.Do(ctx, "SMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k, "m1", "m2").Val()).To(Equal(int64(2)))
		Expect(c.SCard(ctx, k).Val()).To(Equal(int64(2)))
	})

	It("no exists", func() {
		Expect(c.SCard(ctx, "scardnoskey").Val()).To(Equal(int64(0)))
	})
})
