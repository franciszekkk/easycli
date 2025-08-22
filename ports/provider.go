package ports

type Provider interface {
	CheckForUpdates() (bool, error)
	PerformUpdate(path string) error
}
