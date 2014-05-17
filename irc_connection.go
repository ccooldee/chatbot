package bot

type ircConnection interface {
	Privmsg(target, message string)
	Notice(target, message string)
	Join(channel string)
	Part(channel string)
	Quit()
	GetNick() string
}
