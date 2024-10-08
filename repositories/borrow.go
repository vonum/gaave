package repositories

import (
	"github.com/vonum/gaave/aave"
	"github.com/vonum/gaave/models"
	"gorm.io/gorm"
)

type BorrowRepo struct {
  DB *gorm.DB
}

func (r *BorrowRepo) AddBorrows(borrowEvents []aave.Borrow) error {
  var borrows []*models.Borrow
  for _, event := range borrowEvents {
    var borrow models.Borrow

    borrow.BlockNumber = event.BlockNumber
    borrow.TxHash = event.TxHash.Hex()
    borrow.LogIndex = event.Index
    borrow.Reserve = event.Reserve.Hex()
    borrow.User = event.User.Hex()
    borrow.OnBehalfOf = event.OnBehalfOf.Hex()
    borrow.Amount = event.Amount.String()
    borrow.BorrowRateMode = event.BorrowRateMode.String()
    borrow.BorrowRate = event.BorrowRate.String()
    borrow.Referal = event.Referal

    borrows = append(borrows, &borrow)
  }

  result := r.DB.Create(borrows)
  return result.Error
}

func (r *BorrowRepo) FindBorrows() ([]models.Borrow, error) {
  var borrows []models.Borrow

  result := r.DB.Find(&borrows)

  if result.Error != nil {
    return nil, result.Error
  }

  return borrows, nil

}
