package bot

type Engine interface {
	Start() error
	Stop() error
	Restart() error
}
