package abi_test

import (
	"bytes"
	"testing/quick"

	"github.com/renproject/abi"
	"github.com/renproject/surge"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("8-bit unsigned integer", func() {
	Context("when marshaling and unmarshaling", func() {
		It("should equal itself", func() {
			f := func(x uint8) bool {
				buf := new(bytes.Buffer)

				y := abi.NewU8(x)
				_, err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U8{}
				_, err = z.Unmarshal(buf, surge.MaxBytes)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when marshaling and unmarshaling to/from JSON", func() {
		It("should equal itself", func() {
			f := func(x uint8) bool {
				y := abi.NewU8(x)
				data, err := y.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())

				z := abi.U8{}
				err = z.UnmarshalJSON(data)
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
				_, err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U16{}
				_, err = z.Unmarshal(buf, surge.MaxBytes)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when marshaling and unmarshaling to/from JSON", func() {
		It("should equal itself", func() {
			f := func(x uint16) bool {
				y := abi.NewU16(x)
				data, err := y.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())

				z := abi.U16{}
				err = z.UnmarshalJSON(data)
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
				_, err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U32{}
				_, err = z.Unmarshal(buf, surge.MaxBytes)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when marshaling and unmarshaling to/from JSON", func() {
		It("should equal itself", func() {
			f := func(x uint32) bool {
				y := abi.NewU32(x)
				data, err := y.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())

				z := abi.U32{}
				err = z.UnmarshalJSON(data)
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
				_, err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U64{}
				_, err = z.Unmarshal(buf, surge.MaxBytes)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when marshaling and unmarshaling to/from JSON", func() {
		It("should equal itself", func() {
			f := func(x uint64) bool {
				y := abi.NewU64(x)
				data, err := y.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())

				z := abi.U64{}
				err = z.UnmarshalJSON(data)
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
				_, err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U128{}
				_, err = z.Unmarshal(buf, surge.MaxBytes)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when marshaling and unmarshaling to/from JSON", func() {
		It("should equal itself", func() {
			f := func(x [16]byte) bool {
				y := abi.NewU128(x)
				data, err := y.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())

				z := abi.U128{}
				err = z.UnmarshalJSON(data)
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
				_, err := y.Marshal(buf)
				Expect(err).ToNot(HaveOccurred())

				z := abi.U256{}
				_, err = z.Unmarshal(buf, surge.MaxBytes)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when marshaling and unmarshaling to/from JSON", func() {
		It("should equal itself", func() {
			f := func(x [32]byte) bool {
				y := abi.NewU256(x)
				data, err := y.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())

				z := abi.U256{}
				err = z.UnmarshalJSON(data)
				Expect(err).ToNot(HaveOccurred())

				Expect(y).To(Equal(z))
				return true
			}

			err := quick.Check(f, nil)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
