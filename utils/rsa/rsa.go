package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

type Rsa struct {
	PrivateKey []byte `json:"private_key"`
	PublicKey  []byte `json:"public_key"`
}

func Gen() (*Rsa, error) {
	pvKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(pvKey)

	derPkix, err := x509.MarshalPKIXPublicKey(&pvKey.PublicKey)
	if err != nil {
		return nil, err
	}

	return &Rsa{
		PrivateKey: pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: derStream,
		}),
		PublicKey: pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: derPkix,
		}),
	}, nil
}

// Encrypt public key encrypt
func (r *Rsa) Encrypt(s string) (string, error) {
	if block, err := pemDecode(r.PublicKey); err != nil {
		return "", err
	} else {
		pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return "", err
		}
		pub := pubInterface.(*rsa.PublicKey)
		encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(s))
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(encrypted), nil
	}
}

// Decrypt private key decrypt
func (r *Rsa) Decrypt(s string) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	if block, err := pemDecode(r.PrivateKey); err != nil {
		return "", err
	} else {
		private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, private, encrypted)
		if err != nil {
			return "", err
		}
		return string(decrypted), nil
	}
}

// Signature private key signature
func (r *Rsa) Signature(decrypted []byte) ([]byte, error) {
	rng := rand.Reader
	hashed := sha256.Sum256(decrypted)
	if block, err := pemDecode(r.PrivateKey); err != nil {
		return nil, err
	} else {
		private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		signature, err := rsa.SignPKCS1v15(rng, private, crypto.SHA256, hashed[:])
		if err != nil {
			return nil, err
		}
		return signature, nil
	}
}

// SignatureVerify public key verify
func (r *Rsa) SignatureVerify(encrypted []byte, signature []byte) error {
	hashed := sha256.Sum256(encrypted)
	if block, err := pemDecode(r.PublicKey); err != nil {
		return err
	} else {
		pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return err
		}
		pub := pubInterface.(*rsa.PublicKey)
		return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
	}
}

func pemDecode(key []byte) (*pem.Block, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("key error")
	}
	return block, nil
}
