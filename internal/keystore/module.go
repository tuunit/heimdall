package keystore

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/rs/zerolog"
	"go.uber.org/fx"

	"github.com/dadrus/heimdall/internal/config"
	"github.com/dadrus/heimdall/internal/heimdall"
	"github.com/dadrus/heimdall/internal/x/errorchain"
)

// nolint
var Module = fx.Options(
	fx.Provide(newKeyStore),
)

func newKeyStore(conf config.Configuration, logger zerolog.Logger) (KeyStore, error) {
	const rsa2048 = 2048

	if conf.Signer == nil {
		logger.Warn().Msg("Signer is not configured. NEVER DO IT IN PRODUCTION!!!! Generating an RSA key pair.")

		privateKey, err := rsa.GenerateKey(rand.Reader, rsa2048)
		if err != nil {
			return nil, errorchain.
				NewWithMessage(heimdall.ErrInternal, "failed to generate RSA-2048 key pair").
				CausedBy(err)
		}

		return NewKeyStoreFromKey(privateKey)
	}

	return NewKeyStoreFromPEMFile(conf.Signer.File, conf.Signer.Password)
}
