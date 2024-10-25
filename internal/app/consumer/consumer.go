package consumer

// Consumer defines an interface with a Start method, which is expected to start the
// event consumption process and return an error if there are any issues
type Consumer interface {
	Start() error
}