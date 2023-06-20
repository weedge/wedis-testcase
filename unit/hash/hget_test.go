package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hget Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "hget"
		Expect(c.HGet(ctx, k, "other").Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
