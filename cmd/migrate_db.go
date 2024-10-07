package main

import (
	"fmt"
	"log"

	"github.com/vonum/gaave/db"
	"github.com/vonum/gaave/repositories"
)

func main() {
  url := "host=localhost user=api password=api dbname=gaave port=5432"
  db := db.Init(url)
  fmt.Println(db)

  depositRepo := repositories.DepositRepo{DB: db}
  withdrawalRepo := repositories.WithdrawalRepo{DB: db}
  borrowRepo := repositories.BorrowRepo{DB: db}
  repayRepo := repositories.RepayRepo{DB: db}
  flashLoanRepo := repositories.FlashLoanRepo{DB: db}

  deposits, err := depositRepo.FindDeposits()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Deposits:", len(deposits), deposits[0])

  withdrawals, err := withdrawalRepo.FindWitdhrawals()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Withdarals:", len(withdrawals), withdrawals[0])

  borrows, err := borrowRepo.FindBorrows()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Borrows:", len(borrows), borrows[0])

  repays, err := repayRepo.FindRepays()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Repays:", len(repays), repays[0])

  flashLoans, err := flashLoanRepo.FindFlashLoans()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("flashLoans:", len(flashLoans), flashLoans[0])
}
