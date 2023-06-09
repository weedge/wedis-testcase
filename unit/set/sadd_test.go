package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sadd Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		skey := "sadd"
		Expect(c.Do(ctx, "SMCLEAR", skey).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, skey, "m1", "m2").Val()).To(Equal(int64(2)))
		Expect(c.SAdd(ctx, skey, "m1", "m2").Val()).To(Equal(int64(0)))
	})
})
