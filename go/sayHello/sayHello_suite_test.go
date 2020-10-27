package sayHello_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSayHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hello Suite")
}
