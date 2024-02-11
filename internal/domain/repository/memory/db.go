package memory

import (
	"context"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/model"
	"github.com/ksaritek/paper-rock-scissors/internal/domain/repository"
	repositoryerr "github.com/ksaritek/paper-rock-scissors/internal/domain/repository/errors"
	"github.com/ksaritek/paper-rock-scissors/internal/observability"
	"go.uber.org/zap"
)

type memoryStore struct {
	sync.RWMutex
	data     map[string]model.Session
	ttl      time.Duration
	expireAt map[string]time.Time
	logger   *zap.Logger
}

func NewMemoryStore(ctx context.Context, l *zap.Logger, ttl time.Duration) repository.Session {
	ctx, span := observability.Tracer().Start(ctx, "memory_db_new")
	defer span.End()

	db := &memoryStore{
		data:     make(map[string]model.Session),
		ttl:      ttl,
		expireAt: make(map[string]time.Time),
		logger:   l.With(zap.String("component", "sessionMap")),
	}

	go func(ctx context.Context, db *memoryStore) {
		ctx, span := observability.Tracer().Start(ctx, "memory_db_get")
		defer span.End()

		db.logger.Info("starting session cleanup routine")

		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				db.logger.Info("stopping session cleanup routine")
				return
			case <-ticker.C:
				db.logger.Debug("running session cleanup routine")
				db.deleteExpired()
			}
		}
	}(ctx, db)

	return db
}

func (m *memoryStore) Get(ctx context.Context, sessionId string) (_ *model.Session, retErr error) {
	_, span := observability.Tracer().Start(ctx, "memory_db_get")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	m.RLock()
	defer m.RUnlock()

	session, ok := m.data[sessionId]
	if !ok {
		return nil, errors.Wrap(repositoryerr.ErrNotFound, "session not found")
	}

	m.expireAt[session.Id] = time.Now().Add(m.ttl)
	return &session, nil
}

func (m *memoryStore) Create(ctx context.Context, session model.Session) (retErr error) {
	_, span := observability.Tracer().Start(ctx, "memory_db_create")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	m.Lock()
	defer m.Unlock()

	m.data[session.Id] = session
	m.expireAt[session.Id] = time.Now().Add(m.ttl)

	return nil
}

func (m *memoryStore) UpdateRoundResult(ctx context.Context, sessionId string, player model.RoundResult) (_ *model.Session, retErr error) {
	_, span := observability.Tracer().Start(ctx, "memory_db_update")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	m.Lock()
	defer m.Unlock()

	session, ok := m.data[sessionId]
	if !ok {
		return nil, errors.Wrap(repositoryerr.ErrNotFound, "session not found")
	}

	switch player {
	case model.WIN:
		session.Wins++
	case model.LOSS:
		session.Losses++
	case model.DRAW:
		session.Draws++
	}

	m.data[session.Id] = session
	m.expireAt[session.Id] = time.Now().Add(m.ttl)

	return &session, nil
}

func (m *memoryStore) Delete(ctx context.Context, sessionId string) (_ *model.Session, retErr error) {
	_, span := observability.Tracer().Start(ctx, "memory_db_delete")
	defer span.End()
	defer observability.RecordErr(span, retErr)

	m.Lock()
	defer m.Unlock()

	session, ok := m.data[sessionId]
	if !ok {
		return nil, errors.Wrap(repositoryerr.ErrNotFound, "session not found")
	}

	delete(m.data, sessionId)
	delete(m.expireAt, sessionId)

	return &session, nil
}

func (m *memoryStore) deleteExpired() {
	m.Lock()
	defer m.Unlock()

	for sessionId := range m.data {
		if m.isExpired(sessionId) {
			m.logger.Debug("deleting expired session", zap.String("session_id", sessionId))

			delete(m.data, sessionId)
			delete(m.expireAt, sessionId)
		}
	}

}

func (m *memoryStore) isExpired(key string) bool {
	expireAt, ok := m.expireAt[key]
	if !ok {
		return false
	}

	return time.Now().After(expireAt)
}
