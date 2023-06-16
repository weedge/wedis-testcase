package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hmset Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.HMSet(ctx, "k1", "f1", "v1", "f2", "v2").Err()).NotTo(HaveOccurred())
	})

	It("err", func() {
		Expect(c.HMSet(ctx, "k1", "f1", "v1", "f2", "v2", "f3").Err().Error()).To(ContainSubstring("ERR wrong number of arguments"))
	})
})
