package server

import (
	"context"
	"fmt"
	"os"
	"team-manager/internal/biz"
	"team-manager/internal/service"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotServer struct {
	log        *log.Helper
	bot        *tgbotapi.BotAPI
	botService *service.BotService
	userBiz    *biz.UsersUsecase
}

func NewBotServer(
	logger log.Logger,
	botService *service.BotService,
	userBiz *biz.UsersUsecase,
) *BotServer {
	bs := &BotServer{
		log:        log.NewHelper(log.With(logger, "module", "server/bot")),
		botService: botService,
		userBiz:    userBiz,
	}

	tgBot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Errorf("failed on create bot: %v", err)
	} else {
		bs.bot = tgBot
	}

	return bs
}

func (bs *BotServer) Start(ctx context.Context) error {
	bs.telegramBot()
	bs.log.Info("bot started")

	return nil
}

func (bs *BotServer) Stop(ctx context.Context) error {
	bs.bot.StopReceivingUpdates()
	bs.log.Info("bot stopped")

	return nil
}

func (bs *BotServer) telegramBot() {

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates := bs.bot.GetUpdatesChan(u)

	requesting := service.Request("")

	for update := range updates {

		// handle request
		reply, request, err := bs.botService.BotRequest(context.Background(), update, &requesting)
		if err != nil {
			bs.log.Errorf("BotRequest error: %v", err)
			bs.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong. Please, try again later."))
		}

		// reset request
		if request != nil {
			requesting = *request
		} else {
			requesting = ""
		}

		// send reply
		if reply != nil {
			bs.bot.Send(reply)
		} else {
			bs.log.Errorf("Reply not found")
			if update.Message != nil {
				bs.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong. Please, try again later."))
			}
		}
	}
}

func (bs *BotServer) Request(reply tgbotapi.Chattable) {
	_, err := bs.bot.Request(reply)
	if err != nil {
		bs.log.Errorf("Request error: %v", err)
	}
}

func (bs *BotServer) Send(reply tgbotapi.Chattable) {
	_, err := bs.bot.Send(reply)
	if err != nil {
		bs.log.Errorf("Send error: %v", err)
	}
}

func (bs *BotServer) SendQuotes() error {
	quote := bs.botService.GetQuote()

	users, err := bs.userBiz.GetUsers(context.Background())
	if err != nil {
		bs.log.Errorf("GetUsers error: %v", err)
		return err
	}

	for _, user := range users {
		reply := tgbotapi.NewMessage(user.TgID, fmt.Sprintf("%s\n\n©%s", quote.Quote, quote.Author))
		_, err := bs.bot.Send(reply)
		if err != nil {
			bs.log.Errorf("Send error: %v", err)
			return err
		}
	}

	fmt.Println("Quotes sent: ", time.Now())

	return nil
}
