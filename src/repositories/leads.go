package repositories

import (
	"arbif/src/models"
	"database/sql"
)

type leads struct {
	db *sql.DB
}

func NewLeadsRepository(db *sql.DB) *leads {
	return &leads{db}
}

func (repository leads) Create(lead models.Lead) error {
	statement, err := repository.db.Prepare(`
		insert into leads (
			id, 
			credit_value, 
			total_installment, 
			cnpj, 
			name, 
			cpf, 
			email, 
			phone, 
			opt_in
		) 
		values(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		lead.Id,
		lead.CreditValue,
		lead.TotalInstallment,
		lead.Cnpj,
		lead.Name,
		lead.Cpf,
		lead.Email,
		lead.Phone,
		lead.OptIn,
	)
	if err != nil {
		return err
	}

	return nil
}
