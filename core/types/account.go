/*
   Here we will store account type.
*/

package types

import (
	"strconv"
	"time"

	"github.com/eXtern-OS/common/utils"
)

// Account implements user account for eXtern OS services
type Account struct {
	ID                   string            `json:"id"                    bson:"id"`                    //Internal ID
	DisplayName          string            `json:"display_name"          bson:"display_name"`          // User-defined name displayed publicly
	Username             string            `json:"username" bson:"username"`                           // Username, displayed publicly
	Email                string            `json:"email"                 bson:"email"`                 // Email, could be displayed publicly
	PasswordHashed       string            `json:"password_hashed"       bson:"password_hashed"`       // Hashed password stored in DB
	PfpURL               string            `json:"pfp_url"               bson:"pfp_url"`               // URL to user's PFP, PFP displayed publicly
	AdditionalProperties map[string]string `json:"additional_properties" bson:"additional_properties"` // Additional properties, some of which are internal, some could be exported
}

// ExportedAccount - for packing Account into JSON
type ExportedAccount struct {
	Name               string            `json:"name"`
	Username           string            `json:"username"`
	Email              string            `json:"email"`
	PfPURL             string            `json:"pf_purl"`
	ExportedProperties map[string]string `json:"exported_properties"`
}

// NewAccount returns empty account with set id, and ready-to-go
func NewAccount() Account {
	return Account{
		ID:                   strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(utils.NumberWLen(4)),
		AdditionalProperties: make(map[string]string),
	}
}

// AllowedForExport is a list of packages allowed for export
var AllowedForExport = []string{"packages"}

func exportProperties(properties map[string]string) map[string]string {
	res := make(map[string]string)
	for _, allowed := range AllowedForExport {
		if v, ok := properties[allowed]; ok {
			res[allowed] = v
		}
	}
	return res
}

// Export returns exported account
func (a *Account) Export() ExportedAccount {
	return ExportedAccount{
		Name:               a.DisplayName,
		Email:              a.Email,
		PfPURL:             a.PfpURL,
		Username:           a.Username,
		ExportedProperties: exportProperties(a.AdditionalProperties),
	}
}

// Verify checks that all inputs are not empty
func (a *Account) Verify() bool {
	return a.Email != "" && a.PasswordHashed != "" && a.DisplayName != "" && a.Username != ""
}

// SetDiff updates different fields with non-empty values
func (a *Account) SetDiff(ax Account) {
	a.Email = utils.Ternary(ax.Email != "", ax.Email, a.Email)

	a.DisplayName = utils.Ternary(ax.DisplayName != "", ax.DisplayName, a.DisplayName)

	a.PasswordHashed = utils.Ternary(ax.PasswordHashed != "", ax.PasswordHashed, a.PasswordHashed)

	a.PfpURL = utils.Ternary(ax.PfpURL != "", ax.PfpURL, a.PfpURL)

	a.Username = utils.Ternary(ax.Username != "", ax.Username, a.Username)
}
