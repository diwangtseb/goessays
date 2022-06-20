package channel

import (
	"testing"

	"github.com/bytedance/sonic"
)

func BenchmarkUserUnMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sonic.Unmarshal([]byte(`{"id":1,"name":"test"}`), &User{})
	}
}

func BenchmarkUserUnMarshalWithPool(b *testing.B) {
	u := NewUserPool()
	for n := 0; n < b.N; n++ {
		user := u.Get().(*User)
		sonic.Unmarshal([]byte(`{"id":1,"name":"test"}`), user)
		u.Put(user)
	}
}
