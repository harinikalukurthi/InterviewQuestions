package main
import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)
var _ = Describe("Sum function",func(){
	Context("when adding two numbers",func(){
		It("adding two positive numbers",func(){
			result:=Sum(2,3)
			Expect(result).To(Equal(5))
		})
		It("addition of two negative numbers",func(){
			result:=Sum(-2,-3)
			Expect(result).To(Equal(-5))
		})
		It("addition of one positive and negative",func(){
			result:=Sum(-4,2)
			Expect(result).To(Equal(-2))
		})
	})
})