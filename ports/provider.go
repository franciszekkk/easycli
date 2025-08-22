package ports

type Provider interface {
	GetNewestReleaseName() (string, error)
	PerformUpdate(binaryPath string) error
}
