package repositories

import (
	"github.com/vonum/gaave/aave"
	"github.com/vonum/gaave/models"
	"gorm.io/gorm"
)

type DepositRepo struct {
  DB *gorm.DB
}

func (r *DepositRepo) AddDeposits(depositEvents []aave.Deposit) error {
  var deposits []*models.Deposit
  for _, event := range depositEvents {
    var deposit models.Deposit

    deposit.BlockNumber = event.BlockNumber
    deposit.TxHash = event.TxHash.Hex()
    deposit.Reserve = event.Reserve.Hex()
    deposit.User = event.User.Hex()
    deposit.OnBehalfOf = event.OnBehalfOf.Hex()
    deposit.Amount = event.Amount.String()
    deposit.Referal = event.Referal

    deposits = append(deposits, &deposit)
  }

  result := r.DB.Create(deposits)
  return result.Error
}

func (r *DepositRepo) FindDeposits() ([]models.Deposit, error) {
  var deposits []models.Deposit

  result := r.DB.Find(&deposits)

  if result.Error != nil {
    return nil, result.Error
  }

  return deposits, nil
}
