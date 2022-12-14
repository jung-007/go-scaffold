package usecase

import (
	"context"
	"fmt"

	"go-scaffold/internal/entity"
)

type TranslationUseCase struct {
	repo   TranslationRepo
	webAPI TranslationWebAPI
}

func New(r TranslationRepo, w TranslationWebAPI) *TranslationUseCase {
	return &TranslationUseCase{
		repo:   r,
		webAPI: w,
	}
}

// History - getting translate history from store
func (uc *TranslationUseCase) History(ctx context.Context) ([]entity.Translation, error) {
	translations, err := uc.repo.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}
	return translations, nil
}

// Translate
func (uc *TranslationUseCase) Translate(ctx context.Context, t entity.Translation) (entity.Translation, error) {
	Translation, err := uc.webAPI.Translate(t)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	}
	return Translation, nil
}
