package main

import "testing"

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Channel()
	}
	b.Log(b.Name(), "channel benchmark finish")

}

func BenchmarkWG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WG()
	}
}

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

func BenchmarkChannel2(b *testing.B) {

	Channel()
}

//func TestChannel(t *testing.T) {
//	if "" != Channel() {
//		t.Error("fail")
//	}
//
//	t.Log("success")
//}
