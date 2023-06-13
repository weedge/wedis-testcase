package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("APPEND Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k1, v1 := "k1", "v1"
		Expect(c.Del(ctx, k1).Err()).NotTo(HaveOccurred())
		Expect(c.Append(ctx, k1, v1).Val()).To(Equal(int64(len(v1))))
		Expect(c.Append(ctx, k1, v1).Val()).To(Equal(int64(2 * len(v1))))
	})
})
