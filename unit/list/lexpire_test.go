package list

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lexpire Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.RPush(ctx, "lexpirel1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "LEXPIRE", "lexpirel1", "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "LTTL", "lexpirel1").Val()).To(Equal(int64(100)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "LTTL", "lexpirel1").Val()).To(Equal(int64(99)))
	})

	It("no", func() {
		Expect(c.Do(ctx, "LEXPIRE", "lexpirek110", "100").Val()).To(Equal(int64(0)))
	})
})
