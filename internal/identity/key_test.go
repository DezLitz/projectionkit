package identity_test

import (
	. "github.com/dezlitz/projectionkit/internal/identity"
	"github.com/dogmatiq/dogma"
	. "github.com/dogmatiq/enginekit/enginetest/stubs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Key()", func() {
	It("returns the identity key of the handler", func() {
		h := &ProjectionMessageHandlerStub{
			ConfigureFunc: func(c dogma.ProjectionConfigurer) {
				c.Identity("<name>", "<key>")
				c.Routes(
					dogma.HandlesEvent[EventStub[TypeA]](),
				)
			},
		}

		Expect(Key(h)).To(Equal("<key>"))
	})
})
