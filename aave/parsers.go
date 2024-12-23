package aave

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)



type AbiParser interface {
  parseData(eventName string, data []byte) interface{}
}

type DepositParser struct {
  abi *abi.ABI
}
func (dp *DepositParser) parseData(logs []types.Log) []Deposit {
  var deposits []Deposit
  for _, l := range logs {
    var deposit Deposit
    dp.abi.UnpackIntoInterface(&deposit, DepositEventName, l.Data)
    deposit.BlockNumber = l.BlockNumber
    deposit.TxHash = l.TxHash
    deposit.Index = l.Index
    deposits = append(deposits, deposit)
  }

  return deposits
}

type WithdrawalParser struct {
  abi *abi.ABI
}
func (wp *WithdrawalParser) parseData(logs []types.Log) []Withdrawal {
  var withdrawals []Withdrawal
  for _, l := range logs {
    var withdrawal Withdrawal
    wp.abi.UnpackIntoInterface(&withdrawal, WithdrawalEventName, l.Data)
    withdrawal.BlockNumber = l.BlockNumber
    withdrawal.TxHash = l.TxHash
    withdrawal.Index = l.Index
    withdrawals = append(withdrawals, withdrawal)
  }

  return withdrawals
}

type BorrowParser struct {
  abi *abi.ABI
}
func (bp *BorrowParser) parseData(logs []types.Log) []Borrow {
  var borrows []Borrow
  for _, l := range logs {
    var borrow Borrow
    bp.abi.UnpackIntoInterface(&borrow, BorrowEventName, l.Data)
    borrow.BlockNumber = l.BlockNumber
    borrow.TxHash = l.TxHash
    borrow.Index = l.Index
    borrows = append(borrows, borrow)
  }

  return borrows
}


type RepayParser struct {
  abi *abi.ABI
}
func (bp *RepayParser) parseData(logs []types.Log) []Repay {
  var repays []Repay
  for _, l := range logs {
    var repay Repay
    bp.abi.UnpackIntoInterface(&repay, RepayEventName, l.Data)
    repay.BlockNumber = l.BlockNumber
    repay.TxHash = l.TxHash
    repay.Index = l.Index
    repays = append(repays, repay)
  }

  return repays
}


type FlashLoanParser struct {
  abi *abi.ABI
}
func (flp *FlashLoanParser) parseData(logs []types.Log) []FlashLoan {
  var flashLoans []FlashLoan
  for _, l := range logs {
    var flashLoan FlashLoan
    flp.abi.UnpackIntoInterface(&flashLoan, FlashLoanEventName, l.Data)
    flashLoan.BlockNumber = l.BlockNumber
    flashLoan.TxHash = l.TxHash
    flashLoan.Index = l.Index
    flashLoans = append(flashLoans, flashLoan)
  }

  return flashLoans
}
