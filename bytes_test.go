package abi_test

import (
	"bytes"
	"crypto/rand"
	"testing/quick"

	"github.com/renproject/abi"
	"github.com/renproject/surge"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bytes", func() {
	Context("when unmarshaling and unmarshaling", func() {
		It("should equal itself", func() {
			f := func(n uint32) bool {
				buf := new(bytes.Buffer)

				n = n % 100000
				random := make([]byte, n)
				_, err := rand.Reader.Read(random)
				Expect(err).ToNot(HaveOccurred())

				y := abi.Bytes(random)
				_, err = y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.Bytes{}
				_, err = z.Unmarshal(buf, surge.MaxBytes)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when unmarshaling and unmarshaling to/from JSON", func() {
		It("should equal itself", func() {
			f := func(n uint32) bool {
				n = n % 100000
				random := make([]byte, n)
				_, err := rand.Reader.Read(random)
				Expect(err).ToNot(HaveOccurred())

				y := abi.Bytes(random)
				data, err := y.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())

				z := abi.Bytes{}
				err = z.UnmarshalJSON(data)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when unmarshaling into unallocated bytes", func() {
		It("should allocate the required length and capacity, and copy the data", func() {
			data := []byte{
				0, 0, 0, 16, // Length prefix
				1, 2, 3, 4, 5, 6, 7, 8,
				9, 0, 1, 2, 3, 4, 5, 6, // Data
			}
			b := abi.Bytes{}
			_, err := b.Unmarshal(bytes.NewReader(data), surge.MaxBytes)
			Expect(err).ToNot(HaveOccurred())
			Expect([]byte(b)).To(HaveLen(len(data) - 4))
			Expect([]byte(b)).To(HaveCap(len(data) - 4))
			Expect(bytes.Equal(data[4:], b[:])).To(BeTrue())
		})
	})
})

var _ = Describe("Bytes32", func() {
	Context("when unmarshaling and unmarshaling", func() {
		It("should equal itself", func() {
			f := func(x [32]byte) bool {
				buf := new(bytes.Buffer)

				y := abi.Bytes32(x)
				_, err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.Bytes32{}
				_, err = z.Unmarshal(buf, surge.MaxBytes)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when unmarshaling and unmarshaling to/from JSON", func() {
		It("should equal itself", func() {
			f := func(x [32]byte) bool {
				y := abi.Bytes32(x)
				data, err := y.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())

				z := abi.Bytes32{}
				err = z.UnmarshalJSON(data)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when unmarshaling into empty bytes", func() {
		It("should copy the data", func() {
			data := [32]byte{
				1, 2, 3, 4, 5, 6, 7, 8, // Data
				9, 0, 1, 2, 3, 4, 5, 6,
				7, 8, 9, 0, 1, 2, 3, 4,
				5, 6, 7, 8, 9, 0, 1, 2,
			}
			b32 := abi.Bytes32{}
			_, err := b32.Unmarshal(bytes.NewReader(data[:]), surge.MaxBytes)
			Expect(err).ToNot(HaveOccurred())
			Expect(b32).To(HaveLen(32))
			Expect(b32).To(HaveCap(32))
			Expect(bytes.Equal(data[:], b32[:])).To(BeTrue())
		})
	})
})
