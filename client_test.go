package ascendex

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/yvv4git/ascendex/mock"
)

func TestReadMessagesFromChannel_ConnectionNotInitialized(t *testing.T) {
	urlCreator := NewAscendexURL()
	client := NewClient(urlCreator)

	require.Nil(t, client.conn, "must be nil")
	require.ErrorIs(t, ErrConnectionClosed, client.SubscribeToChannel("BTC_USDT"), "must return error")
}

func TestReadMessagesFromChannel_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urlCreator := NewAscendexURL()
	client := NewClient(urlCreator)

	mockConn := mock.NewMockConnection(ctrl)
	mockConn.EXPECT().WriteJSON(gomock.Any()).AnyTimes().DoAndReturn(func(v interface{}) error {
		val, ok := v.(map[string]string)
		require.True(t, ok)
		require.Equal(t, "bbo:BTC/USDT", val["ch"])

		return nil
	})
	mockConn.EXPECT().ReadJSON(gomock.Any()).AnyTimes().DoAndReturn(func(v interface{}) error {
		err := json.Unmarshal([]byte(`{"m":"bbo","symbol":"BTC/USDT","data":{"ts":1573067212324,"bid":["29903","0.0033"],"ask":["30186","0.0017"]}}`), v)
		require.NoError(t, err)

		return nil
	})
	client.conn = mockConn

	require.NoError(t, client.SubscribeToChannel("BTC_USDT"))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	chBestOrderBook := make(chan BestOrderBook)
	go client.ReadMessagesFromChannel(chBestOrderBook)

	select {
	case <-ctx.Done():
		t.Fail()
	case <-chBestOrderBook:
		return
	}
}

func TestReadMessagesFromChannel_InvalidMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urlCreator := NewAscendexURL()
	client := NewClient(urlCreator)

	mockConn := mock.NewMockConnection(ctrl)
	mockConn.EXPECT().WriteJSON(gomock.Any()).AnyTimes().DoAndReturn(func(v interface{}) error {
		val, ok := v.(map[string]string)
		require.True(t, ok)
		require.Equal(t, "bbo:BTC/USDT", val["ch"])

		return nil
	})
	mockConn.EXPECT().ReadJSON(gomock.Any()).AnyTimes().DoAndReturn(func(v interface{}) error {
		err := json.Unmarshal([]byte(`{"symbol":"BTC/USDT","data":{"ts":1573067212324,"bid":["29903","0.0033"],"ask":["30186","0.0017"]}}`), v)
		require.NoError(t, err)

		return nil
	})
	client.conn = mockConn

	require.NoError(t, client.SubscribeToChannel("BTC_USDT"))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	chBestOrderBook := make(chan BestOrderBook)
	go client.ReadMessagesFromChannel(chBestOrderBook)

	select {
	case <-ctx.Done():
		return
	case <-chBestOrderBook:
		t.Fail()
	}
}

func TestSubscribeToChannel_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urlCreator := NewAscendexURL()
	client := NewClient(urlCreator)

	mockConn := mock.NewMockConnection(ctrl)
	mockConn.EXPECT().WriteJSON(gomock.Any()).AnyTimes().DoAndReturn(func(v interface{}) error {
		val, ok := v.(map[string]string)
		require.True(t, ok)
		require.Equal(t, "bbo:BTC/USDT", val["ch"])

		return nil
	})
	client.conn = mockConn

	require.NoError(t, client.SubscribeToChannel("BTC_USDT"))
}

func TestSubscribeToChannel_InvalidParameters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urlCreator := NewAscendexURL()
	client := NewClient(urlCreator)

	mockConn := mock.NewMockConnection(ctrl)
	mockConn.EXPECT().WriteJSON(gomock.Any()).AnyTimes().DoAndReturn(func(v interface{}) error {
		val, ok := v.(map[string]string)
		require.True(t, ok)
		require.Equal(t, "bbo:BTC/USDT", val["ch"])

		return nil
	})
	client.conn = mockConn

	require.ErrorIs(t, client.SubscribeToChannel("BTC-USDT"), ErrInvalidSymbolValue)
}

func TestSubscribeToChannel_ProblemWithConnection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urlCreator := NewAscendexURL()
	client := NewClient(urlCreator)

	mockConn := mock.NewMockConnection(ctrl)
	mockConn.EXPECT().WriteJSON(gomock.Any()).AnyTimes().DoAndReturn(func(v interface{}) error {
		val, ok := v.(map[string]string)
		require.True(t, ok)
		require.Equal(t, "bbo:BTC/USDT", val["ch"])

		return errors.New("something's wrong")
	})
	mockConn.EXPECT().Close()
	client.conn = mockConn

	require.Error(t, client.SubscribeToChannel("BTC_USDT"))
}

func TestWriteMessagesToChannel_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urlCreator := NewAscendexURL()
	client := NewClient(urlCreator)

	mockConn := mock.NewMockConnection(ctrl)
	mockConn.EXPECT().WriteJSON(gomock.Any()).AnyTimes().DoAndReturn(func(v interface{}) error {
		val, ok := v.(map[string]string)
		require.True(t, ok)
		require.Equal(t, "ping", val["op"])

		return errors.New("something's wrong")
	})

	client.conn = mockConn

	client.WriteMessagesToChannel()
}

func TestWriteMessagesToChannel_NilClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urlCreator := NewAscendexURL()
	client := NewClient(urlCreator)

	client.WriteMessagesToChannel()
}
