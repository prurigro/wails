package store_test

import (
	"github.com/matryer/is"
	"github.com/wailsapp/wails/v2/pkg/store"
	"testing"
)

func TestNew(t *testing.T) {
	is2 := is.New(t)
	intStore := store.New("test", 1)
	is2.Equal(intStore.Get(), 1)

	stringStore := store.New("test", "test")

	var subscriberValue string
	stringStore.Subscribe(func(data string) {
		subscriberValue = data
	})

	is2.Equal(stringStore.Get(), "test")
	stringStore.Set("test2")
	is2.Equal(stringStore.Get(), "test2")
	is2.Equal(subscriberValue, "test2")

	stringStore.Update(func(s string) string {
		return s + "2"
	})
	is2.Equal(stringStore.Get(), "test22")
	is2.Equal(subscriberValue, "test22")

}
