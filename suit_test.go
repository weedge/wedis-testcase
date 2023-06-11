package wedistestcase

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRESPCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "redis RESP Cmd Suite")
}

var _ = BeforeSuite(func() {

})

var _ = AfterSuite(func() {

})
