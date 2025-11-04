package security

import (
    "crypto/rand"
    "encoding/base64"
    "errors"

    "golang.org/x/crypto/argon2"
)

const (
    timeCost    = 1
    memoryCost  = 64 * 1024
    threads     = 4
    keyLength   = 32
)

func HashPassword(password string) (string, error) {
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, threads, keyLength)
    encoded := base64.RawStdEncoding.EncodeToString(append(salt, hash...))
    return encoded, nil
}

func VerifyPassword(encoded, password string) error {
    data, err := base64.RawStdEncoding.DecodeString(encoded)
    if err != nil {
        return err
    }
    if len(data) < 16 {
        return errors.New("invalid hash data")
    }
    salt := data[:16]
    hash := data[16:]
    test := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, threads, uint32(len(hash)))
    for i := range hash {
        if hash[i] != test[i] {
            return errors.New("invalid password")
        }
    }
    return nil
}
