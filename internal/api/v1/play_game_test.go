package v1

import (
	"context"
	"testing"
	"time"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/repository/memory"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/service"
	"github.com/ksaritek/paper-rock-scissors/internal/logger"
	pb "github.com/ksaritek/paper-rock-scissors/proto/gen/go/game/v1"
	"github.com/stretchr/testify/assert"
)

func Test_controller_PlayGame(t *testing.T) {
	ctx := context.Background()

	repo := memory.NewMemoryStore(ctx, logger.NewLogger("debug"), time.Hour)
	l := logger.NewLogger("debug")
	service := service.NewSessionService(context.Background(), l, func() string { return "123456" }, func() model.Choice { return model.PAPER }, repo)
	ctrl := &controller{
		logger: l,
		srv:    service,
	}
	playerId := "player1"
	ctrl.StartSession(ctx, &pb.StartSessionRequest{
		PlayerId: &playerId,
	})

	playReq := &pb.PlayGameRequest{
		PlayerMove: pb.Move_MOVE_ROCK,
		SessionId:  "123456",
	}

	playResp, err := ctrl.PlayGame(ctx, playReq)

	assert.NoError(t, err)
	assert.NotNil(t, playResp)
	assert.Equal(t, pb.Move_MOVE_PAPER, playResp.GetComputerMove())
	assert.Equal(t, "123456", playResp.GetSession().GetId())
	assert.Equal(t, int32(0), playResp.GetSession().GetWins())
	assert.Equal(t, int32(1), playResp.GetSession().GetLosses())
	assert.Equal(t, int32(0), playResp.GetSession().GetDraws())
}
