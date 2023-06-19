package list

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("brpop Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.RPush(ctx, "brpopl1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.BRPop(ctx, 1*time.Second, "brpopl1").Val()).To(Equal([]string{"brpopl1", "v3"}))
		Expect(c.BRPop(ctx, 1*time.Second, "brpopl1").Val()).To(Equal([]string{"brpopl1", "v2"}))
		Expect(c.BRPop(ctx, 1*time.Second, "brpopl1").Val()).To(Equal([]string{"brpopl1", "v1"}))
		Expect(c.BRPop(ctx, 1*time.Second, "brpopl1").Val()).To(Equal([]string{}))
	})
	It("block no exists", func() {
		s := time.Now()
		Expect(c.BRPop(ctx, 2*time.Second, "brpopl110").Val()).To(Equal([]string{}))
		Expect(time.Since(s).Seconds() > 2).To(BeTrue())
	})
})
