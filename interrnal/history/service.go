package history

import "context"

type Service interface {
	CreateHistory(ctx context.Context, dto CreateHistoryDTO) *History

	GetHistory(ctx context.Context, limit, offset int) []*History
	GetEntityById(ctx context.Context, id string) *History
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) CreateHistory(ctx context.Context, dto CreateHistoryDTO) *History {
	return nil
}

func (s *service) GetHistory(ctx context.Context, limit, offset int) []*History {
	return s.storage.GetAll(limit, offset)
}

func (s *service) GetEntityById(ctx context.Context, id string) *History {
	return s.storage.GetOne(id)
}
