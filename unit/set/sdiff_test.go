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
		Expect(c.SAdd(ctx, "s1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "s2", "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SDiff(ctx, "s1", "s2").Val()).To(Equal([]string{"m1", "m4"}))
	})

	It("diff same", func() {
		Expect(c.SAdd(ctx, "ss1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "ss2", "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SDiff(ctx, "ss1", "ss2").Val()).To(Equal([]string{}))
	})
})
