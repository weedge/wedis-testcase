package srv

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var (
	ctx context.Context
	c   redis.UniversalClient
)

var _ = Describe("HelloCmd", func() {
	BeforeEach(func() {
		ctx = context.Background()
		c = gSrv.NewRedisClient()
	})

	AfterEach(func() {
		Expect(c.Close()).NotTo(HaveOccurred())
	})

	It("hello with protocol 2", func() {
		worldList := c.Do(ctx, "HELLO", "2").Val().([]interface{})
		Expect(worldList[2]).To(Equal("proto"))
		Expect(worldList[3]).To(Equal(int64(2)))
	})

	It("hello with protocol 3", func() {
		worldList := c.Do(ctx, "HELLO", "3").Val().([]interface{})
		Expect(worldList[2]).To(Equal("proto"))
		Expect(worldList[3]).To(Equal(int64(2)))
	})

	It("hello with wrong protocol", func() {
		Expect(c.Do(ctx, "HELLO", 1).Err().Error()).To(ContainSubstring("NOPROTO unsupported protocol version"))
		Expect(c.Do(ctx, "HELLO", 888).Err().Error()).To(ContainSubstring("NOPROTO unsupported protocol version"))
	})

	It("hello with non protocol", func() {
		Expect(c.Do(ctx, "HELLO", "WORLD").Err().Error()).To(ContainSubstring("Protocol version is not an integer or out of range"))
	})

	It("hello with setname", func() {
		name := "weDis"
		worldList := c.Do(ctx, "HELLO", "2", "SETNAME", name).Val().([]interface{})
		Expect(worldList[2]).To(Equal("proto"))
		Expect(worldList[3]).To(Equal(int64(2)))
		Expect(c.Do(ctx, "CLIENT", "GETNAME").Val()).To(Equal(name))
	})

	It("hello with auth default user", func() {
		name := "weDis"
		worldList := c.Do(ctx, "HeLLO", "2", "Auth", "deFault", "aaa", "SETnAME", name).Val().([]interface{})
		Expect(worldList[2]).To(Equal("proto"))
		Expect(worldList[3]).To(Equal(int64(2)))
		Expect(c.Do(ctx, "CLIENT", "GETNAME").Val()).To(Equal(name))
	})

	It("hello with auth no default user", func() {
		Expect(c.Do(ctx, "HeLLO", "2", "Auth", "oligei", "aaa", "SETnAME", "weDis").Err().Error()).To(ContainSubstring("ERR invalid password"))
	})

	It("hello with wrong option", func() {
		Expect(c.Do(ctx, "HELLO", "2", "Auth", "deFault").Err().Error()).To(ContainSubstring("ERR syntax error in HELLO option"))
		Expect(c.Do(ctx, "HELLO", "2", "Auth", "deFault", "1", "setname").Err().Error()).To(ContainSubstring("ERR syntax error in HELLO option"))
		Expect(c.Do(ctx, "HELLO", "2", "Auth", "deFault", "1", "setname", "2", "ca").Err().Error()).To(ContainSubstring("ERR syntax error in HELLO option"))
	})

})
