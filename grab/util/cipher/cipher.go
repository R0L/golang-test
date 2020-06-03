package cipher

type Cipher interface {
	Encrypt(src string) (string, error)
	Decrypt(src string) (string, error)
}
