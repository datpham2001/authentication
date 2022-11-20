package api

import (
	"authentication/dao"

	"github.com/labstack/echo"
)

type authHandler struct {
	Dao dao.AuthDAO
}

func NewAuthHandler(dao dao.AuthDAO) *authHandler {
	return &authHandler{
		Dao: dao,
	}
}

func (h *authHandler) Register(c echo.Context) error {

}
