package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/vonum/gaave/aave"
)

func main() {
  const BlockRange = 10000
  const ContractAddress = "0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"
  ctx := context.Background()

  err := godotenv.Load()
  rpcUrl := os.Getenv("RPC_URL")

  client, err := ethclient.Dial(rpcUrl)

  maxBn, err := client.BlockNumber(ctx)
  if err != nil {
    log.Fatal("Can't fetch block number")
  }
  minBn := maxBn - BlockRange
  fmt.Println(maxBn, minBn)

  aaveClient := aave.NewAaveClient(rpcUrl, client)
  deposits := aaveClient.GetDeposits(int64(minBn), int64(maxBn))
  fmt.Println("\n\n", deposits[0].Amount)
  fmt.Println("\n\n", len(deposits))

  ws := aaveClient.GetWithdrawals(int64(minBn), int64(maxBn))
  fmt.Println("\n\n", ws[0].Amount)
  fmt.Println("\n\n", len(ws))

  fls := aaveClient.GetFlashLoans(int64(minBn), int64(maxBn))
  fmt.Println("\n\n", fls[0].Amount)
  fmt.Println("\n\n", len(fls))
}
