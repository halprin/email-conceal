package context

type EnvironmentGateway interface {
	GetEnvironmentValue(key string) string
}
