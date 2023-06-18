package zset

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("zinterstore Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key1 := "zinterstore1"
		arrZ1 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key1, arrZ1...).Val()).To(Equal(int64(len(arrZ1))))
		key2 := "zinterstore2"
		arrZ2 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key2, arrZ2...).Val()).To(Equal(int64(len(arrZ2))))
		key3 := "zinterstoreSum"
		Expect(c.ZInterStore(ctx, key3, &redis.ZStore{
			Keys:      []string{key1, key2},
			Weights:   []float64{2, 3},
			Aggregate: "", //sum
		}).Val()).To(Equal(int64(3)))
		Expect(c.ZRangeWithScores(ctx, key3, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 5, Member: "m1"},
			{Score: 10, Member: "m2"},
			{Score: 15, Member: "m3"},
		}))
	})

	It("sum", func() {
		key1 := "zinterstore1"
		arrZ1 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key1, arrZ1...).Val()).To(Equal(int64(len(arrZ1))))
		key2 := "zinterstore2"
		arrZ2 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key2, arrZ2...).Val()).To(Equal(int64(len(arrZ2))))
		key3 := "zinterstoreSum"
		Expect(c.ZInterStore(ctx, key3, &redis.ZStore{
			Keys:      []string{key1, key2},
			Weights:   []float64{2, 3},
			Aggregate: "sum",
		}).Val()).To(Equal(int64(3)))
		Expect(c.ZRangeWithScores(ctx, key3, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 5, Member: "m1"},
			{Score: 10, Member: "m2"},
			{Score: 15, Member: "m3"},
		}))
	})
	It("min", func() {
		key1 := "zinterstore1"
		arrZ1 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key1, arrZ1...).Val()).To(Equal(int64(len(arrZ1))))
		key2 := "zinterstore2"
		arrZ2 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key2, arrZ2...).Val()).To(Equal(int64(len(arrZ2))))
		key3 := "zinterstoreSum"
		Expect(c.ZInterStore(ctx, key3, &redis.ZStore{
			Keys:      []string{key1, key2},
			Weights:   []float64{2, 3},
			Aggregate: "min",
		}).Val()).To(Equal(int64(3)))
		Expect(c.ZRangeWithScores(ctx, key3, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 2, Member: "m1"},
			{Score: 4, Member: "m2"},
			{Score: 6, Member: "m3"},
		}))
	})
	It("max", func() {
		key1 := "zinterstore1"
		arrZ1 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key1, arrZ1...).Val()).To(Equal(int64(len(arrZ1))))
		key2 := "zinterstore2"
		arrZ2 := []redis.Z{
			{Score: 1, Member: "m1"},
			{Score: 2, Member: "m2"},
			{Score: 3, Member: "m3"},
		}
		Expect(c.ZAdd(ctx, key2, arrZ2...).Val()).To(Equal(int64(len(arrZ2))))
		key3 := "zinterstoreSum"
		Expect(c.ZInterStore(ctx, key3, &redis.ZStore{
			Keys:      []string{key1, key2},
			Weights:   []float64{2, 3},
			Aggregate: "max",
		}).Val()).To(Equal(int64(3)))
		Expect(c.ZRangeWithScores(ctx, key3, 0, -1).Val()).To(Equal([]redis.Z{
			{Score: 3, Member: "m1"},
			{Score: 6, Member: "m2"},
			{Score: 9, Member: "m3"},
		}))
	})
})
