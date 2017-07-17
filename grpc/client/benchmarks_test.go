package main

import "testing"

func BenchmarkGetOrders(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, e := GetOrders("127.0.0.1:8089")
		if e != nil {
			b.Fatal("failed to connect order server: ", e)
		}
		if r.GetRslt().GetRslt() != 0 {
			b.Fatal("failed to get orders: ", r.GetRslt().GetDescription())
		}
	}

}
