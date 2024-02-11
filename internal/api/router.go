package api

import (
	"context"

	api_v1 "github.com/ksaritek/paper-rock-scissors/internal/api/v1"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/service"
	pb "github.com/ksaritek/paper-rock-scissors/proto/gen/go/game/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Routes(ctx context.Context, l *zap.Logger, s *grpc.Server, srv service.Session) {
	reflection.Register(s)

	controller := api_v1.NewController(ctx, l, srv)
	pb.RegisterGameServiceServer(s, controller)
}
