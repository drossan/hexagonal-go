package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthMiddleware struct {
	username          string
	password          string
	maxFailedAttempts int
	lockoutDuration   time.Duration
	failedAttempts    map[string]int
	lockoutEndTime    map[string]time.Time
	mutex             sync.Mutex
}

func NewAuthMiddleware(username, password string, maxFailedAttempts int, lockoutDuration time.Duration) *AuthMiddleware {
	return &AuthMiddleware{
		username:          username,
		password:          password,
		maxFailedAttempts: maxFailedAttempts,
		lockoutDuration:   lockoutDuration,
		failedAttempts:    make(map[string]int),
		lockoutEndTime:    make(map[string]time.Time),
	}
}

func (a *AuthMiddleware) BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
		clientIP := c.RealIP()
		a.mutex.Lock()
		defer a.mutex.Unlock()

		if lockoutEnd, ok := a.lockoutEndTime[clientIP]; ok {
			if time.Now().Before(lockoutEnd) {
				return false, echo.NewHTTPError(http.StatusTooManyRequests, "Too many failed attempts. Try again later.")
			}
			delete(a.lockoutEndTime, clientIP)
			delete(a.failedAttempts, clientIP)
		}

		if u == a.username && p == a.password {
			delete(a.failedAttempts, clientIP)
			return true, nil
		}

		a.failedAttempts[clientIP]++
		if a.failedAttempts[clientIP] >= a.maxFailedAttempts {
			a.lockoutEndTime[clientIP] = time.Now().Add(a.lockoutDuration)
			return false, echo.NewHTTPError(http.StatusTooManyRequests, "Too many failed attempts. Try again later.")
		}

		return false, nil
	})
}
