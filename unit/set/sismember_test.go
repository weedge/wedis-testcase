package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sismember Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "sismember"
		Expect(c.Do(ctx, "SMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k, "m1", "m2").Val()).To(Equal(int64(2)))
		Expect(c.SIsMember(ctx, k, "m1").Val()).To(BeTrue())
		Expect(c.SIsMember(ctx, k, "m2").Val()).To(BeTrue())
		Expect(c.SIsMember(ctx, k, "m3").Val()).To(BeFalse())
	})
	It("nokey", func() {
		k := "sismembernoskey"
		Expect(c.Do(ctx, "SMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.SIsMember(ctx, k, "m1").Val()).To(BeFalse())
	})
})
