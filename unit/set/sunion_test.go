package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sunion Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k1, k2 := "sunions1", "sunions2"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k1, "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, k2, "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SUnion(ctx, k1, k2).Val()).To(Equal([]string{"m1", "m2", "m4", "m3", "m5"}))
	})

	It("union same", func() {
		k1, k2 := "sunionss1", "sunionss2"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k1, "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, k2, "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SUnion(ctx, k1, k2).Val()).To(Equal([]string{"m1", "m2", "m4"}))
	})
	It("union empty", func() {
		k1, k2 := "sunionsss1", "sunionsss2"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SUnion(ctx, k1, k2).Val()).To(Equal([]string{}))
	})
})
