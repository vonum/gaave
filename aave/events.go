package aave

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Deposit struct {
  BlockNumber uint64
  TxHash common.Hash
  Index uint
  Reserve common.Address
  User common.Address
  OnBehalfOf common.Address
  Amount *big.Int
  Referal uint16
}

type Withdrawal struct {
  BlockNumber uint64
  TxHash common.Hash
  Index uint
  Reserve common.Address
  User common.Address
  To common.Address
  Amount *big.Int
}

type Borrow struct {
  BlockNumber uint64
  TxHash common.Hash
  Index uint
  Reserve common.Address
  User common.Address
  OnBehalfOf common.Address
  Amount *big.Int
  BorrowRateMode *big.Int
  BorrowRate *big.Int
  Referal uint16
}

type Repay struct {
  BlockNumber uint64
  TxHash common.Hash
  Index uint
  Reserve common.Address
  User common.Address
  Repayer common.Address
  Amount *big.Int
}

type FlashLoan struct {
  BlockNumber uint64
  TxHash common.Hash
  Index uint
  Target common.Address
  Initiator common.Address
  Asset common.Address
  Amount *big.Int
  Premium *big.Int
  ReferalCode uint16
}
