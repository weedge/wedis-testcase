package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("smclear Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "smclear"
		Expect(c.SAdd(ctx, k, "m1", "m2").Val()).To(Equal(int64(2)))
		Expect(c.Do(ctx, "SMCLEAR", k).Val()).To(Equal(int64(1)))
		Expect(c.SMembers(ctx, k).Val()).To(Equal([]string{}))
	})

	It("smclear empty", func() {
		Expect(c.Do(ctx, "sMCLEAR", "ss11").Val()).To(Equal(int64(0)))
	})
})
