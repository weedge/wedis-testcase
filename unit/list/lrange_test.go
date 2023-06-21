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
		k := "lrangell1"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.LPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.LRange(ctx, k, 0, -1).Val()).To(Equal([]string{"v3", "v2", "v1"}))
		Expect(c.LRange(ctx, k, 1, -1).Val()).To(Equal([]string{"v2", "v1"}))
		Expect(c.LRange(ctx, k, 2, -1).Val()).To(Equal([]string{"v1"}))
		Expect(c.LRange(ctx, k, 3, -1).Val()).To(Equal([]string{}))
	})
	It("no key", func() {
		k := "lrangenokey"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.LRange(ctx, "lrangenokey", 0, -1).Val()).To(Equal([]string{}))
	})
})
