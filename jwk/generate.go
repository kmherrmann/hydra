package jwk

import (
	"context"
	"crypto/x509"

	"github.com/gofrs/uuid"

	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"

	"github.com/ory/x/josex"
)

func GenerateJWK(ctx context.Context, alg jose.SignatureAlgorithm, kid, use string) (*jose.JSONWebKeySet, error) {
	bits := 0
	if alg == jose.RS256 || alg == jose.RS384 || alg == jose.RS512 {
		bits = 4096
	}

	_, priv, err := josex.NewSigningKey(alg, bits)
	if err != nil {
		return nil, errors.Wrapf(ErrUnsupportedKeyAlgorithm, "%s", err)
	}

	if len(kid) == 0 {
		kid = uuid.Must(uuid.NewV4()).String()
	}

	if len(use) == 0 {
		use = "sig"
	}

	return &jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{
				Algorithm:                   string(alg),
				Key:                         priv,
				Use:                         use,
				KeyID:                       kid,
				Certificates:                []*x509.Certificate{},
				CertificateThumbprintSHA256: []byte{},
				CertificateThumbprintSHA1:   []byte{},
			},
		},
	}, nil
}
