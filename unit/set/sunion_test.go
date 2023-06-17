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
		Expect(c.SAdd(ctx, "sunions1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sunions2", "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SUnion(ctx, "sunions1", "sunions2").Val()).To(Equal([]string{"m1", "m2", "m4", "m3", "m5"}))
	})

	It("union same", func() {
		Expect(c.SAdd(ctx, "sunions1ss1", "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, "sunions1ss2", "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SUnion(ctx, "sunions1ss1", "sunions1ss2").Val()).To(Equal([]string{"m1", "m2", "m4"}))
	})
	It("union empty", func() {
		Expect(c.SUnion(ctx, "sunions1ss11", "sunions1ss12").Val()).To(Equal([]string{}))
	})
})
