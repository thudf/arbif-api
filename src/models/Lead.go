package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
)

type Lead struct {
	Id               string    `json:"leadId,omitempty"`
	CreditValue      float64   `json:"creditValue,omitempty"`
	TotalInstallment uint8     `json:"totalInstallment,omitempty"`
	Cnpj             string    `json:"cnpj,omitempty"`
	Name             string    `json:"name,omitempty"`
	Cpf              string    `json:"cpf,omitempty"`
	Email            string    `json:"email,omitempty"`
	Phone            string    `json:"phone,omitempty"`
	OptIn            bool      `json:"optIn,omitempty"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
}

func (lead *Lead) Prepare() error {
	if err := lead.validate(); err != nil {
		return err
	}

	if err := lead.format(); err != nil {
		return err
	}

	return nil
}

func (lead *Lead) validate() error {
	if lead.CreditValue == 0.0 {
		return errors.New("credit value is required")
	}

	if lead.TotalInstallment == 0 {
		return errors.New("total installment is required")
	}

	if lead.Cnpj == "" {
		return errors.New("cnpj is required")
	}

	if lead.Name == "" {
		return errors.New("name is required")
	}

	if lead.Cpf == "" {
		return errors.New("cpf is required")
	}

	if lead.Email == "" {
		return errors.New("email is required")
	}

	if err := checkmail.ValidateFormat(lead.Email); err != nil {
		return errors.New("email invalid")
	}

	if lead.Phone == "" {
		return errors.New("phone is required")
	}

	if !lead.OptIn {
		return errors.New("optin must be true")
	}

	return nil
}

func (lead *Lead) format() error {
	lead.Name = strings.TrimSpace(lead.Name)
	lead.Email = strings.TrimSpace(lead.Email)
	lead.Id = uuid.New().String()

	if lead.Id == "" {
		return errors.New("error creating lead id")
	}

	return nil
}
