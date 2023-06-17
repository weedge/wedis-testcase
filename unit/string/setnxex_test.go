package string

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("setnxex Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, ex, val := "setnxex", 100, "val"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "setnxex", key, ex, val).Val()).To(Equal(int64(1)))
		Expect(c.TTL(ctx, key).Val()).To(Equal(100 * time.Second))
	})

	It("exists", func() {
		key, ex, val := "key", 100, "val"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, key, val, 0).Val()).To(Equal("OK"))
		Expect(c.Do(ctx, "setnxex", key, ex, val).Val()).To(Equal(int64(0)))
		Expect(c.TTL(ctx, key).Val()).To(Equal(time.Duration(-1)))
	})
})
