package abi_test

import (
	"bytes"

	"github.com/renproject/abi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bytes", func() {
	Context("when unmarshaling into unallocated bytes", func() {
		It("should allocate the required length and capacity, and copy the data", func() {
			data := []byte{
				0, 0, 0, 16, // Length prefix
				1, 2, 3, 4, 5, 6, 7, 8,
				9, 0, 1, 2, 3, 4, 5, 6, // Data
			}
			b := abi.Bytes{}
			err := b.Unmarshal(bytes.NewReader(data))
			Expect(err).ToNot(HaveOccurred())
			Expect([]byte(b)).To(HaveLen(len(data) - 4))
			Expect([]byte(b)).To(HaveCap(len(data) - 4))
			Expect(bytes.Equal(data[4:], b[:])).To(BeTrue())
		})
	})

	Context("when unmarshaling into 32 bytes", func() {
		It("should copy the data", func() {
			data := [32]byte{
				1, 2, 3, 4, 5, 6, 7, 8, // Data
				9, 0, 1, 2, 3, 4, 5, 6,
				7, 8, 9, 0, 1, 2, 3, 4,
				5, 6, 7, 8, 9, 0, 1, 2,
			}
			b32 := abi.Bytes32{}
			err := b32.Unmarshal(bytes.NewReader(data[:]))
			Expect(err).ToNot(HaveOccurred())
			Expect(b32).To(HaveLen(32))
			Expect(b32).To(HaveCap(32))
			Expect(bytes.Equal(data[:], b32[:])).To(BeTrue())
		})
	})
})
