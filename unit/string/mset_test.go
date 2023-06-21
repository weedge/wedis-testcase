package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MSET Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("mset mget", func() {
		Expect(c.MSet(ctx, map[string]interface{}{
			"MSETk1": "v1",
			"MSETk2": "2",
			"MSETk3": "v3 \n\r\n\r",
		}).Err()).NotTo(HaveOccurred())
		Expect(c.MGet(ctx, "MSETk1", "MSETk2", "MSETk3").Val()).To(Equal([]interface{}{"v1", "2", "v3 \n\r\n\r"}))
	})

	It("mset wrong number of args", func() {
		Expect(c.MSet(ctx, "x", "10", "y", "foo bar", "z").Err().Error()).To(ContainSubstring("wrong number of argument"))
	})
})
