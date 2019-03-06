package demo

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/cernbox/reva/pkg/err"
	"github.com/cernbox/reva/pkg/token"
	"github.com/cernbox/reva/pkg/token/manager/registry"
)

var errors = err.New("demo")

func init() {
	registry.Register("demo", New)
}

// New returns a new token manager.
func New(m map[string]interface{}) (token.Manager, error) {
	mngr := manager{}
	return &mngr, nil
}

type manager struct{}

func (m *manager) MintToken(claims token.Claims) (string, error) {
	token, err := encode(claims)
	if err != nil {
		return "", errors.Wrap(err, "error encoding claims")
	}
	return token, nil
}

func (m *manager) DismantleToken(token string) (token.Claims, error) {
	claims, err := decode(token)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding claims")
	}
	return claims, nil
}

// from https://stackoverflow.com/questions/28020070/golang-serialize-and-deserialize-back
// go binary encoder
func encode(m token.Claims) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

// from https://stackoverflow.com/questions/28020070/golang-serialize-and-deserialize-back
// go binary decoder
func decode(str string) (token.Claims, error) {
	m := token.Claims{}
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
