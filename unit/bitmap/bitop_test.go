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
		k1, v1 := "BITOPkey1", "foobar"
		k2, v2 := "BITOPkey2", "abcdef"
		Expect(c.Set(ctx, k1, v1, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, k2, v2, 0).Err()).NotTo(HaveOccurred())
		Expect(c.BitOpAnd(ctx, "BITOPdest", k1, k2).Val()).To(Equal(int64(6)))
		Expect(c.Get(ctx, "BITOPdest").Val()).To(Equal("`bc`ab"))
	})
})
