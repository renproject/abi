package ext

import (
	"github.com/renproject/abi"
)

// Extension types. The standard unmarshaler will not check these extension
// types. Instead, you must use the unmarshalers in this packages (and
// sub-packages).
const (
	// Cryptographic types.
	TypeCrypto                  = abi.Type(1000)
	TypeCryptoS256N             = TypeCrypto + abi.Type(1)
	TypeCryptoS256P             = TypeCrypto + abi.Type(2)
	TypeCryptoS256PubKey        = TypeCrypto + abi.Type(3)
	TypeCryptoS256PrivKey       = TypeCrypto + abi.Type(4)
	TypeCryptoShamirS256N       = TypeCrypto + abi.Type(5)
	TypeCryptoShamirS256P       = TypeCrypto + abi.Type(6)
	TypeCryptoShamirS256PrivKey = TypeCrypto + abi.Type(7)

	// RenVM types.
	Type          = abi.Type(1100)
	TypeTx        = Type + abi.Type(1)
	TypeArgument  = Type + abi.Type(2)
	TypeArguments = Type + abi.Type(3)

	// Ethereum types.
	TypeEthereum        = abi.Type(1200)
	TypeEthereumAddress = TypeEthereum + abi.Type(1)
	TypeEthereumTx      = TypeEthereum + abi.Type(2)

	// Bitcoin types.
	TypeBitcoin          = abi.Type(1300)
	TypeBitcoinAddress   = TypeBitcoin + abi.Type(1)
	TypeBitcoinUTXOIndex = TypeBitcoin + abi.Type(2)
	TypeBitcoinUTXO      = TypeBitcoin + abi.Type(3)

	// ZCash types.
	TypeZCash        = abi.Type(1400)
	TypeZCashAddress = TypeZCash + abi.Type(1)

	// Bitcoin Cash types.
	TypeBitcoinCash        = abi.Type(1500)
	TypeBitcoinCashAddress = TypeBitcoinCash + abi.Type(1)

	// Litecoin types.
	TypeLitecoin        = abi.Type(1700)
	TypeLitecoinAddress = TypeLitecoin + abi.Type(1)
)
