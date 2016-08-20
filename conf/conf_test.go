package conf

import (
	"testing"
)

func TestConf(t *testing.T) {
	if RETRY_TIME != 5 {
		t.Errorf("RETRY_TIME expected %d, got %d", 5, RETRY_TIME)
	}

	if ID_CHAN_SIZE != 10 {
		t.Errorf("ID_CHAN_SIZE expected %d, got %d", 10, ID_CHAN_SIZE)
	}

	if THREAD_POOL_SIZE != 5 {
		t.Errorf("THREAD_POOL_SIZE expected %d, got %d", 5, THREAD_POOL_SIZE)
	}

	if l := len(TOR_URL_TEMPLATES); l != 4 {
		t.Errorf("len(TOR_URL_TEMPLATES) expected %d, got %d", 4, l)
	}

	if l := len(IMG_URL_TEMPLATES); l != 1 {
		t.Errorf("len(IMG_URL_TEMPLATES) expected %d, got %d", 1, l)
	}
}
