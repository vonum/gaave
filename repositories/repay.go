package repositories

import (
	"github.com/vonum/gaave/aave"
	"github.com/vonum/gaave/models"
	"gorm.io/gorm"
)

type RepayRepo struct {
  DB *gorm.DB
}

func (r *RepayRepo) AddRepays(repayEvents []aave.Repay) error {
  var repays []*models.Repay
  for _, event := range repayEvents {
    var repay models.Repay

    repay.Reserve = event.Reserve.Hex()
    repay.User = event.User.Hex()
    repay.Repayer = event.Repayer.Hex()
    repay.Amount = event.Amount.String()

    repays = append(repays, &repay)
  }

  result := r.DB.Create(repays)
  return result.Error
}

func (r *RepayRepo) FindRepays() ([]models.Repay, error) {
  var repays []models.Repay

  result := r.DB.Find(&repays)

  if result.Error != nil {
    return nil, result.Error
  }

  return repays, nil

}
