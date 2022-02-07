package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestProductService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Servicio de pruebas Suite")
}
