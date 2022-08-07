package memory

import (
	"github.com/alimgiray/getir/adapter/sqlite"
	"gorm.io/gorm"
)

type MemoryService struct {
	db *gorm.DB
}

func NewMemoryService() *MemoryService {
	return &MemoryService{db: sqlite.Connect(&Record{})}
}

func (s *MemoryService) CreateNewRecord(record *Record) error {
	return s.db.Save(&record).Error
}

func (s *MemoryService) FindRecord(key string) (Record, error) {
	record := Record{Key: key}
	err := s.db.First(&record).Error
	return record, err
}
