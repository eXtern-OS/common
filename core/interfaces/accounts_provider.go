package interfaces

import "github.com/eXtern-OS/common/core/types"

type AccountsProvider interface {
	Init(params map[string]string) error                    // Inits provider
	AccountExists(credentials string) (bool, error)         // Checks if account exists, by searching for either email, username or userId
	FindAccount(credentials string) (*types.Account, error) // Returns account or error
	InsertAccount(account types.Account) error              // Inserts account or returns an error
	DeleteAccount(credentials string) error                 // Deletes an account or returns an error
	UpdateAccount(account types.Account) error              // Updates an account or returns an error
	Close() error                                           // Closes connection or cleans up
}
