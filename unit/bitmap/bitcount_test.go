package bitmap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("BITCOUNT Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, val := "key", "foobar"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.BitCount(ctx, key, nil).Val()).To(Equal(int64(26)))
		Expect(c.BitCount(ctx, key, &redis.BitCount{Start: 0, End: 0}).Val()).To(Equal(int64(4)))
		Expect(c.BitCount(ctx, key, &redis.BitCount{Start: 1, End: 1}).Val()).To(Equal(int64(6)))
	})
})
