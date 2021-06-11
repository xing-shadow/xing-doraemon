package RuleEngine

import (
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/storage/tsdb"
	"io/ioutil"
	"os"
	"time"
)

//MockStorage temporary storage
type MockStorage struct {
	storage.Storage
	dir string
}

// Close delete temporary files when closed
func (s MockStorage) Close() error {
	if err := s.Storage.Close(); err != nil {
		return err
	}
	return os.RemoveAll(s.dir)
}

// NewMockStorage create temporary storage
func NewMockStorage() (storage.Storage, error) {
	dir, err := ioutil.TempDir("", "mock_storage")
	if err != nil {
		return nil, err
	}
	db, err := tsdb.Open(dir, nil, nil, &tsdb.Options{
		MinBlockDuration:  model.Duration(time.Hour * 2),
		MaxBlockDuration:  model.Duration(time.Hour * 24),
		RetentionDuration: model.Duration(15 * 24 * time.Hour),
	})
	if err != nil {
		return nil, err
	}

	return MockStorage{Storage: tsdb.Adapter(db, int64(0)), dir: dir}, nil
}
