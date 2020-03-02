package access_token

import (
	"bookstore/bookstore_OAuth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationtime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (token *AccessToken) Validate() *errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if token.AccessToken == "" {
		return errors.NewBadRequestError("invalid Accesstoken")
	}
	if token.UserId <= 0 {
		return errors.NewBadRequestError("invalid userId")

	}
	if token.ClientId <= 0 {
		return errors.NewBadRequestError("invalid clientId")

	}
	if token.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")

	}
	return nil
}
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationtime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
