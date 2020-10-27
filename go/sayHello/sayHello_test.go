package sayHello_test

import (
	. "github.com/fejnartal/code-sandbox/go/sayHello"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HelloSayer", func() {
	It("Returns VMware greetings", func() {
		helloMsg := SayHello("VMware")
		Expect(helloMsg).To(Equal("Hello VMware!"))
	})
	It("Returns hello pizza message", func() {
		helloMsg := SayHello("Bosh")
		Expect(helloMsg).To(Equal("Hello Bosh!"))
	})
})
