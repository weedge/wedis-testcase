package string

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SETNX Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("setnx key no exists", func() {
		key, val := "SETNXkey", "val"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.SetNX(ctx, key, val, 0).Val()).To(BeTrue())
	})

	It("setnx key exists", func() {
		key, val := "SETNXkey", "val"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.SetNX(ctx, key, val, 0).Val()).To(BeFalse())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})

	It("setnx un-expired key", func() {
		key, val, val1 := "SETNXkey", "val", "val1"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Expire(ctx, key, 10*time.Second).Err()).NotTo(HaveOccurred())
		Expect(c.SetNX(ctx, key, val1, 0).Val()).To(BeFalse())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})
})
