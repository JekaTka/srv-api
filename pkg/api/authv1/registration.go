package authv1

import (
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"net/http"
	"github.com/JekaTka/cryptohex-api/pkg/domain/entity/user"
)

type RegistrationHandler echo.HandlerFunc

type dependancies struct {
	dig.In

	uService user.Service
}

// @Summary Registration for new user
// @Description Should return status 200 with user/JWT
// @Tags api, search
// @ID v1-auth-register
// @Produce json
// @Param q query string true "Query"
// @Param limit query int false "The maximum number of entries to return. If the value exceeds the maximum, then the maximum value will be used. [default: 100]"
// @Param offset query int false "The (zero-based) offset of the first item in the collection to return. [default: 0]"
// @Success 200 {object} registrationResponse
// @Router /v1/auth/register [post]
func makeRegistrationHandler(deps dependancies) RegistrationHandler {
	return func(c echo.Context) error {
		params := registerQueryParams{}

		if err := c.Bind(&params); err != nil {
			return err
		}
		if err := c.Validate(params); err != nil {
			return err
		}

		id, err = deps.uService.Create(params)
		return c.JSON(http.StatusOK, registrationResponse{Status: "Success"})
	}
}
