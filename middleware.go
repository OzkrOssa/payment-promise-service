package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func checkTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Obtener el token del encabezado
		token := c.Request().Header.Get("Authorization")

		// Verificar el token (esto es solo un ejemplo, debes implementar la verificación real)
		if token != os.Getenv("TOKEN") {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token inválido"})
		}
		// Continuar al siguiente manejador si el token es válido
		return next(c)
	}
}
