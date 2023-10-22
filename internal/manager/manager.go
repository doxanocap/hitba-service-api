package manager

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/hitba-service-api/pkg/redis"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
	"sync"
)

type Manager struct {
	cacheConn *redis.Conn
	cfg       *model.Config
	db        *gorm.DB

	service       interfaces.IService
	serviceRunner sync.Once

	repository       interfaces.IRepository
	repositoryRunner sync.Once

	processor       interfaces.IProcessor
	processorRunner sync.Once
}

func InitManager(cfg *model.Config) *Manager {
	return &Manager{
		cfg: cfg,
	}
}

func (m *Manager) Repository() interfaces.IRepository {
	m.repositoryRunner.Do(func() {
		m.repository = InitRepositoryManager(m.db)
	})
	return m.repository
}

func (m *Manager) Service() interfaces.IService {
	m.serviceRunner.Do(func() {
		m.service = InitServiceManager(m)
	})
	return m.service
}

func (m *Manager) Processor() interfaces.IProcessor {
	m.processorRunner.Do(func() {
		m.processor = InitProcessor(m, m.cfg)
	})
	return m.processor
}

func (m *Manager) SetCoreDB(db *gorm.DB) {
	m.db = db
}

func (m *Manager) SetCacheConnection(cacheConn *redis.Conn) {
	m.cacheConn = cacheConn
}
