package aave

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Deposit struct {
  Reserve common.Address
  User common.Address
  OnBehalfOf common.Address
  Amount *big.Int
  Referal uint16
}

type Withdrawal struct {
  Reserve common.Address
  User common.Address
  To common.Address
  Amount *big.Int
}

type Borrow struct {
  Reserve common.Address
  User common.Address
  OnBehalfOf common.Address
  Amount *big.Int
  BorrowRateMode *big.Int
  BorrowRate *big.Int
  Referal uint16
}

type Repay struct {
  Reserve common.Address
  User common.Address
  Repayer common.Address
  Amount *big.Int
}

type FlashLoan struct {
  Target common.Address
  Initiatior common.Address
  Asset common.Address
  Amount *big.Int
  Premium *big.Int
  ReferalCode uint16
}
