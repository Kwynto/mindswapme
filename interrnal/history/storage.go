package history

type Storage interface {
	GetOne(id string) *History
	GetAll(limit, offset int) []*History
	Create(history *History) *History
	Delete(id string) error
}
