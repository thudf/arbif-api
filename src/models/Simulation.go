package models

import (
	"errors"
	"math"
	"time"
)

type Simulation struct {
	LeadId           string  `json:"leadId,omitempty"`
	CreditValue      float64 `json:"creditValue,omitempty"`
	TotalInstallment uint8   `json:"totalInstallment,omitempty"`
	FirstInstallment float64 `json:"firstInstallment,omitempty"`
	LastInstallment  float64 `json:"lastInstallment,omitempty"`
}

func (simulation *Simulation) Prepare() error {
	if err := simulation.validate(); err != nil {
		return err
	}

	if err := simulation.format(); err != nil {
		return err
	}

	return nil
}

func (simulation *Simulation) validate() error {
	if simulation.LeadId == "" {
		return errors.New("leadId value is required")
	}

	if simulation.CreditValue == 0.0 {
		return errors.New("credit value is required")
	}

	if simulation.TotalInstallment == 0 {
		return errors.New("total installment is required")
	}

	return nil
}

func (simulation *Simulation) format() error {
	var fee float64 = 3
	initialAmount := simulation.CreditValue
	totalInstallment := float64(simulation.TotalInstallment)

	_, _, day := time.Now().Date()
	if day < 15 {
		fee = 5
	}

	if simulation.TotalInstallment > 36 {
		fee = fee + 10
	}

	feeInDecimal := (fee / 100) + 1
	installmentsInYears := totalInstallment / 12
	finalAmount := initialAmount * math.Pow(feeInDecimal, installmentsInYears)
	feeValue := finalAmount - initialAmount
	rest := math.Mod(initialAmount, totalInstallment)
	regularInstallment := (initialAmount - rest) / totalInstallment

	simulation.FirstInstallment = regularInstallment + feeValue
	simulation.LastInstallment = regularInstallment + rest

	return nil
}
