package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("exists Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect("").To(Equal(""))
	})
	It("ok", func() {
		k := "keyexists"
		Expect(c.Set(ctx, k, "m1", 0).Val()).To(Equal("OK"))
		Expect(c.Exists(ctx, k).Val()).To(Equal(int64(1)))
	})
	It("no", func() {
		Expect(c.Exists(ctx, "kk110").Val()).To(Equal(int64(0)))
	})
})
