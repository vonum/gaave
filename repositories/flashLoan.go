package repositories

import (
	"github.com/vonum/gaave/aave"
	"github.com/vonum/gaave/models"
	"gorm.io/gorm"
)

type FlashLoanRepo struct {
  DB *gorm.DB
}

func (r *FlashLoanRepo) AddFlashLoans(flashLoanEvents []aave.FlashLoan) error {
  var flashLoans []*models.FlashLoan
  for _, event := range flashLoanEvents {
    var flashLoan models.FlashLoan

    flashLoan.BlockNumber = event.BlockNumber
    flashLoan.TxHash = event.TxHash.Hex()
    flashLoan.LogIndex = event.Index
    flashLoan.Target = event.Target.Hex()
    flashLoan.Initiator = event.Initiator.Hex()
    flashLoan.Asset = event.Asset.Hex()
    flashLoan.Amount = event.Amount.String()
    flashLoan.Premium = event.Premium.String()
    flashLoan.Referal = event.ReferalCode

    flashLoans = append(flashLoans, &flashLoan)
  }

  result := r.DB.Create(flashLoans)
  return result.Error
}

func (r *FlashLoanRepo) FindFlashLoans() ([]models.FlashLoan, error) {
  var flashLoans []models.FlashLoan

  result := r.DB.Find(&flashLoans)

  if result.Error != nil {
    return nil, result.Error
  }

  return flashLoans, nil

}
