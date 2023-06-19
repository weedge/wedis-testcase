package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lmclear Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "lmclearl1"
		Expect(c.RPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "LMCLEAR", k).Val()).To(Equal(int64(1)))
		Expect(c.LRange(ctx, k, 0, -1).Val()).To(Equal([]string{}))
	})

	It("lmclear empty", func() {
		Expect(c.Do(ctx, "LMCLEAR", "lmcleark11").Val()).To(Equal(int64(0)))
	})
})
