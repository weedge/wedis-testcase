package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("srem Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "srem"
		Expect(c.Do(ctx, "SMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k, "m1", "m2").Val()).To(Equal(int64(2)))
		Expect(c.SRem(ctx, k, "m1").Val()).To(Equal(int64(1)))
		Expect(c.SRem(ctx, k, "m1").Val()).To(Equal(int64(0)))
		Expect(c.SRem(ctx, k, "m2").Val()).To(Equal(int64(1)))
		Expect(c.SRem(ctx, k, "m2").Val()).To(Equal(int64(0)))
		Expect(c.SRem(ctx, k, "m3").Val()).To(Equal(int64(0)))
		Expect(c.SMembers(ctx, k).Val()).To(Equal([]string{}))
	})
	It("no exists key", func() {
		k := "sremnokey"
		Expect(c.Do(ctx, "SMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.SRem(ctx, k).Val()).To(Equal(int64(0)))
	})
})
