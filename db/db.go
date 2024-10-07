package db

import (
	"log"

	"github.com/vonum/gaave/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
  db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

  if err != nil {
    log.Fatal("Could not connect to postgres")
  }

  db.AutoMigrate(
    &models.Deposit{},
    &models.Withdrawal{},
    &models.Borrow{},
    &models.Repay{},
    &models.FlashLoan{},
  )

  return db
}
