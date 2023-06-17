package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sdiffstore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.SAdd(ctx, "s1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "s2", "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SDiffStore(ctx, "ss1", "s1", "s2").Val()).To(Equal(int64(2)))
		Expect(c.SMembers(ctx, "ss1").Val()).To(Equal([]string{"m1", "m4"}))
	})

	It("diff same", func() {
		Expect(c.SAdd(ctx, "sss1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sss2", "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SDiffStore(ctx, "sss3", "sss1", "sss2").Val()).To(Equal(int64(0)))
		Expect(c.SMembers(ctx, "sss3").Val()).To(Equal([]string{}))
	})
})
