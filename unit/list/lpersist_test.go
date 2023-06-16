package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lpersist Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("no exists key", func() {
		Expect(c.Do(ctx, "LPERSIST", "kkk").Val()).To(Equal(int64(0)))
	})

	It("ok", func() {
		k := "l1"
		Expect(c.RPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "LEXPIRE", k, "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(100)))
		Expect(c.Do(ctx, "LPERSIST", k).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(-1)))
	})
})
