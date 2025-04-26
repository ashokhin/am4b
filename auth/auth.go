package auth

import (
	"log/slog"
)

type AuthAgent struct {
	logger *slog.Logger
}

func (a *AuthAgent) Auth() {
	a.logger.Info("Hello from auth")
}

func New(l *slog.Logger) AuthAgent {
	return AuthAgent{
		logger: l,
	}
}
