package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/ksaritek/paper-rock-scissors/internal/api"
	"github.com/ksaritek/paper-rock-scissors/internal/cmd"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/repository/memory"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/service"
	"github.com/ksaritek/paper-rock-scissors/internal/logger"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
	"github.com/ksaritek/paper-rock-scissors/internal/version"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	gitRepo   = "none"
	gitCommit = "none"
	buildDate = "none"
	buildUser = "none"
	buildTag  = "none"
)

func main() {
	initFigure := figure.NewFigure("!! Game Server !!", "", true)
	initFigure.Print()

	port := flag.String("port", "50051", "port to listen")
	flag.Parse()

	ver := version.Set(
		gitRepo,
		gitCommit,
		buildDate,
		buildUser,
		buildTag,
	)

	l := logger.NewLogger("debug")
	defer l.Sync()

	l = l.With(zap.String("type", "grpc-server"))

	ctx, stop := cmd.AppContext()
	defer stop()

	if value, present := os.LookupEnv("OTEL_EXPORTER_ENABLED"); present && value == "true" {
		l.Info("otel exporter is enabled")
		otelShutdown, err := observability.SetupOTelSDK(ctx)
		if err != nil {
			return
		}

		defer func() {
			err = errors.Join(err, otelShutdown(context.Background()))
		}()
	}

	address := fmt.Sprintf(":%s", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		l.Fatal("failed to listen", zap.Error(err))
	}

	l.Info("listening on", zap.String("address", address))

	grpcServer := cmd.Server(ctx, l)

	repo := memory.NewMemoryStore(ctx, l, time.Minute*10)
	srv := service.NewSessionService(ctx, l, service.GenerateSessionID, service.RandomMove, repo)
	api.Routes(ctx, l, grpcServer, srv)

	defer grpcServer.GracefulStop()
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			l.Fatal("failed to serve", zap.Error(err))
		}
	}()

	l.Info(
		"game-server is started",
		getLogParameters(ver)...,
	)

	<-ctx.Done()
	l.Info("shutting down")
}

func getLogParameters(ver *version.Version) []zapcore.Field {
	return []zapcore.Field{
		zap.String("GoVersion", ver.GoVersion),
		zap.String("GoOs", ver.GoOs),
		zap.String("GoArch", ver.GoArch),
		zap.String("GitRepo", ver.GitRepo),
		zap.String("GitCommit", ver.GitCommit),
		zap.String("BuildDate", ver.BuildDate),
		zap.String("BuildUser", ver.BuildUser),
		zap.String("BuildTag", ver.BuildTag),
	}
}
