package lib

type PlaybookRunner interface {
	Run(name string, args []string) error
}
