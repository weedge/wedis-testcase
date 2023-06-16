package bitmap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BITOP Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k1, v1 := "key1", "foobar"
		k2, v2 := "key2", "abcdef"
		Expect(c.Set(ctx, k1, v1, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, k2, v2, 0).Err()).NotTo(HaveOccurred())
		Expect(c.BitOpAnd(ctx, "dest", k1, k2).Val()).To(Equal(int64(6)))
		Expect(c.Get(ctx, "dest").Val()).To(Equal("`bc`ab"))
	})
})
