package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestExamQuestion1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ExamQuestion1 Suite")
}
