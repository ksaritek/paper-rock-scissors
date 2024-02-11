package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/repository"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
	"go.uber.org/zap"
)

type Session interface {
	StartSession(ctx context.Context, playerId string) (*model.Session, error)
	GetSession(ctx context.Context, sessionId string) (*model.Session, error)
	DeleteSession(ctx context.Context, sessionId string) (*model.Session, error)
	Move(ctx context.Context, move model.Move) (*model.Session, model.Choice, error)
}

type sessionService struct {
	sessionRepo       repository.Session
	logger            *zap.Logger
	generateSessionId func() string
	computerMove      func() model.Choice
}

func NewSessionService(ctx context.Context,
	l *zap.Logger,
	generateSessionId func() string,
	computerMove func() model.Choice,
	sessionRepo repository.Session) Session {
	_, span := observability.Tracer().Start(ctx, "session_service_new")
	defer span.End()

	return &sessionService{
		sessionRepo:       sessionRepo,
		logger:            l.With(zap.String("service", "game")),
		generateSessionId: generateSessionId,
		computerMove:      computerMove,
	}
}

func GenerateSessionID() string {
	timestamp := time.Now().UnixNano()
	hash := sha256.Sum256([]byte(strconv.FormatInt(timestamp, 10)))
	return hex.EncodeToString(hash[:])
}

func RandomMove() model.Choice {
	moves := []model.Choice{model.ROCK, model.PAPER, model.SCISSORS}
	randomIndex := rand.Intn(len(moves))
	return moves[randomIndex]
}
