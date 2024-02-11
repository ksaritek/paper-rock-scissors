package service

import (
	"testing"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func Test_DeterminePlayerResult(t *testing.T) {
	testCases := []struct {
		playerMove     model.Choice
		computerMove   model.Choice
		expectedResult model.RoundResult
		expectedError  error
	}{
		{
			playerMove:     model.ROCK,
			computerMove:   model.ROCK,
			expectedResult: model.DRAW,
			expectedError:  nil,
		},
		{
			playerMove:     model.ROCK,
			computerMove:   model.PAPER,
			expectedResult: model.LOSS,
			expectedError:  nil,
		},
		{
			playerMove:     model.ROCK,
			computerMove:   model.SCISSORS,
			expectedResult: model.WIN,
			expectedError:  nil,
		},
		{
			playerMove:     model.PAPER,
			computerMove:   model.ROCK,
			expectedResult: model.WIN,
			expectedError:  nil,
		},
		{
			playerMove:     model.PAPER,
			computerMove:   model.PAPER,
			expectedResult: model.DRAW,
			expectedError:  nil,
		},
		{
			playerMove:     model.PAPER,
			computerMove:   model.SCISSORS,
			expectedResult: model.LOSS,
			expectedError:  nil,
		},
		{
			playerMove:     model.SCISSORS,
			computerMove:   model.ROCK,
			expectedResult: model.LOSS,
			expectedError:  nil,
		},
		{
			playerMove:     model.SCISSORS,
			computerMove:   model.PAPER,
			expectedResult: model.WIN,
			expectedError:  nil,
		},
		{
			playerMove:     model.SCISSORS,
			computerMove:   model.SCISSORS,
			expectedResult: model.DRAW,
			expectedError:  nil,
		},
		{
			playerMove:     model.INVALID_CHOICE,
			computerMove:   model.ROCK,
			expectedResult: model.INVALID_RESULT,
			expectedError:  ErrInvalidChoice,
		},
	}

	for _, tc := range testCases {
		result, err := determinePlayerResult(tc.playerMove, tc.computerMove)

		assert.Equal(t, tc.expectedResult, result)
		assert.Equal(t, tc.expectedError, err)
	}
}
