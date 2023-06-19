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
		Expect(c.SAdd(ctx, "sdiffstores1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sdiffstores2", "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SDiffStore(ctx, "sdiffstoress1", "sdiffstores1", "sdiffstores2").Val()).To(Equal(int64(2)))
		Expect(c.SMembers(ctx, "sdiffstoress1").Val()).To(Equal([]string{"m1", "m4"}))
	})

	It("diff same", func() {
		Expect(c.SAdd(ctx, "sdiffstoresss1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sdiffstoresss2", "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SDiffStore(ctx, "sdiffstoresss3", "sdiffstoresss1", "sdiffstoresss2").Val()).To(Equal(int64(0)))
		Expect(c.SMembers(ctx, "sdiffstoresss3").Val()).To(Equal([]string{}))
	})
})
