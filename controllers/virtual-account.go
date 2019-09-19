package controllers

import (
	"github.com/edwinyoyada/bopay/models"
	"github.com/labstack/echo/v4"

	"net/http"
	"log"
)

type VirtualAccountController struct {
	VirtualAccountService models.IVirtualAccountService
}

func NewVAController(vas models.IVirtualAccountService) VirtualAccountController {
	return VirtualAccountController{
		VirtualAccountService: vas,
	}
}

func (vac *VirtualAccountController) UpdateVACallback(c echo.Context) error {
	va := new(models.VirtualAccount)

	if err := c.Bind(va); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err := vac.VirtualAccountService.SaveVA(*va)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, va)
}
