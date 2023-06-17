package srv

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("client Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("getname", func() {
		name := "weDis"
		Expect(c.Do(ctx, "HELLO", "2", "SETNAME", name).Err()).NotTo(HaveOccurred())
		Expect(c.ClientGetName(ctx).Val()).To(Equal(name))
	})
})
