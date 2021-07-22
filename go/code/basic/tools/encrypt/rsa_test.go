package encrypt

import "testing"

func TestGetRsaKey(t *testing.T) {
	err := GetRsaKey()
	if err != nil {
		t.Error(err)
	}
}
