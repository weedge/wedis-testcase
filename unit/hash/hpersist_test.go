package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hpersist Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("no exists key", func() {
		k := "hpersistnokey"
		Expect(c.Do(ctx, "HPERSIST", k).Val()).To(Equal(int64(0)))
	})

	It("ok", func() {
		k, f, v := "hpersist", "field", "val"
		Expect(c.HSet(ctx, k, f, v).Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "HEXPIRE", k, "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "HTTL", k).Val()).To(Equal(int64(100)))
		Expect(c.Do(ctx, "HPERSIST", k).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "HTTL", k).Val()).To(Equal(int64(-1)))
	})
})
