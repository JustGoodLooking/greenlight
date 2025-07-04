package data

import "time"

type keypair struct {
	ID        int64 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int64 `json:"-"`
	Alias     string `json:"alias"`
	KeyType   string `json:"key_type"`
	PublicKey string `json:"public_key"`
	Address   string `json:"address"`
	Status    string `json:"status"`
}
