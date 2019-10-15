package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/mrbotchi-team/mrbotchi/config"
	"golang.org/x/crypto/argon2"
)

var (
	InvalidHashError         = errors.New("The encoded hash isn't in the correct format.")
	IncompatibleVersionError = errors.New("Incompatible version of argon2.")
)

func GenerateHashedPassword(password string, param config.Argon2Config) (string, error) {
	salt, err := GenerateRandomBytes(param.SaltLength)
	if nil != err {
		return "", err
	}
	hashedPassword := argon2.IDKey([]byte(password), salt, param.Iterations, param.Memory, param.Parallelism, param.KeyLength)

	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hashedPassword)

	argon2Hash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, param.Memory, param.Iterations, param.Parallelism, base64Salt, base64Hash)

	return argon2Hash, nil
}

func VerifyPassword(password, encodedHash string) (bool, error) {
	param, salt, hash, err := decodeHash(encodedHash)
	if nil != err {
		return false, err
	}
	otherHash := argon2.IDKey([]byte(password), salt, param.Iterations, param.Memory, param.Parallelism, param.KeyLength)

	if 1 == subtle.ConstantTimeCompare(hash, otherHash) {
		return true, nil
	}

	return false, nil
}

func decodeHash(encodedHash string) (*config.Argon2Config, []byte, []byte, error) {
	vals := strings.Split(encodedHash, "$")
	if 6 != len(vals) {
		return nil, nil, nil, InvalidHashError
	}

	var version int
	if _, err := fmt.Sscanf(vals[2], "v=%d", &version); nil != err {
		return nil, nil, nil, err
	}
	if argon2.Version != version {
		return nil, nil, nil, IncompatibleVersionError
	}

	param := &config.Argon2Config{}
	if _, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &param.Memory, &param.Iterations, &param.Parallelism); nil != err {
		return nil, nil, nil, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	param.SaltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	param.KeyLength = uint32(len(hash))

	return param, salt, hash, nil
}
