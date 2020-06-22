package humansrv

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matiasresca/xmen-meli/mocks"
	"github.com/matiasresca/xmen-meli/pkg/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestService_CheckMutant_NewHuman(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepo := mocks.NewMockHumanRepository(ctrl)

	service := NewService(mockRepo)

	dna := []string{
		"ATGCGA",
		"CAGTCC",
		"TTCTAT",
		"AGAGGT",
		"CCGCTA",
		"TGCTTC",
	}

	mockRepo.
		EXPECT().
		GetByDna(dna).
		Return(nil, errors.New("mongo: no documents in result"))

	mockRepo.
		EXPECT().
		Save(domain.Human{dna, false}).
		Return(nil)

	isMutant, err := service.CheckMutant(dna)

	assert.Equal(t, false, isMutant)
	assert.Empty(t, err)
}

func TestService_CheckMutant_NewMutant(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	mockRepo := mocks.NewMockHumanRepository(ctrl)

	service := NewService(mockRepo)

	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TCATGT",
		"AGAAGG",
		"CCCATA",
		"TCACTG",
	}

	mockRepo.
		EXPECT().
		GetByDna(dna).
		Return(nil, errors.New("mongo: no documents in result"))

	mockRepo.
		EXPECT().
		Save(domain.Human{dna, true}).
		Return(nil)

	//Aca se prueba.-
	isMutant, err := service.CheckMutant(dna)

	assert.Equal(t, true, isMutant)
	assert.Empty(t, err)
}

func TestService_CheckMutant_GetHuman(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	mockRepo := mocks.NewMockHumanRepository(ctrl)

	service := NewService(mockRepo)

	dna := []string{
		"ATGCGA",
		"CAGTCC",
		"TTCTAT",
		"AGAGGT",
		"CCGCTA",
		"TGCTTC",
	}

	mockRepo.
		EXPECT().
		GetByDna(dna).
		Return(&domain.Human{dna, false}, nil)

	isMutant, err := service.CheckMutant(dna)

	assert.Equal(t, false, isMutant)
	assert.Empty(t, err)
}

func TestService_CheckMutant_GetMutant(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepo := mocks.NewMockHumanRepository(ctrl)

	service := NewService(mockRepo)

	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TCATGT",
		"AGAAGG",
		"CCCATA",
		"TCACTG",
	}

	mockRepo.
		EXPECT().
		GetByDna(dna).
		Return(&domain.Human{dna, true}, nil)

	//Aca se prueba.-
	isMutant, err := service.CheckMutant(dna)

	assert.Equal(t, true, isMutant)
	assert.Empty(t, err)
}

func TestService_CheckMutant_ErrorNxN(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepo := mocks.NewMockHumanRepository(ctrl)

	service := NewService(mockRepo)

	dna := []string{
		"ATGCG",
		"CAGTGC",
		"TCATG",
		"AGAAGG",
		"CCCATA",
		"TCACTG",
	}

	//Aca se prueba.-
	isMutant, err := service.CheckMutant(dna)

	assert.Equal(t, false, isMutant)
	assert.Error(t, err)
}

func TestService_CheckMutant_ErrorLetter(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepo := mocks.NewMockHumanRepository(ctrl)

	service := NewService(mockRepo)

	dna := []string{
		"ATGCGW",
		"CAGTGH",
		"TCATGA",
		"AGAAGG",
		"CCCATA",
		"TCACTG",
	}

	//Aca se prueba.-
	isMutant, err := service.CheckMutant(dna)

	assert.Equal(t, false, isMutant)
	assert.Error(t, err)
}

func TestService_GetStatsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepo := mocks.NewMockHumanRepository(ctrl)

	service := NewService(mockRepo)

	humans := []domain.Human{}
	mockRepo.
		EXPECT().
		GetAll().
		//Return(humans, nil)
		Return(humans, nil)

	//Aca se prueba.-
	stats, err := service.GetStats()

	assert.Equal(t, &domain.Stats{0, 0, 0}, stats)
	assert.Empty(t, err)
}

func TestService_GetStatsValue(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepo := mocks.NewMockHumanRepository(ctrl)

	service := NewService(mockRepo)

	humans := []domain.Human{
		{nil, true},
		{nil, true},
		{nil, true},
		{nil, false},
		{nil, false},
		{nil, false},
		{nil, false},
		{nil, false},
		{nil, false},
		{nil, false},
		{nil, false},
		{nil, false},
		{nil, false},
	}
	mockRepo.
		EXPECT().
		GetAll().
		//Return(humans, nil)
		Return(humans, nil)

	//Aca se prueba.-
	stats, err := service.GetStats()

	assert.Equal(t, &domain.Stats{3, 10, 0.3}, stats)
	assert.Empty(t, err)
}

func TestService_GetStatsError(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepo := mocks.NewMockHumanRepository(ctrl)

	service := NewService(mockRepo)

	humans := []domain.Human{}

	mockRepo.
		EXPECT().
		GetAll().
		Return(humans, errors.New("ERROR"))

	//Aca se prueba.-
	stats, err := service.GetStats()

	assert.Equal(t, &domain.Stats{}, stats)
	assert.Error(t, err)
}
