package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sinterstore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.SAdd(ctx, "ss11", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "ss12", "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SInterStore(ctx, "ss13", "ss11", "ss12").Val()).To(Equal(int64(1)))
		Expect(c.SMembers(ctx, "ss13").Val()).To(Equal([]string{"m2"}))
	})

	It("inter same", func() {
		Expect(c.SAdd(ctx, "sss111", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sss122", "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SInterStore(ctx, "ssss1", "sss111", "sss122").Val()).To(Equal(int64(3)))
		Expect(c.SMembers(ctx, "ssss1").Val()).To(Equal([]string{"m1", "m2", "m4"}))
	})
})
