package aave

import (
	"context"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const DepositEventName = "Deposit"
const WithdrawalEventName = "Withdraw"
const BorrowEventName = "Borrow"
const RepayEventName = "Repay"
const FlashLoanEventName = "FlashLoan"

const ContractAddress = "0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"

type AaveClient struct {
  client *ethclient.Client
  abi abi.ABI
}

func NewAaveClient(abiPath string, client *ethclient.Client) *AaveClient {
  abiJson, err := os.ReadFile("./contracts/aavev2_pool.json")
  if err != nil {
    log.Fatal("Can't read ABI file")
  }

  abi, err := abi.JSON(strings.NewReader(string(abiJson)))
  if err != nil {
    log.Fatal("Can't parse ABI")
  }


  return &AaveClient{client, abi}
}

func (ac *AaveClient) GetDeposits(from int64, to int64) []Deposit {
  logs := ac.getEvents(from, to, DepositEventName)
  parser := DepositParser{&ac.abi}
  deposits := parser.parseData(logs)

  return deposits
}

func (ac *AaveClient) GetWithdrawals(from int64, to int64) []Withdrawal {
  logs := ac.getEvents(from, to, WithdrawalEventName)
  parser := WithdrawalParser{&ac.abi}
  withdrawals := parser.parseData(logs)

  return withdrawals
}

func (ac *AaveClient) GetBorrows(from int64, to int64) []Borrow {
  logs := ac.getEvents(from, to, BorrowEventName)
  parser := BorrowParser{&ac.abi}
  borrows := parser.parseData(logs)

  return borrows
}

func (ac *AaveClient) GetRepays(from int64, to int64) []Repay {
  logs := ac.getEvents(from, to, RepayEventName)
  parser := RepayParser{&ac.abi}
  repays := parser.parseData(logs)

  return repays
}

func (ac *AaveClient) GetFlashLoans(from int64, to int64) []FlashLoan {
  logs := ac.getEvents(from, to, FlashLoanEventName)
  parser := FlashLoanParser{&ac.abi}
  flashLoans := parser.parseData(logs)

  return flashLoans
}

func (ac *AaveClient) getEvents(from int64, to int64, eventName string) []types.Log {
  depositSig := ac.abi.Events[eventName].Sig
  depositSelector := ac.Selector(depositSig)

  filter := ac.createFilter(from, to, depositSelector)
  logs, err := ac.client.FilterLogs(context.Background(), filter)
  if err != nil {
    log.Fatal("Failed to pull logs")
  }

  return logs
}

func (ac *AaveClient) Selector(signature string) []byte {
  return crypto.Keccak256([]byte(signature))
}

func (ac *AaveClient) createFilter(from int64, to int64, selector []byte)  ethereum.FilterQuery {
  contractAddress := common.HexToAddress(ContractAddress)

  topic := common.Hash(selector)

  return ethereum.FilterQuery{
    FromBlock: big.NewInt(from),
    ToBlock: big.NewInt(to),
    Addresses: []common.Address{contractAddress},
    Topics: [][]common.Hash{
      {topic},
    },
  }
}

