package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("STRLEN Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("strlen", func() {
		Expect("").To(Equal(""))
	})
})
