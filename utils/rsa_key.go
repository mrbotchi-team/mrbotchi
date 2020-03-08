package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func ReadRSAPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if nil != err {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if nil == block {
		return nil, errors.New("invalid private key")
	}

	var key *rsa.PrivateKey
	switch block.Type {
	case "RSA PRIVATE KEY":
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if nil != err {
			return nil, err
		}

	case "PRIVATE KEY":
		keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if nil != err {
			return nil, err
		}

		var ok bool
		key, ok = keyInterface.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("this key isnt RSA private key")
		}

	default:
		return nil, fmt.Errorf("invalid private key type: %s", block.Type)
	}

	key.Precompute()

	if err := key.Validate(); nil != err {
		return nil, err
	}

	return key, nil
}

func ReadRSAPublicKey(filePath string) (*rsa.PublicKey, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if nil != err {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if nil == block {
		return nil, errors.New("invalid public key")
	}

	if "PUBLIC KEY" != block.Type {
		return nil, fmt.Errorf("invalid public key type: %s", block.Type)
	}

	keyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if nil != err {
		return nil, err
	}

	key, ok := keyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("this key isnt RSA public key")
	}

	return key, nil

}
