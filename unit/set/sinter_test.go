package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sinter Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.SAdd(ctx, "s11", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "s12", "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SInter(ctx, "s11", "s12").Val()).To(Equal([]string{"m2"}))
	})

	It("inter same", func() {
		Expect(c.SAdd(ctx, "sss11", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sss12", "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SInter(ctx, "sss11", "sss12").Val()).To(Equal([]string{"m1", "m2", "m4"}))
	})
})
