package easyjson_example

import (
	"encoding/json"
	"testing"
	"time"
)

func BenchmarkEasyJSONMarshal(b *testing.B) {
	user := &User{
		ID:        1,
		Name:      "Иван Иванов",
		Email:     "ivan@example.com",
		CreatedAt: time.Now(),
		Tags:      []string{"developer", "golang"},
		Profile: &Profile{
			Bio:      "Go разработчик",
			Settings: map[string]string{"theme": "dark", "lang": "ru"},
			IsActive: true,
			Score:    95.5,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = user.MarshalJSON()
	}
}

func BenchmarkStandardJSONMarshal(b *testing.B) {
	user := &User{
		ID:        1,
		Name:      "Иван Иванов",
		Email:     "ivan@example.com",
		CreatedAt: time.Now(),
		Tags:      []string{"developer", "golang"},
		Profile: &Profile{
			Bio:      "Go разработчик",
			Settings: map[string]string{"theme": "dark", "lang": "ru"},
			IsActive: true,
			Score:    95.5,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(user)
	}
}

func BenchmarkEasyJSONUnmarshal(b *testing.B) {
	jsonData := []byte(`{"id":1,"name":"Иван Иванов","email":"ivan@example.com","created_at":"2024-01-01T00:00:00Z","tags":["developer","golang"],"profile":{"bio":"Go разработчик","settings":{"theme":"dark","lang":"ru"},"is_active":true,"score":95.5}}`)
	var user User

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = user.UnmarshalJSON(jsonData)
	}
}

func BenchmarkStandardJSONUnmarshal(b *testing.B) {
	jsonData := []byte(`{"id":1,"name":"Иван Иванов","email":"ivan@example.com","created_at":"2024-01-01T00:00:00Z","tags":["developer","golang"],"profile":{"bio":"Go разработчик","settings":{"theme":"dark","lang":"ru"},"is_active":true,"score":95.5}}`)
	var user User

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(jsonData, &user)
	}
}
