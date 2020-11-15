package cache

const (
	DefaultLockSecondsDuration = 10
)

type ParallelLockControl interface {
	CreateLock(key string, secondsDuration int) error
	ReleaseLock(key string) error
	IsLocked(key string) (bool, error)
}
