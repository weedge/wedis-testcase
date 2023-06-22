package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zunionstore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key1 := "zunionstore1"
		Expect(c.Do(ctx, "zMCLEAR", key1).Err()).NotTo(HaveOccurred())
		arrZ1 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key1, arrZ1...).Val()).To(Equal(int64(len(arrZ1))))
		key2 := "zunionstore2"
		Expect(c.Do(ctx, "zMCLEAR", key2).Err()).NotTo(HaveOccurred())
		arrZ2 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
			{Score: 4, Member: "m4"},
		}
		Expect(c.ZAdd(ctx, key2, arrZ2...).Val()).To(Equal(int64(len(arrZ2))))
		key3 := "zunionstoreSum"
		Expect(c.Do(ctx, "zMCLEAR", key3).Err()).NotTo(HaveOccurred())
		Expect(c.ZUnionStore(ctx, key3, &redis.ZStore{
			Keys:      []string{key1, key2},
			Weights:   []float64{2, 3},
			Aggregate: "", //sum
		}).Val()).To(Equal(int64(4)))
		Expect(c.ZRangeWithScores(ctx, key3, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 5, Member: "m1"},
			{Score: 10, Member: "m2"},
			{Score: 12, Member: "m4"},
			{Score: 15, Member: "m3"},
		}))
	})
	It("sum", func() {
		key1 := "zunionstoresum1"
		Expect(c.Do(ctx, "zMCLEAR", key1).Err()).NotTo(HaveOccurred())
		arrZ1 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key1, arrZ1...).Val()).To(Equal(int64(len(arrZ1))))
		key2 := "zunionstoresum2"
		Expect(c.Do(ctx, "zMCLEAR", key2).Err()).NotTo(HaveOccurred())
		arrZ2 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
			{Score: 4, Member: "m4"},
		}
		Expect(c.ZAdd(ctx, key2, arrZ2...).Val()).To(Equal(int64(len(arrZ2))))
		key3 := "zunionstoreSum"
		Expect(c.Do(ctx, "zMCLEAR", key3).Err()).NotTo(HaveOccurred())
		Expect(c.ZUnionStore(ctx, key3, &redis.ZStore{
			Keys:      []string{key1, key2},
			Weights:   []float64{2, 3},
			Aggregate: "sum",
		}).Val()).To(Equal(int64(4)))
		Expect(c.ZRangeWithScores(ctx, key3, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 5, Member: "m1"},
			{Score: 10, Member: "m2"},
			{Score: 12, Member: "m4"},
			{Score: 15, Member: "m3"},
		}))
	})
	It("min", func() {
		key1 := "zunionstoremin1"
		Expect(c.Do(ctx, "zMCLEAR", key1).Err()).NotTo(HaveOccurred())
		arrZ1 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key1, arrZ1...).Val()).To(Equal(int64(len(arrZ1))))
		key2 := "zunionstoremin2"
		Expect(c.Do(ctx, "zMCLEAR", key2).Err()).NotTo(HaveOccurred())
		arrZ2 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
			{Score: 4, Member: "m4"},
		}
		Expect(c.ZAdd(ctx, key2, arrZ2...).Val()).To(Equal(int64(len(arrZ2))))
		key3 := "zunionstoreMin"
		Expect(c.Do(ctx, "zMCLEAR", key3).Err()).NotTo(HaveOccurred())
		Expect(c.ZUnionStore(ctx, key3, &redis.ZStore{
			Keys:      []string{key1, key2},
			Weights:   []float64{2, 3},
			Aggregate: "min",
		}).Val()).To(Equal(int64(4)))
		Expect(c.ZRangeWithScores(ctx, key3, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 2, Member: "m1"},
			{Score: 4, Member: "m2"},
			{Score: 6, Member: "m3"},
			{Score: 12, Member: "m4"},
		}))
	})
	It("max", func() {
		key1 := "zunionstoremax1"
		Expect(c.Do(ctx, "zMCLEAR", key1).Err()).NotTo(HaveOccurred())
		arrZ1 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key1, arrZ1...).Val()).To(Equal(int64(len(arrZ1))))
		key2 := "zunionstoremax2"
		Expect(c.Do(ctx, "zMCLEAR", key2).Err()).NotTo(HaveOccurred())
		arrZ2 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
			{Score: 4, Member: "m4"},
		}
		Expect(c.ZAdd(ctx, key2, arrZ2...).Val()).To(Equal(int64(len(arrZ2))))
		key3 := "zunionstoreMax"
		Expect(c.Do(ctx, "zMCLEAR", key3).Err()).NotTo(HaveOccurred())
		Expect(c.ZUnionStore(ctx, key3, &redis.ZStore{
			Keys:      []string{key1, key2},
			Weights:   []float64{2, 3},
			Aggregate: "max",
		}).Val()).To(Equal(int64(4)))
		Expect(c.ZRangeWithScores(ctx, key3, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 3, Member: "m1"},
			{Score: 6, Member: "m2"},
			{Score: 9, Member: "m3"},
			{Score: 12, Member: "m4"},
		}))
	})
})
