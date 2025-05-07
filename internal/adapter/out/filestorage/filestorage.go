package filestorage

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/felipeversiane/donation-server/config"
)

var (
	once     sync.Once
	instance *fileStorage
)

type fileStorage struct {
	config        config.FileStorageConfig
	basePath      string
	permissions   os.FileMode
	isInitialized bool
	mutex         sync.RWMutex
}

type Interface interface {
	SaveFile(ctx context.Context, path string, content []byte) error
	ReadFile(ctx context.Context, path string) ([]byte, error)
	DeleteFile(ctx context.Context, path string) error
	Exists(ctx context.Context, path string) (bool, error)
	HealthCheck(ctx context.Context) error
	Close()
}

func New(config config.FileStorageConfig) (Interface, error) {
	var err error
	once.Do(func() {
		slog.Info("initializing local storage...")
		basePath, pathErr := filepath.Abs(config.BasePath)
		if pathErr != nil {
			err = fmt.Errorf("failed to resolve absolute path: %w", pathErr)
			slog.Error("error resolving storage path", "error", err)
			return
		}

		if mkdirErr := os.MkdirAll(basePath, os.FileMode(config.DirectoryPermissions)); mkdirErr != nil {
			err = fmt.Errorf("failed to create storage directory: %w", mkdirErr)
			slog.Error("error creating storage directory", "error", err)
			return
		}

		instance = &fileStorage{
			config:        config,
			basePath:      basePath,
			permissions:   os.FileMode(config.FilePermissions),
			isInitialized: true,
		}

		slog.Info("attempting to verify storage accessibility")

		if err = instance.HealthCheck(context.Background()); err != nil {
			instance.Close()
			err = fmt.Errorf("failed to initialize storage: %w", err)
			slog.Error("error initializing storage", "error", err)
		} else {
			slog.Info("local storage initialized successfully")
		}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (s *fileStorage) SaveFile(ctx context.Context, path string, content []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.isInitialized {
		return fmt.Errorf("storage not initialized")
	}

	fullPath := s.getFullPath(path)
	dir := filepath.Dir(fullPath)

	if err := os.MkdirAll(dir, os.FileMode(s.config.DirectoryPermissions)); err != nil {
		slog.Error("failed to create directory", "path", dir, "error", err)
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(fullPath, content, s.permissions); err != nil {
		slog.Error("failed to write file", "path", fullPath, "error", err)
		return fmt.Errorf("failed to write file: %w", err)
	}

	slog.Info("file saved successfully", "path", fullPath)
	return nil
}

func (s *fileStorage) ReadFile(ctx context.Context, path string) ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.isInitialized {
		return nil, fmt.Errorf("storage not initialized")
	}

	fullPath := s.getFullPath(path)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		slog.Error("failed to read file", "path", fullPath, "error", err)
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return content, nil
}

func (s *fileStorage) DeleteFile(ctx context.Context, path string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.isInitialized {
		return fmt.Errorf("storage not initialized")
	}

	fullPath := s.getFullPath(path)
	if err := os.Remove(fullPath); err != nil {
		slog.Error("failed to delete file", "path", fullPath, "error", err)
		return fmt.Errorf("failed to delete file: %w", err)
	}

	slog.Info("file deleted successfully", "path", fullPath)
	return nil
}

func (s *fileStorage) Exists(ctx context.Context, path string) (bool, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.isInitialized {
		return false, fmt.Errorf("storage not initialized")
	}

	fullPath := s.getFullPath(path)
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		slog.Error("failed to check file existence", "path", fullPath, "error", err)
		return false, fmt.Errorf("failed to check file existence: %w", err)
	}

	return true, nil
}

func (s *fileStorage) HealthCheck(ctx context.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.isInitialized {
		return fmt.Errorf("storage not initialized")
	}

	testPath := s.getFullPath(fmt.Sprintf("healthcheck_%d.tmp", time.Now().UnixNano()))
	if err := os.WriteFile(testPath, []byte("healthcheck"), s.permissions); err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}

	if err := os.Remove(testPath); err != nil {
		slog.Warn("failed to clean up health check file", "path", testPath, "error", err)
	}

	return nil
}

func (s *fileStorage) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.isInitialized {
		s.isInitialized = false
		slog.Info("local storage closed")
	}
}

func (s *fileStorage) getFullPath(path string) string {
	return filepath.Join(s.basePath, filepath.Clean(path))
}
