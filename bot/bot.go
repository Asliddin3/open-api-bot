package bot

import (
	"log"

	"github.com/Asliddin3/open-api-bot/api"
	"github.com/Asliddin3/open-api-bot/config"
	"github.com/Asliddin3/open-api-bot/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	cfg     config.Config
	storage storage.StorageI
	bot     *tgbotapi.BotAPI
	api     api.ApiI
}

func New(cfg config.Config, strg storage.StorageI) BotHandler {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return BotHandler{
		cfg:     cfg,
		storage: strg,
		bot:     bot,
		api:     api.RegisterApi(),
	}
}

func (h *BotHandler) Start() {
	log.Printf("Authorized on account %s", h.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.bot.GetUpdatesChan(u)

	for update := range updates {
		go h.HandleBot(update)
	}
}

func (h *BotHandler) HandleBot(update tgbotapi.Update) {
	user, err := h.storage.GetOrCreate(update.Message.From.ID, update.Message.From.FirstName)
	if err != nil {
		log.Println("failed to call storage.GetOrCreate: ", err)
		h.SendMessage(user, "error happened")
	}
	if update.Message.Command() == "start" {
		
	}

	if err != nil {
		log.Println("failed to handle message: ", err)
		h.SendMessage(user, "error happened")
	}
}

func (h *BotHandler) SendMessage(user *storage.User, message string) {
	msg := tgbotapi.NewMessage(user.TgID, message)
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}
}
