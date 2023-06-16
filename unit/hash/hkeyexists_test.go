package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hkeyexists Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k, f, v := "key", "field", "val"
		Expect(c.HSet(ctx, k, f, v).Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "HKEYEXISTS", k).Val()).To(Equal(int64(1)))
	})
	It("no", func() {
		Expect(c.Do(ctx, "HKEYEXISTS", "k11l").Val()).To(Equal(int64(0)))
	})
})
