package telegram

const msgHelpEn = `I can save and keep your pages. Also I can offer you them to read.
In order to save the page, just send me all link to it.
You can also tag the link to make it easier to find in the future.
Additionally, you can add a description to your link by writing it after the tag or using #desc:.`

const msgHelpRu = `Я могу сохранять и хранить ваши страницы. Также я могу предложить вам их для прочтения.
Чтобы сохранить страницу, просто отправьте мне ссылку на неё.
Вы также можете пометить ссылку тегом, чтобы в дальнейшем было удобней ее получать.
Также вы можете написать описание для вашей ссылки написав ее после тега или с помощью "#desc:".`

const msgHelpCmdEn = `In all examples, insert your own data without brackets.

Send the link like this: [Your link] [Your tag] (optional) [Your description] (optional)
Or like this: [Your link] #desc: [Your description]

/rnd - sends a random link from the saved ones.

/tag Sends all links contained in the specified tag. 
Send it like this: /tag [Your tag]

/rndtag sends a random link from the saved ones related to the specified tag. 
Send it like this: /rndtag [Your tag]`

const msgHelpCmdRu = `Во всех примерах вставлять свои данные без скобок.

Присылать ссылку вот так: [Ваша ссылка] [Ваш тег](не обязательно) [Ваше описание](не обязательно)
Либо вот так: [Ваша ссылка] #desc: [Ваше описание]

/rnd - отправляет случайную ссылки из сохраненных.

/tag присылает все ссылки лежащие в данном теге. 
Присылать вот так: /tag [Ваш тег]

/rndtag присылает случайную ссылку из сохраненных, относящуюся к данному тегу. 
Присылать вот так: /rndtag [Ваш тег]`

const msgHelloEn = "Hi there! 👋😃 \n\n" + msgHelpEn + "\nYou can use /help_en on english or /help_ru on russian for help"

const (
	msgUnknownCommand = "Unknown command 🤨"
	msgNoSavedPages   = "You have no saved pages 😎"
	msgSaved          = "Saved! 🫡"
	msgAlreadyExists  = "You have already have this page in your list 🤫"
	msgTagIsEmpty     = "This tag is empty 😅"
	msgWrongTagCmd    = "You need to send it like this: /tag [Your tag] 🤓"
	msgWrongRndTagCmd = "You need to send it like this: /rndtag [Your tag] 🤓"
)
