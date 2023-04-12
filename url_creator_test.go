package ascendex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAscendexURL_Create(t *testing.T) {
	{
		// Default use case.
		urlCreator := NewAscendexURL()
		result := urlCreator.Create()
		require.Equal(t, "wss://ascendex.com/0/api/pro/v1/stream", result)
	}
	{
		// Setup host.
		urlCreator := NewAscendexURL()
		urlCreator.SetHost("ws://localhost:8080")
		result := urlCreator.Create()
		require.Equal(t, "ws://localhost:8080/0/api/pro/v1/stream", result)
	}
	{
		// Setup account group.
		urlCreator := NewAscendexURL()
		urlCreator.SetAccountGroup("1")
		result := urlCreator.Create()
		require.Equal(t, "wss://ascendex.com/1/api/pro/v1/stream", result)
	}
	{
		// Setup path.
		urlCreator := NewAscendexURL()
		urlCreator.SetPath("api/pro/v2/stream")
		result := urlCreator.Create()
		require.Equal(t, "wss://ascendex.com/0/api/pro/v2/stream", result)
	}
	{
		// Try setup empty host.
		urlCreator := NewAscendexURL()
		urlCreator.SetHost("")
		result := urlCreator.Create()
		require.Equal(t, "wss://ascendex.com/0/api/pro/v1/stream", result)
	}
	{
		// Try setup empty account group.
		urlCreator := NewAscendexURL()
		urlCreator.SetAccountGroup("")
		result := urlCreator.Create()
		require.Equal(t, "wss://ascendex.com/0/api/pro/v1/stream", result)
	}
	{
		// Try setup empty path.
		urlCreator := NewAscendexURL()
		urlCreator.SetPath("")
		result := urlCreator.Create()
		require.Equal(t, "wss://ascendex.com/0/api/pro/v1/stream", result)
	}
}
