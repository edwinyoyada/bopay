package repository

import (
	"github.com/edwinyoyada/bopay/models"

	"log"
	"database/sql"
)

type VirtualAccountRepository struct {
	db *sql.DB
}

func NewVARepo(db *sql.DB) models.IVirtualAccountRepository {
	return &VirtualAccountRepository{
		db: db,
	}
}

func (vr *VirtualAccountRepository) GetVAByID(id string) (models.VirtualAccount, error) {
	va := models.VirtualAccount{}

	sql := `SELECT id, bank_code, is_closed, expected_amount, external_id, account_number, "name", expiration_date, status 
	FROM virtual_accounts WHERE id = $1`

	row := vr.db.QueryRow(sql, id)
	err := row.Scan(&va.ID,
		&va.BankCode,
		&va.IsClosed,
		&va.ExpectedAmount,
		&va.ExternalID,
		&va.AccountNumber,
		&va.Name,
		&va.ExpirationDate,
		&va.Status)

	if err != nil {
		log.Print(err)
		return models.VirtualAccount{}, err
	}

	return va, nil

}

func (vr *VirtualAccountRepository) SaveVA(va models.VirtualAccount) error {
	sql := `INSERT INTO virtual_accounts
	(id, bank_code, is_closed, expected_amount, external_id, account_number, "name", expiration_date, status)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	_, err := vr.db.Exec(sql,
		va.ID,
		va.BankCode,
		va.IsClosed,
		va.ExpectedAmount,
		va.ExternalID,
		va.AccountNumber,
		va.Name,
		va.ExpirationDate,
		va.Status)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}