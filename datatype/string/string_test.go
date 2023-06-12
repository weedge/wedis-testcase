package string

import (
	"context"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var (
	ctx context.Context
	c   redis.UniversalClient
)

var _ = Describe("StringCmd", func() {
	BeforeEach(func() {
		ctx = context.Background()
		c = gSrv.NewRedisClient()
	})

	AfterEach(func() {
		Expect(c.Close()).NotTo(HaveOccurred())
	})

	It("set get", func() {
		key, val := "key", "val"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})

	It("set get empty val", func() {
		key, val := "key", ""
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})

	It("set get big val", func() {
		key, val := "key", strings.Repeat("oligei", 1000_000)
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})

	It("setnx key no exists", func() {
		key, val := "key", "val"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.SetNX(ctx, key, val, 0).Val()).To(BeTrue())
	})

	It("setnx key exists", func() {
		key, val := "key", "val"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.SetNX(ctx, key, val, 0).Val()).To(BeFalse())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})

	It("setnx un-expired key", func() {
		key, val, val1 := "key", "val", "val1"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Expire(ctx, key, 10*time.Second).Err()).NotTo(HaveOccurred())
		Expect(c.SetNX(ctx, key, val1, 0).Val()).To(BeFalse())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})

})
