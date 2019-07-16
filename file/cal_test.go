package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := add(1, 2)
	if res != 3 {
		t.Fatalf("失败了")
	} else {
		t.Logf("ok")
	}
}

func TestSum(t *testing.T) {
	res := sum(1, 2)
	if res == 5 {
		t.Logf("ok")
	} else {
		t.Fatalf("失败了")
	}
}

func TestStoreMonster(t *testing.T) {
	m := Monster{"213", 1, "123"}
	if m.Store() {
		t.Logf("ok")
	} else {
		t.Fatalf("失败了")
	}
}

func TestReStoreMonster(t *testing.T) {
	m := Monster{}
	if m.ReStore("./Monster.txt") {
		t.Logf("123123    %v", m)
	} else {
		t.Fatalf("失败了")
	}
}
