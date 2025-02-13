package port

type JWT interface {
	Generate(username, role string) string
}