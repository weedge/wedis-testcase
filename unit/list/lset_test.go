package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lset Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "lset"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.LPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.LSet(ctx, k, 0, "vv1").Val()).To(Equal("OK"))
		Expect(c.LRange(ctx, k, 0, -1).Val()).To(Equal([]string{"vv1", "v2", "v1"}))
	})
})
