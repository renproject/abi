package ext

import "github.com/renproject/abi"

// Enumeration of extension types.
const (
	TypeUnrecognised = abi.Type(1000)

	// Cryptographic types.
	TypeCrypto = abi.Type(1100)

	// RenVM types.
	TypeRenVM          = abi.Type(1200)
	TypeRenVMTx        = TypeRenVM + abi.Type(1)
	TypeRenVMArgument  = TypeRenVM + abi.Type(2)
	TypeRenVMArguments = TypeRenVM + abi.Type(3)

	// Ethereum types.
	TypeEthereum        = abi.Type(1300)
	TypeEthereumAddress = TypeEthereum + abi.Type(1)
	TypeEthereumTx      = TypeEthereum + abi.Type(2)

	// Bitcoin types.
	TypeBitcoin          = abi.Type(1400)
	TypeBitcoinAddress   = TypeBitcoin + abi.Type(1)
	TypeBitcoinUTXOIndex = TypeBitcoin + abi.Type(2)
	TypeBitcoinUTXO      = TypeBitcoin + abi.Type(3)

	// ZCash types.
	TypeZCash        = abi.Type(1500)
	TypeZCashAddress = TypeZCash + abi.Type(1)
	// TypeZCashUTXOIndex = TypeZCash + abi.Type(2)
	// TypeZCashUTXO      = TypeZCash + abi.Type(3)

	// Bitcoin Cash types.
	TypeBCash        = abi.Type(1600)
	TypeBCashAddress = TypeBCash + abi.Type(1)
	// TypeBCashUTXOIndex = TypeBCash + abi.Type(2)
	// TypeBCashUTXO      = TypeBCash + abi.Type(3)

	// Litecoin types.
	TypeLitecoin        = abi.Type(1700)
	TypeLitecoinAddress = TypeLitecoin + abi.Type(1)
	// TypeLitecoinUTXOIndex = TypeLitecoin + abi.Type(2)
	// TypeLitecoinUTXO      = TypeLitecoin + abi.Type(3)
)
