package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hgetall Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "hgetallkey"
		Expect(c.HSet(ctx, k, "f1", "v1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, k, "f2", "v2").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, k, "f3", "v3").Err()).NotTo(HaveOccurred())
		Expect(c.HGetAll(ctx, k).Val()).To(Equal(map[string]string{"f1": "v1", "f2": "v2", "f3": "v3"}))
	})

	It("no", func() {
		Expect(c.HGetAll(ctx, "hgetallk110").Val()).To(Equal(map[string]string{}))
	})
})
