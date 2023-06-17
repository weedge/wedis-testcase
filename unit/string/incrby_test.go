package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("INCRBY Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, val := "INCRBY", "1"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.IncrBy(ctx, key, 10).Val()).To(Equal(int64(11)))
		Expect(c.IncrBy(ctx, key, -3).Val()).To(Equal(int64(8)))
	})

	It("nokey", func() {
		Expect(c.IncrBy(ctx, "incrbynokey", -10).Val()).To(Equal(int64(-10)))
		Expect(c.Get(ctx, "incrbynokey").Val()).To(Equal("-10"))
	})
})
