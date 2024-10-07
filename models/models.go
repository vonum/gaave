package models

import "gorm.io/gorm"

type Deposit struct {
  gorm.Model
  Reserve string    `json:"reserve"`
  User string       `json:"user"`
  OnBehalfOf string `json:"onBehalfOf"`
  Amount string     `json:"amount"`
  Referal uint16    `json:"referal"`
}

type Withdrawal struct {
  gorm.Model
  Reserve string    `json:"reserve"`
  User string       `json:"user"`
  To string         `json:"to"`
  Amount string     `json:"amount"`
}

type Borrow struct {
  gorm.Model
  Reserve string        `json:"reserve"`
  User string           `json:"user"`
  OnBehalfOf string     `json:"onBehalfOf"`
  Amount string         `json:"amount"`
  BorrowRateMode string `json:"borrowRateMode"`
  BorrowRate string     `json:"borrowRate"`
  Referal uint16        `json:"referal"`
}

type Repay struct {
  gorm.Model
  Reserve string    `json:"reserve"`
  User string       `json:"user"`
  Repayer string    `json:"repayer"`
  Amount string     `json:"amount"`
}

type FlashLoan struct {
  gorm.Model
  Target string     `json:"target"`
  Initiator string  `json:"initiator"`
  Asset string      `json:"asset"`
  Amount string     `json:"amount"`
  Premium string    `json:"premium"`
  Referal uint16    `json:"referal"`
}
