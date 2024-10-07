package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vonum/gaave/repositories"
	"gorm.io/gorm"
)

type handler struct {
  DB *gorm.DB
}

func DefaultHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, "Hello")
}

func (h handler) GetDeposits(ctx *gin.Context) {
  depositRepo := repositories.DepositRepo{DB: h.DB}

  deposits, err := depositRepo.FindDeposits()

  if err != nil {
      ctx.AbortWithError(http.StatusNotFound, err)
      return
  }

  ctx.JSON(http.StatusOK, &deposits)
}

func (h handler) GetWithdrawals(ctx *gin.Context) {
  withdrawalRepo := repositories.WithdrawalRepo{DB: h.DB}

  withdrawals, err := withdrawalRepo.FindWithdrawals()

  if err != nil {
      ctx.AbortWithError(http.StatusNotFound, err)
      return
  }

  ctx.JSON(http.StatusOK, &withdrawals)
}

func (h handler) GetBorrows(ctx *gin.Context) {
  borrowRepo := repositories.BorrowRepo{DB: h.DB}

  borrows, err := borrowRepo.FindBorrows()

  if err != nil {
      ctx.AbortWithError(http.StatusNotFound, err)
      return
  }

  ctx.JSON(http.StatusOK, &borrows)
}

func (h handler) GetRepays(ctx *gin.Context) {
  repayRepo := repositories.RepayRepo{DB: h.DB}

  repays, err := repayRepo.FindRepays()

  if err != nil {
      ctx.AbortWithError(http.StatusNotFound, err)
      return
  }

  ctx.JSON(http.StatusOK, &repays)
}

func (h handler) GetFlashLoans(ctx *gin.Context) {
  flashLoanRepo := repositories.FlashLoanRepo{DB: h.DB}

  flashLoans, err := flashLoanRepo.FindFlashLoans()

  if err != nil {
      ctx.AbortWithError(http.StatusNotFound, err)
      return
  }

  ctx.JSON(http.StatusOK, &flashLoans)
}
