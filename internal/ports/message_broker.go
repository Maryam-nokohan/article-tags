package ports

type MessageBroker interface {
    Publish(subject string, data []byte) error
    Subscribe(subject string, handler func([]byte) error) error
    Close() error
}
