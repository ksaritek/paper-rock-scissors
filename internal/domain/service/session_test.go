package service

import (
	"context"
	"testing"
	"time"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/repository/memory"
	"github.com/ksaritek/paper-rock-scissors/internal/logger"
	"github.com/stretchr/testify/assert"
)

func Test_Start_Session(t *testing.T) {
	ctx := context.Background()

	repo := memory.NewMemoryStore(ctx, logger.NewLogger("debug"), time.Hour)

	service := NewSessionService(context.Background(), repo)

	playerId := "player1"
	expectedSession := &model.Session{
		PlayerId: playerId,
		Wins:     0,
		Losses:   0,
		Draws:    0,
	}

	session, err := service.StartSession(ctx, playerId)
	expectedSession.Id = session.Id

	assert.NoError(t, err)

	assert.Equal(t, expectedSession, session)
}

func Test_Get_Session(t *testing.T) {
	ctx := context.Background()

	repo := memory.NewMemoryStore(ctx, logger.NewLogger("debug"), time.Hour)

	service := NewSessionService(context.Background(), repo)

	playerId := "player1"
	session, err := service.StartSession(ctx, playerId)

	assert.NoError(t, err)

	s, err := service.GetSession(ctx, session.Id)

	assert.NoError(t, err)

	assert.Equal(t, session, s)
}

func Test_Delete_Session(t *testing.T) {
	ctx := context.Background()

	repo := memory.NewMemoryStore(ctx, logger.NewLogger("debug"), time.Hour)

	service := NewSessionService(context.Background(), repo)

	playerId := "player1"
	session, err := service.StartSession(ctx, playerId)

	assert.NoError(t, err)

	s, err := service.DeleteSession(ctx, session.Id)

	assert.NoError(t, err)

	assert.Equal(t, session, s)
}

func Test_Move(t *testing.T) {
	ctx := context.Background()

	repo := memory.NewMemoryStore(ctx, logger.NewLogger("debug"), time.Hour)

	service := NewSessionService(context.Background(), repo)

	playerId := "player1"
	session, err := service.StartSession(ctx, playerId)

	assert.NoError(t, err)

	move := model.Move{
		Id:     session.Id,
		Choice: model.PAPER,
	}

	s, c, err := service.Move(ctx, move, func() model.Choice { return model.PAPER })
	expectedSession := &model.Session{
		Id:       session.Id,
		PlayerId: playerId,
		Wins:     0,
		Losses:   0,
		Draws:    1,
	}
	assert.NoError(t, err)

	assert.Equal(t, s, expectedSession)
	assert.Equal(t, model.PAPER, c)
}
