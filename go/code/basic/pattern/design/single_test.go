package design

import "testing"

func TestSingle(t *testing.T) {
    p := Init("john")
    j := Init("joe")
    if p.GetName() != j.GetName() {
        t.Error("none single")
    }
}

func TestSingleImpl_GetTitle(t *testing.T) {
    a := NewSingleImpl("apple")
    b := NewSingleImpl("bob")
    if a.GetTitle() != b.GetTitle() {
        t.Error("none single Impl")
    }
    t.Log(a.GetTitle())
}
