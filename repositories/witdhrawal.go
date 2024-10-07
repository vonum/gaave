package repositories

import (
	"github.com/vonum/gaave/aave"
	"github.com/vonum/gaave/models"
	"gorm.io/gorm"
)

type WithdrawalRepo struct {
  DB *gorm.DB
}

func (r *WithdrawalRepo) AddWithdrawals(withdrawalEvents []aave.Withdrawal) error {
  var withdrawals []*models.Withdrawal
  for _, event := range withdrawalEvents {
    var withdrawal models.Withdrawal

    withdrawal.BlockNumber = event.BlockNumber
    withdrawal.TxHash = event.TxHash.Hex()
    withdrawal.Reserve = event.Reserve.Hex()
    withdrawal.User = event.User.Hex()
    withdrawal.To = event.To.Hex()
    withdrawal.Amount = event.Amount.String()

    withdrawals = append(withdrawals, &withdrawal)
  }

  result := r.DB.Create(withdrawals)
  return result.Error
}

func (r *WithdrawalRepo) FindWithdrawals() ([]models.Withdrawal, error) {
  var withdrawals []models.Withdrawal

  result := r.DB.Find(&withdrawals)

  if result.Error != nil {
    return nil, result.Error
  }

  return withdrawals, nil

}
