/*
	The purpose of this package is to define token type and provide utilities to create it
*/
package token

import (
	"../utils"
	"time"
)

/*
	Token is a special structure, which simplifies cross-application interaction
	On the outside, token is a string, which begins with t0, has to be passed through applications. Token represents account id, but not to leak it,
	We use temporary tokens. 24 hours is a reasonable timeout https://external-preview.redd.it/SFOGfkqV63OQyQAJKk3727MlLO7WIvGJaf7ylKteAbQ.jpg?auto=webp&s=ff8bd90246ceb10bf38203d419bf3b06590a3835
*/
type Token struct {
	Created int64  `json:"created"  bson:"created"`
	Expires int64  `json:"expires"  bson:"expires"`
	TokenID string `json:"token_id" bson:"token_id"`
	UserID  string `json:"user_id"  bson:"user_id"`
}

/*
	ExportedToken will be used to communicate between servers
*/
type ExportedToken struct {
	TokenID string `json:"token_id"`
	UserID  string `json:"user_id"`
}

// Export converts Token to ExportedToken
func (t *Token) Export() ExportedToken {
	return ExportedToken{
		TokenID: t.TokenID,
		UserID:  t.UserID,
	}
}

// NewToken creates new Token, sets expiration to 24 hours
func NewToken(userID string) Token {
	t := time.Now()
	return Token{
		Created: t.Unix(),
		Expires: t.Add(24 * time.Hour).Unix(),
		TokenID: "t0" + utils.HashInt64(t.UnixNano()+int64(utils.NumberWLen(8))),
		UserID:  userID,
	}
}

// Verify verifies Token on the timeout basis, and checks whether it has empty userId
func (t *Token) Verify() bool {
	return time.Now().Unix() > t.Expires && t.UserID != ""
}
