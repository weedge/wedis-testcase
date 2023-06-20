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
		k := "hmsetk1"
		Expect(c.HMSet(ctx, k, "f1", "v1", "f2", "v2").Err()).NotTo(HaveOccurred())
		Expect(c.HMGet(ctx, k, "f1", "f2").Val()).To(Equal([]interface{}{"v1", "v2"}))
	})

	It("err", func() {
		k := "hmsetk2"
		Expect(c.HMSet(ctx, k, "f1", "v1", "f2", "v2", "f3").Err().Error()).To(ContainSubstring("ERR wrong number of arguments"))
	})
})
