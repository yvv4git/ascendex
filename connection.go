package ascendex

// Connection - common contract for all connection implementation.
type Connection interface {
	ReadJSON(interface{}) error
	WriteJSON(interface{}) error
	Close() error
}
