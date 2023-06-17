package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("skeyexists Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})
	It("ok", func() {
		k := "skeyexists"
		Expect(c.SAdd(ctx, k, "m1", "m2").Val()).To(Equal(int64(2)))
		Expect(c.Do(ctx, "SKEYEXISTS", k).Val()).To(Equal(int64(1)))
	})
	It("no", func() {
		Expect(c.Do(ctx, "SKEYEXISTS", "k11s").Val()).To(Equal(int64(0)))
	})
})
