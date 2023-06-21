package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MGET Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("mget against non existing key", func() {
		Expect(c.MGet(ctx, "MGETk1000", "MGETk2000").Val()).To(Equal([]interface{}{nil, nil}))
	})
})
