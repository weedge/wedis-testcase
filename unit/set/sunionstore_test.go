package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sunionstore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})
	It("ok", func() {
		Expect(c.SAdd(ctx, "sunionstores1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sunionstores2", "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SUnionStore(ctx, "sunionstores111", "sunionstores1", "sunionstores2").Val()).To(Equal(int64(5)))
		Expect(c.SMembers(ctx, "sunionstores111").Val()).To(Equal([]string{"m1", "m2", "m4", "m3", "m5"}))
	})

	It("union same", func() {
		Expect(c.SAdd(ctx, "sunionstoresss1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sunionstoresss2", "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SUnionStore(ctx, "sunionstoresss11", "sunionstoresss1", "sunionstoresss2").Val()).To(Equal(int64(3)))
		Expect(c.SMembers(ctx, "sunionstoresss11").Val()).To(Equal([]string{"m1", "m2", "m4"}))
	})
	It("union empty", func() {
		Expect(c.SUnionStore(ctx, "sunionstoresss11", "sunionstoresss12").Val()).To(Equal(int64(0)))
	})
})
