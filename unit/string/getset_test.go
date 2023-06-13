package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringCmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("getset new val", func() {
		key, val := "key", "val"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.GetSet(ctx, key, val).Val()).To(Equal(""))
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})
	It("getset replace old val", func() {
		key, val, newVal := "key", "val", "newVal"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.GetSet(ctx, key, newVal).Val()).To(Equal(val))
		Expect(c.Get(ctx, key).Val()).To(Equal(newVal))
	})

})
