package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("setrange Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, val := "setrangekey", "Hello World"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, key, val, 0).Val()).To(Equal("OK"))
		Expect(c.SetRange(ctx, key, 6, "Wedis").Val()).To(Equal(int64(11)))
		Expect(c.Get(ctx, key).Val()).To(Equal("Hello Wedis"))
	})

	It("zero padding", func() {
		key := "setrangezeropadding"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.SetRange(ctx, key, 6, "Wedis").Val()).To(Equal(int64(11)))
		Expect(c.Get(ctx, key).Val()).To(Equal("\x00\x00\x00\x00\x00\x00Wedis"))
	})
})
