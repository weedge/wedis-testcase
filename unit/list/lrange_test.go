package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lrange Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.LPush(ctx, "ll1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.LRange(ctx, "ll1", 0, -1).Val()).To(Equal([]string{"v3", "v2", "v1"}))
		Expect(c.LRange(ctx, "ll1", 1, -1).Val()).To(Equal([]string{"v2", "v1"}))
		Expect(c.LRange(ctx, "ll1", 2, -1).Val()).To(Equal([]string{"v1"}))
		Expect(c.LRange(ctx, "ll1", 3, -1).Val()).To(Equal([]string{}))
	})
	It("no key", func() {
		Expect(c.LRange(ctx, "nokey", 0, -1).Val()).To(Equal([]string{}))
	})
})
