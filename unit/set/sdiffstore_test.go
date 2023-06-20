package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sdiffstore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k1, k2, k3 := "sdiffstores1", "sdiffstores2", "sdiffstores3"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k1, "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, k2, "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SDiffStore(ctx, k3, k1, k2).Val()).To(Equal(int64(2)))
		Expect(c.SMembers(ctx, k3).Val()).To(Equal([]string{"m1", "m4"}))
	})

	It("diff same", func() {
		k1, k2, k3 := "sdiffstoress1", "sdiffstoress2", "sdiffstoress3"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k1, "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, k2, "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SDiffStore(ctx, k3, k1, k2).Val()).To(Equal(int64(0)))
		Expect(c.SMembers(ctx, k3).Val()).To(Equal([]string{}))
	})
})
