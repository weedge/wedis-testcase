package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lkeyexists Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.RPush(ctx, "lkeyexistsl1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "LKEYEXISTS", "lkeyexistsl1").Val()).To(Equal(int64(1)))
	})
	It("no", func() {
		Expect(c.Do(ctx, "LKEYEXISTS", "lkeyexistsk11l").Val()).To(Equal(int64(0)))
	})
})
