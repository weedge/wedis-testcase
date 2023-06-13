package srv

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PingCmd", func() {
	It("Start Server Ping Ok", func() {
		println("start server ping ok")
	})

	It("Start Server,TCP client Ping", func() {
		c := gSrv.NewTCPClient()
		defer func() {
			Expect(c.Close()).NotTo(HaveOccurred())
		}()
		Expect(c.WriteArgs("PING")).NotTo(HaveOccurred())
		c.MustRead("+PONG")
		Expect(c.WriteArgs("PING", "PONG")).NotTo(HaveOccurred())
		c.MustRead("$4")
		c.MustRead("PONG")
		Expect(c.WriteArgs("ping", "hello", "world")).NotTo(HaveOccurred())
		c.MustRead("-ERR wrong number of arguments")
	})
})
