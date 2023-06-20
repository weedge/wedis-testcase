package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sdiff Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.SAdd(ctx, "sdiffs1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sdiffs2", "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SDiff(ctx, "sdiffs1", "sdiffs2").Val()).To(Equal([]string{"m1", "m4"}))
	})

	It("diff same", func() {
		Expect(c.SAdd(ctx, "sdiffss1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sdiffss2", "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SDiff(ctx, "sdiffss1", "sdiffss2").Val()).To(Equal([]string{}))
	})
})
