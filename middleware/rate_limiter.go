package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RateLimiter() echo.MiddlewareFunc {
	return middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20))
}
