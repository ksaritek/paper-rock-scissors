package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ksaritek/paper-rock-scissors/internal/cmd"
	pb "github.com/ksaritek/paper-rock-scissors/proto/gen/go/game/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	moveMapToProto = map[string]pb.Move{
		"rock":     pb.Move_MOVE_ROCK,
		"scissors": pb.Move_MOVE_SCISSORS,
		"paper":    pb.Move_MOVE_PAPER,
	}

	moveMapFromProto = map[pb.Move]string{
		pb.Move_MOVE_ROCK:     "rock",
		pb.Move_MOVE_SCISSORS: "scissors",
		pb.Move_MOVE_PAPER:    "paper",
	}
)

func main() {
	addr := flag.String("address", "localhost:50051", "address to connect to")
	flag.Parse()

	ctx, stop := cmd.AppContext()
	defer stop()

	fmt.Printf("Connecting to %s\n", *addr)
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGameServiceClient(conn)

	var session *pb.Session

	fmt.Print("Do you want to play the game (y/n): ")
	var command string
	fmt.Scanln(&command)
	if strings.ToLower(command) == "n" {
		fmt.Println("Quitting the game...")
		return
	}

	session = startGame(ctx, client)

outer:
	for {
		fmt.Print("Enter your move (rock / paper / scissors) or quit(q / quit): ")
		var move string
		fmt.Scanln(&move)

		if strings.TrimSpace(move) == "" {
			fmt.Println("Invalid move. Please try again.")
			continue
		}
		if strings.TrimSpace(move) == "quit" || strings.TrimSpace(move) == "q" {
			break outer
		}

		moveCode, ok := moveMapToProto[strings.ToLower(move)]
		if !ok {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		request := &pb.PlayGameRequest{
			SessionId:  session.GetId(),
			PlayerMove: moveCode,
		}
		response, err := client.PlayGame(ctx, request)
		if err != nil {
			log.Fatalf("Failed to play the game: %v", err)
		}

		fmt.Printf("Computer move: %s\n", moveMapFromProto[response.GetComputerMove()])
		printScore(response.GetSession())
	}

	endGame(ctx, client, session)

	fmt.Println("tot zien!....")
}

func endGame(ctx context.Context, client pb.GameServiceClient, session *pb.Session) {
	if session == nil {
		return
	}
	score, err := client.EndSession(ctx, &pb.EndSessionRequest{Session: session})
	if err != nil {
		log.Fatalf("Failed to end the session: %v", err)
	}

	fmt.Printf("summary: draw: %d: %s win: %d computer win: %d\n",
		score.GetSession().GetDraws(),
		score.GetSession().PlayerId,
		score.Session.GetWins(),
		score.Session.GetLosses(),
	)
	printScore(score.GetSession())
}

func startGame(ctx context.Context, client pb.GameServiceClient) *pb.Session {
	fmt.Print("Enter your player ID (or leave blank for a random name): ")
	var playerID string
	fmt.Scanln(&playerID)

	request := &pb.StartSessionRequest{}
	if strings.TrimSpace(playerID) != "" {
		request.PlayerId = &playerID
	}
	response, err := client.StartSession(ctx, request)
	if err != nil {
		log.Fatalf("Failed to start a new session: %v", err)
	}

	fmt.Printf("Session started with ID: %s and Name: %s\n", response.GetSession().GetId(), response.GetSession().GetPlayerId())
	return response.GetSession()
}

func whoIsOnTop(session *pb.Session) string {
	if session.GetWins() > session.GetLosses() {
		return fmt.Sprintf("%s is on top! (+%d)", session.GetPlayerId(), session.GetWins()-session.GetLosses())
	} else if session.GetWins() < session.GetLosses() {
		return fmt.Sprintf("Computer is on top! (+%d)", session.GetLosses()-session.GetWins())
	}
	return "It's a draw!"
}

func printScore(session *pb.Session) {
	totalRound := session.GetWins() + session.GetLosses() + session.GetDraws()
	fmt.Printf("Score: %s Total round: %d\n", whoIsOnTop(session), totalRound)
}
