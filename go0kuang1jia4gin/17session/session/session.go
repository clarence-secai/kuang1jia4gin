package session

type Session interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}
