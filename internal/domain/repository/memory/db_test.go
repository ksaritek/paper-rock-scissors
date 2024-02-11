package memory

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	repositoryerr "github.com/ksaritek/paper-rock-scissors/internal/domain/repository/errors"
	"go.uber.org/zap"
)

func setupTest(t *testing.T) *memoryStore {
	ctx := context.Background()
	logger := zap.NewNop()
	ttl := time.Minute

	db := NewMemoryStore(ctx, logger, ttl)

	ms := db.(*memoryStore)

	if ms == nil {
		t.Error("NewMemoryStore returned nil")
	}

	if ms.data == nil {
		t.Error("db.data is nil")
	}
	if ms.ttl != ttl {
		t.Errorf("db.ttl is %v, expected %v", ms.ttl, ttl)
	}
	if ms.expireAt == nil {
		t.Error("db.expireAt is nil")
	}

	ms.data["session1"] = model.Session{
		Id:       "session1",
		PlayerId: "player1",
		Wins:     2,
		Losses:   1,
		Draws:    0,
	}
	ms.expireAt["session1"] = time.Now().Add(ttl)

	ms.data["session2"] = model.Session{
		Id:       "session2",
		PlayerId: "player2",
		Wins:     1,
		Losses:   2,
		Draws:    0,
	}
	ms.expireAt["session2"] = time.Now().Add(ttl)

	ms.data["session3"] = model.Session{
		Id:       "session3",
		PlayerId: "player3",
		Wins:     0,
		Losses:   0,
		Draws:    3,
	}
	ms.expireAt["session3"] = time.Now().Add(ttl)

	return ms
}
func TestStoreGet(t *testing.T) {
	db := setupTest(t)
	ctx := context.Background()

	tests := []struct {
		name      string
		sessionId string
		expected  *model.Session
		err       error
	}{
		{
			name:      "session1",
			sessionId: "session1",
			expected: &model.Session{
				Id:       "session1",
				PlayerId: "player1",
				Wins:     2,
				Losses:   1,
				Draws:    0,
			},
			err: nil,
		},
		{
			name:      "session2",
			sessionId: "session2",
			expected: &model.Session{
				Id:       "session2",
				PlayerId: "player2",
				Wins:     1,
				Losses:   2,
				Draws:    0,
			},
			err: nil,
		},
		{
			name:      "session3",
			sessionId: "session3",
			expected: &model.Session{
				Id:       "session3",
				PlayerId: "player3",
				Wins:     0,
				Losses:   0,
				Draws:    3,
			},
			err: nil,
		},
		{
			name:      "session4",
			sessionId: "session4",
			expected:  nil,
			err:       repositoryerr.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session, err := db.Get(ctx, tt.sessionId)
			if err != nil && tt.err == nil {
				t.Errorf("db.Get returned error: %v", err)
			}
			if session != nil {
				if !reflect.DeepEqual(session, tt.expected) {
					t.Errorf("db.Get returned %v, expected %v", session, tt.expected)
				}
			} else {
				if tt.expected != nil {
					t.Errorf("db.Get returned nil, expected %v", tt.expected)
				}
			}
		})
	}
}

func TestStoreCreate(t *testing.T) {
	db := setupTest(t)
	ctx := context.Background()

	tests := []struct {
		name      string
		sessionId string
		expected  model.Session
		err       error
	}{
		{
			name:      "session1",
			sessionId: "session1",
			expected: model.Session{
				Id:       "session1",
				PlayerId: "player1",
				Wins:     2,
				Losses:   1,
				Draws:    0,
			},
			err: nil,
		},
		{
			name:      "session2",
			sessionId: "session2",
			expected: model.Session{
				Id:       "session2",
				PlayerId: "player2",
				Wins:     1,
				Losses:   2,
				Draws:    0,
			},
			err: nil,
		},
		{
			name:      "session3",
			sessionId: "session3",
			expected: model.Session{
				Id:       "session3",
				PlayerId: "player3",
				Wins:     0,
				Losses:   0,
				Draws:    3,
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// test for Create method
			err := db.Create(ctx, tt.expected)
			if err != nil && tt.err == nil {
				t.Errorf("db.Create returned error: %v", err)
			}
			session := db.data[tt.sessionId]

			if !reflect.DeepEqual(session, tt.expected) {
				t.Errorf("db.Get returned %v, expected %v", session, tt.expected)
			}
		})
	}
}

func TestStoreDelete(t *testing.T) {
	db := setupTest(t)
	ctx := context.Background()

	tests := []struct {
		name      string
		sessionId string
		expected  *model.Session
		err       error
	}{
		{
			name:      "session1",
			sessionId: "session1",
			expected: &model.Session{
				Id:       "session1",
				PlayerId: "player1",
				Wins:     2,
				Losses:   1,
				Draws:    0,
			},
			err: nil,
		},
		{
			name:      "session2",
			sessionId: "session2",
			expected: &model.Session{
				Id:       "session2",
				PlayerId: "player2",
				Wins:     1,
				Losses:   2,
				Draws:    0,
			},
			err: nil,
		},
		{
			name:      "session3",
			sessionId: "session3",
			expected: &model.Session{
				Id:       "session3",
				PlayerId: "player3",
				Wins:     0,
				Losses:   0,
				Draws:    3,
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// test for Delete method
			session, err := db.Delete(ctx, tt.sessionId)
			if err != nil && tt.err == nil {
				t.Errorf("db.Delete returned error: %v", err)
			}
			if session != nil {
				if !reflect.DeepEqual(session, tt.expected) {
					t.Errorf("db.Delete returned %v, expected %v", session, tt.expected)
				}
			} else {
				if tt.expected != nil {
					t.Errorf("db.Delete returned nil, expected %v", tt.expected)
				}
			}
		})
	}
}

func TestUpdateRound(t *testing.T) {
	db := setupTest(t)
	ctx := context.Background()

	tests := []struct {
		name      string
		sessionId string
		expected  *model.Session
		err       error
	}{
		{
			name:      "session1",
			sessionId: "session1",
			expected: &model.Session{
				Id:       "session1",
				PlayerId: "player1",
				Wins:     3,
				Losses:   1,
				Draws:    0,
			},
			err: nil,
		},
		{
			name:      "session2",
			sessionId: "session2",
			expected: &model.Session{
				Id:       "session2",
				PlayerId: "player2",
				Wins:     2,
				Losses:   2,
				Draws:    0,
			},
			err: nil,
		},
		{
			name:      "session3",
			sessionId: "session3",
			expected: &model.Session{
				Id:       "session3",
				PlayerId: "player3",
				Wins:     1,
				Losses:   0,
				Draws:    3,
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session, err := db.UpdateRoundResult(ctx, tt.sessionId, model.WIN)
			if err != nil && tt.err == nil {
				t.Errorf("db.UpdateRoundResult returned error: %v", err)
			}
			if session != nil {
				if !reflect.DeepEqual(session, tt.expected) {
					t.Errorf("db.UpdateRoundResult returned %v, expected %v", session, tt.expected)
				}
			} else {
				if tt.expected != nil {
					t.Errorf("db.UpdateRoundResult returned nil, expected %v", tt.expected)
				}
			}
		})
	}
}
