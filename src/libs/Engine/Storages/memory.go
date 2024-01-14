package Storages

type MemoryStorage struct {
	st map[string]map[int]bool
}

func NewMemoryStorage(capacity int) *MemoryStorage {
	return &MemoryStorage{
		st: make(map[string]map[int]bool, capacity),
	}
}

func (storage *MemoryStorage) Get(text string) *[]int {
	res := make([]int, 0, len(storage.st[text]))
	for k := range storage.st[text] {
		res = append(res, k)
	}
	return &res
}

func (storage *MemoryStorage) Set(id int, text string) {
	if storage.st[text] == nil {
		storage.st[text] = make(map[int]bool, 30)
	}
	storage.st[text][id] = true
}
