package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/vonum/gaave/aave"
	"github.com/vonum/gaave/db"
	"github.com/vonum/gaave/repositories"
)

func main() {
  const BlockRange = 10000
  const ContractAddress = "0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"
  ctx := context.Background()

  err := godotenv.Load()
  rpcUrl := os.Getenv("RPC_URL")
  dbUrl := os.Getenv("DB_URL")

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

  brs := aaveClient.GetBorrows(int64(minBn), int64(maxBn))
  fmt.Println("\n\n", brs[0].Amount)
  fmt.Println("\n\n", len(brs))

  rps := aaveClient.GetRepays(int64(minBn), int64(maxBn))
  fmt.Println("\n\n", rps[0].Amount)
  fmt.Println("\n\n", len(rps))

  fls := aaveClient.GetFlashLoans(int64(minBn), int64(maxBn))
  fmt.Println("\n\n", fls[0].Amount)
  fmt.Println("\n\n", len(fls))

  db := db.Init(dbUrl)

  dbRepo := repositories.DepositRepo{DB: db}
  err = dbRepo.AddDeposits(deposits)
  if err != nil {
    log.Fatal(err)
  }

  wsRepo := repositories.WithdrawalRepo{DB: db}
  err = wsRepo.AddWithdrawals(ws)
  if err != nil {
    log.Fatal(err)
  }

  brsRepo := repositories.BorrowRepo{DB: db}
  err = brsRepo.AddBorrows(brs)
  if err != nil {
    log.Fatal(err)
  }

  rpsRepo := repositories.RepayRepo{DB: db}
  err = rpsRepo.AddRepays(rps)
  if err != nil {
    log.Fatal(err)
  }

  flsRepo := repositories.FlashLoanRepo{DB: db}
  err = flsRepo.AddFlashLoans(fls)
  if err != nil {
    log.Fatal(err)
  }
}
