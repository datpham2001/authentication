package utils

const (
	ACCOUNT = "ACCOUNT"

	ACCOUNT_LENGTH = 6

	STRING_TO_GEN_ID = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenNanoID(alphabet string, length int) string {
	res, err := gonanoid.Generate(alphabet, length)
}
