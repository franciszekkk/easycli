package ports

type Provider interface {
	GetNewestReleaseName() (string, error)
	PerformUpdate(path string) error
}
