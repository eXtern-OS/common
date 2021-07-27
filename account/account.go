/*
   Here we will store account type.
*/

package account

import (
	"strconv"
	"time"

	"../utils"
)

/*
   ID - our internal id, UnixNano + 4 random digits
   DisplayName - We are not FAANG, we dont need exact name
   Email
   PasswordHashed
   PfpURL
   AdditionalProperties - thats where things are getting interesting. Thats a key-value map string:string to store additional info about user. For instance, nicknames or stuff.
*/

type Account struct {
	ID                   string            `json:"id"                    bson:"id"`
	DisplayName          string            `json:"display_name"          bson:"display_name"`
	Email                string            `json:"email"                 bson:"email"`
	PasswordHashed       string            `json:"password_hashed"       bson:"password_hashed"`
	PfpURL               string            `json:"pfp_url"               bson:"pfp_url"`
	AdditionalProperties map[string]string `json:"additional_properties" bson:"additional_properties"`
}

// Exported account - for packing into JSON

type ExportedAccount struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	PfPURL string `json:"pf_purl"`
}

// NewAccount returns empty account with set id, and ready-to-go
func NewAccount() Account {
	return Account{
		ID:                   strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(utils.NumberWLen(4)),
		AdditionalProperties: make(map[string]string),
	}
}

// Export returns exported account
func (a *Account) Export() ExportedAccount {
	return ExportedAccount{
		Name:   a.DisplayName,
		Email:  a.Email,
		PfPURL: a.PfpURL,
	}
}

// Verify checks that all inputs are not empty
func (a *Account) Verify() bool {
	return a.Email != "" && a.PasswordHashed != "" && a.DisplayName != ""
}

// SetDiff updates different fields with non-empty values
func (a *Account) SetDiff(ax Account) {
	if ax.Email != "" {
		a.Email = ax.Email
	}

	if ax.DisplayName != "" {
		a.DisplayName = ax.DisplayName
	}

	if ax.PasswordHashed != "" {
		a.PasswordHashed = utils.HashString(ax.PasswordHashed)
	}
}
