package usecase

import (
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
)

type refreshTokenUseCase struct {
	useRepository  domain.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUseCase(userRepository)
