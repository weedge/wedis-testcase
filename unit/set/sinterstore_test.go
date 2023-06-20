package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sinterstore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k1, k2, k3 := "sinterstores1", "sinterstores2", "sinterstores3"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k1, "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, k2, "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SInterStore(ctx, k3, k1, k2).Val()).To(Equal(int64(1)))
		Expect(c.SMembers(ctx, k3).Val()).To(Equal([]string{"m2"}))
	})

	It("inter same", func() {
		k1, k2, k3 := "sinterstoress1", "sinterstoress2", "sinterstoress3"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k1, "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, k2, "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SInterStore(ctx, k3, k1, k2).Val()).To(Equal(int64(3)))
		Expect(c.SMembers(ctx, k3).Val()).To(Equal([]string{"m1", "m2", "m4"}))
	})
})
