package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vonum/gaave/db"
)

type Server struct {}

func (s *Server) Run(port int, dbUrl string) error {
  db := db.Init(dbUrl)
  h := handler{DB: db}

  r := gin.Default()

  r.GET("/", DefaultHandler)
  r.GET("/deposits", h.GetDeposits)
  r.GET("/withdrawals", h.GetWithdrawals)
  r.GET("/borrows", h.GetBorrows)
  r.GET("/repays", h.GetRepays)
  r.GET("/flashLoans", h.GetFlashLoans)

  err := r.Run(fmt.Sprintf(":%d", port))
  return err
}
