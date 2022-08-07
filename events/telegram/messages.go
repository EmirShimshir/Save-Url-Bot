package telegram

const msgHelp = `I can save your pages. Also I can offer you to read them.

In order to save the page, just send me a link to it.

In order to get a random page from your list, send me command /rnd.
(after that the link will be removed from your list)`

const msgHello = "Hello! 👋\n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command 🤯"
	msgNoSavePages    = "You have no saved pages 😢"
	msgSaved          = "Saved! 👌"
	msgAlreadyExists  = "You have already have this page in your list 😌"
)
