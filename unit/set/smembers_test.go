package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("smembers Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "smembers"
		Expect(c.SAdd(ctx, k, "m1", "m2").Val()).To(Equal(int64(2)))
		Expect(c.SMembers(ctx, k).Val()).To(Equal([]string{"m1", "m2"}))
	})
	It("nokey empty", func() {
		Expect(c.SMembers(ctx, "nokeys").Val()).To(Equal([]string{}))
	})
})
