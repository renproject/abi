package abi_test

import (
	"bytes"
	"testing/quick"

	"github.com/renproject/abi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("8-bit unsigned integer", func() {
	Context("when marshaling and unmarshaling", func() {
		It("should equal itself", func() {
			f := func(x uint8) bool {
				buf := new(bytes.Buffer)

				y := abi.NewU8(x)
				err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U8{}
				err = z.Unmarshal(buf)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var _ = Describe("16-bit unsigned integer", func() {
	Context("when marshaling and unmarshaling", func() {
		It("should equal itself", func() {
			f := func(x uint16) bool {
				buf := new(bytes.Buffer)

				y := abi.NewU16(x)
				err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U16{}
				err = z.Unmarshal(buf)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var _ = Describe("32-bit unsigned integer", func() {
	Context("when marshaling and unmarshaling", func() {
		It("should equal itself", func() {
			f := func(x uint32) bool {
				buf := new(bytes.Buffer)

				y := abi.NewU32(x)
				err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U32{}
				err = z.Unmarshal(buf)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var _ = Describe("64-bit unsigned integer", func() {
	Context("when marshaling and unmarshaling", func() {
		It("should equal itself", func() {
			f := func(x uint64) bool {
				buf := new(bytes.Buffer)

				y := abi.NewU64(x)
				err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U64{}
				err = z.Unmarshal(buf)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var _ = Describe("128-bit unsigned integer", func() {
	Context("when marshaling and unmarshaling", func() {
		It("should equal itself", func() {
			f := func(x [16]byte) bool {
				buf := new(bytes.Buffer)

				y := abi.NewU128(x)
				err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U128{}
				err = z.Unmarshal(buf)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var _ = Describe("256-bit unsigned integer", func() {
	Context("when marshaling and unmarshaling", func() {
		It("should equal itself", func() {
			f := func(x [32]byte) bool {
				buf := new(bytes.Buffer)

				y := abi.NewU256(x)
				err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U256{}
				err = z.Unmarshal(buf)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
