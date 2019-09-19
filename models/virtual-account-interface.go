package models

type IVirtualAccountRepository interface {
	GetVAByID(id string) (VirtualAccount, error)
	SaveVA(va VirtualAccount) error
}

type IVirtualAccountService interface {
	SaveVA(va VirtualAccount) error
}