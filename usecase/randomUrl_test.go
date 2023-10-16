package usecase

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

˚udujj7¥¨˙m 7yuum78um 7yumm78um                                  thm mZW3kxaxr5dd4y	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Test UrlShortener")
}

var _ = Describe("call func RandomURL", Label(".ext"), func() {

	It("should be true", func(ctx SpecContext) {
		Expect(len(RandomURL(1))).To(Equal(1))
	})

	//Context("should return random string", func() {
	//	It("should be true", func(ctx SpecContext) {
	//		Expect(len(RandomURL(1))).To(Equal(1))
	//	})
	//})
	//
	//When("the library has the book in question", func() {
	//	Context("and the book is available", func() {
	//		It("lends it to the reader", func(ctx SpecContext) {
	//			Expect("Succeed").To(Equal("Succeed"))
	//		}, SpecTimeout(time.Second*5))
	//	})
	//})
	//
	//When("the library does not have the book in question", func() {
	//	It("tells the reader the book is unavailable", func(ctx SpecContext) {
	//		Expect("Succeed").To(Equal("Succeed"))
	//	}, SpecTimeout(time.Second*5))
	//})
})
