package bot

type UseCase interface {
	Start()
	Stop()
	Restart()
	SendMsg()
	SendGroupMsg()
}
