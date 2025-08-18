package vector

type Store interface {
	Ping() error
	Close() error
}
