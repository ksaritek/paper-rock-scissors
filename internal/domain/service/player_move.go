package service

import (
	"context"

	"github.com/friendsofgo/errors"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
)

var (
	ErrInvalidChoice = errors.New("invalid choice")
)

func (s *sessionService) Move(ctx context.Context, move model.Move) (_ *model.Session, _ model.Choice, retErr error) {
	ctx, span := observability.Tracer().Start(ctx, "service_move")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	v := s.computerMove()

	playerResult, err := determinePlayerResult(move.Choice, v)
	if err != nil {
		return nil, v, errors.WithMessage(err, "failed to determine who wins the round")
	}

	session, err := s.sessionRepo.UpdateRoundResult(ctx, move.Id, playerResult)
	if err != nil {
		return nil, model.INVALID_CHOICE, errors.WithMessage(err, "failed to update session")
	}

	return session, v, nil
}

func determinePlayerResult(playerMove model.Choice, computerMove model.Choice) (model.RoundResult, error) {
	switch playerMove {
	case model.ROCK:
		switch computerMove {
		case model.ROCK:
			return model.DRAW, nil
		case model.PAPER:
			return model.LOSS, nil
		case model.SCISSORS:
			return model.WIN, nil
		}
	case model.PAPER:
		switch computerMove {
		case model.ROCK:
			return model.WIN, nil
		case model.PAPER:
			return model.DRAW, nil
		case model.SCISSORS:
			return model.LOSS, nil
		}
	case model.SCISSORS:
		switch computerMove {
		case model.ROCK:
			return model.LOSS, nil
		case model.PAPER:
			return model.WIN, nil
		case model.SCISSORS:
			return model.DRAW, nil
		}
	}

	return model.INVALID_RESULT, ErrInvalidChoice
}
