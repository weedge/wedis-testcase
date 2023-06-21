package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lkeyexists Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "lkeyexistsl1"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.RPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "LKEYEXISTS", k).Val()).To(Equal(int64(1)))
	})
	It("no", func() {
		k := "lkeyexistsnokey"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "LKEYEXISTS", k).Val()).To(Equal(int64(0)))
	})
})
