package data

import (
	"database/sql"
	"time"
)

type Keypair struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int64     `json:"-"`
	Alias     string    `json:"alias"`
	KeyType   string    `json:"key_type"`
	PublicKey string    `json:"public_key"`
	Address   string    `json:"address"`
	Status    string    `json:"status"`
}

type KeypairModel struct {
	DB *sql.DB
}


func (m KeypairModel) New(userID int64, alias, keyType string) (*Keypair, any, error) {
	keypair, pri := generateKeypair(userID, alias, keyType)

	return keypair, pri, nil
}


func generateKeypair(userID int64, alias, keyType string) (*Keypair, any) {
	
}