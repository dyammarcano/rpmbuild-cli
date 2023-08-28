package static

import "testing"

func TestResources(t *testing.T) {
	data, err := Resources()
	if err != nil {
		t.Errorf("Resources() = %v; want %v", err, nil)
	}

	if data == nil {
		t.Errorf("Resources() = %v; want %v", data, nil)
	}
}
