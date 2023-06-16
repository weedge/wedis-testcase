package hash

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("httl Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k, f, v := "key", "field", "val"
		Expect(c.HSet(ctx, k, f, v).Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "HEXPIRE", k, "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "HTTL", k).Val()).To(Equal(int64(100)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "HTTL", k).Val()).To(Equal(int64(99)))
		Expect(c.Do(ctx, "HPERSIST", k).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "HTTL", k).Val()).To(Equal(int64(-1)))
	})

	It("no key", func() {
		Expect(c.Do(ctx, "HTTL", "nokey").Val()).To(Equal(int64(-1)))
	})
})
