package storage_type

type StorageType int

const (
	CSV StorageType = iota
	SQLite
)

func (s StorageType) String() string {
	switch s {
	case CSV:
		return "csv"
	case SQLite:
		return "sqlite"
	default:
		return "unknown"
	}
}
