package model

import (
	"testing"
)

func Test_mysql(t *testing.T) {
	want := "root:root@tcp(localhost:3306)/forum?charset=utf8&parseTime=True&loc=Local"
	if got := mysql(); got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
