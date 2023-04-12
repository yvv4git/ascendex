package ascendex

// Message - used as message implementation.
type Message struct {
	M      string      `json:"m"`
	Symbol string      `json:"symbol,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
