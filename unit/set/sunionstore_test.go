package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sunionstore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})
	It("ok", func() {
		k1, k2, k3 := "sunionstores1", "sunionstores2", "sunionstores3"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k1, "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, k2, "m2", "m3", "m5").Val()).To(Equal(int64(3)))
		Expect(c.SUnionStore(ctx, k3, k1, k2).Val()).To(Equal(int64(5)))
		Expect(c.SMembers(ctx, k3).Val()).To(Equal([]string{"m1", "m2", "m3", "m4", "m5"})) // sort by lsm tree in memtable
	})

	It("union same", func() {
		k1, k2, k3 := "sunionstoress1", "sunionstoress2", "sunionstoress3"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k1, "m1", "m2", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SAdd(ctx, k2, "m2", "m1", "m4").Val()).To(Equal(int64(3)))
		Expect(c.SUnionStore(ctx, k3, k1, k2).Val()).To(Equal(int64(3)))
		Expect(c.SMembers(ctx, k3).Val()).To(Equal([]string{"m1", "m2", "m4"}))
	})
	It("union empty", func() {
		k1, k2 := "sunionstoresss11", "sunionstoresss12"
		Expect(c.Do(ctx, "SMCLEAR", k1, k2).Err()).NotTo(HaveOccurred())
		Expect(c.SUnionStore(ctx, k1, k2).Val()).To(Equal(int64(0)))
	})
})
