package list

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("brpoplpush Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.RPush(ctx, "l1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.BRPopLPush(ctx, "l1", "l2", 1*time.Second).Val()).To(Equal("v3"))
		Expect(c.LRange(ctx, "l1", 0, -1).Val()).To(Equal([]string{"v1", "v2"}))
		Expect(c.LRange(ctx, "l2", 0, -1).Val()).To(Equal([]string{"v3"}))
	})

	It("no exists", func() {
		Expect(c.BRPopLPush(ctx, "ll1", "ll2", 1*time.Second).Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
