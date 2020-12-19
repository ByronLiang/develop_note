package design

import "testing"

func TestSingle(t *testing.T) {
    p := Init("john")
    j := Init("joe")
    if p.GetName() != j.GetName() {
        t.Error("none single")
    }
}
