package model

import "time"

// UserData maintains all of the user model related data.
// Credentials are stored differently
type UserData struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Timezone  time.Location // Localization, this will adapt all transactions to the user timezone. All transactions are stored in UTC.
}

type UserCredentials struct {
	Id           int // User Id
	Email        string
	PasswordHash string
}
