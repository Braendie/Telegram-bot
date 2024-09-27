package telegram

const msgHelpEn = `I can save and keep your pages. Also I can offer you them to read.
In order to save the page, just send me all link to it.
In order to get a random page from your list, send me command /rnd.
Caution! After that, this page will be removed from your list!`

const msgHelpRu = `Я могу сохранять и хранить ваши страницы. Также я могу предложить вам их для прочтения.
Чтобы сохранить страницу, просто отправьте мне ссылку на неё.
Чтобы получить случайную страницу из вашего списка, отправьте мне команду /rnd.
Внимание! После этого данная страница будет удалена из вашего списка!`

const msgHelpCmdEn = `\n/rnd - sends a random link from the saved ones.`

const msgHelpCmdRu = `\n/rnd - отправляет случайную ссылки из сохраненных.`

const msgHelloEn = "Hi there! 👋😃 \n\n" + msgHelpEn + "\nYou can use /help_en on english or /help_ru on russian for help"

const (
	msgUnknownCommand = "Unknown command 🤨"
	msgNoSavedPages   = "You have no saved pages 😎"
	msgSaved          = "Saved! 🫡"
	msgAlreadyExists  = "You have already have this page in your list 🤫"
)
