package service

import zodiacRepository "server/internal/services/zodiac/repository"

var _ ZodiacRepository = new(zodiacRepository.ZodiacRepository)

type ZodiacRepository interface {
}

type ZodiacService struct {
	zodiacRepository ZodiacRepository
}

func NewZodiacService(
	zodiacRepository ZodiacRepository,
) *ZodiacService {
	return &ZodiacService{
		zodiacRepository: zodiacRepository,
	}
}
