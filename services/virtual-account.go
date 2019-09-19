package services

import (
	"github.com/edwinyoyada/bopay/models"
)

type VirtualAccountService struct {
	VirtualAccountRepository models.IVirtualAccountRepository
}

func NewVAService(vr models.IVirtualAccountRepository) models.IVirtualAccountService {
	return &VirtualAccountService{
		VirtualAccountRepository: vr,
	}
}

func (vas *VirtualAccountService) SaveVA(va models.VirtualAccount) error {
	return vas.VirtualAccountRepository.SaveVA(va)
}