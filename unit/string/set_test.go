package string

import (
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/weedge/wedis-testcase/help"
)

var _ = Describe("Set Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("set get", func() {
		key, val := "Setkey", "val"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})

	It("set get nil val", func() {
		key := "Setkeynil"
		Expect(c.Set(ctx, key, nil, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, key).Val()).To(Equal(""))
	})

	It("set get empty val", func() {
		key, val := "Setkey", ""
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})

	It("set get big val", func() {
		key, val := "Setkey", strings.Repeat("oligei", 1000_000)
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})

	It("set get big random val", func() {
		var payload []string
		for i := 0; i < 100; i++ {
			buf := help.RandString(1, 100000, help.Alpha)
			payload = append(payload, buf)
			Expect(c.Set(ctx, "Setbigpayload_"+strconv.Itoa(i), buf, 0).Err()).NotTo(HaveOccurred())
		}

		for i := 0; i < 1000; i++ {
			index := help.RandomInt(100)
			key := "Setbigpayload_" + strconv.FormatInt(index, 10)
			Expect(c.Get(ctx, key).Val()).To(Equal(payload[index]))
		}
	})

	It("set mget", func() {
		Expect(c.Del(ctx, "setk1", "setk2").Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, "setk1", "v1", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Set(ctx, "setk2", "v2", 0).Err()).NotTo(HaveOccurred())
		Expect(c.MGet(ctx, "setk1", "setk2").Val()).To(Equal([]interface{}{"v1", "v2"}))
	})
})
