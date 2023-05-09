package data

type SecretRepo interface {
	GetSecret() (*Secret, error)
	SaveSecret(value Secret) error
	UpdateSecret(value Secret) error
}
