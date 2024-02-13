package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"net/http"
	v1 "team-manager/api/helloworld/v1"
	"team-manager/internal/biz"
)

type BotService struct {
	log *log.Helper

	v1.UnimplementedGreeterServer
	userBiz *biz.UsersUsecase
}

// NewBotService new a greeter service.
func NewBotService(
	userBiz *biz.UsersUsecase,
	logger log.Logger,
) *BotService {
	return &BotService{
		userBiz: userBiz,
		log:     log.NewHelper(log.With(logger, "module", "service/bot")),
	}
}

func (bs *BotService) BotRequest(ctx context.Context, update tgbotapi.Update, request *Request) (tgbotapi.Chattable, *Request, error) {
	message := update.Message
	reply := tgbotapi.Chattable(nil)

	if request != nil && *request != "" {
		switch *request {
		case CLUB_NAME:
			if message.Text == "" {
				reply = tgbotapi.NewMessage(message.Chat.ID, "Please, enter a name for the team")
				return reply, request, nil
			}
			//team, err := bs.clubBiz.CreateClub(ctx, tgUser, message.Text)
			//if err != nil {
			//	bs.log.Errorf("CreateClub error: %v", err)
			//	return nil, nil, err
			//}

			//reply = tgbotapi.NewMessage(chatId, GetContent(team, false))
			return reply, nil, nil
		default:
			bs.log.Errorf("Request not found: %v", *request)
			return nil, nil, nil
		}
	}

	// Check if we've gotten a message update.
	if message != nil {
		chatId := update.Message.Chat.ID
		tgUser := update.Message.From

		if message.Command() != "" {
			switch message.Command() {
			case "start":
				bs.RegisterUser(ctx, tgUser)
				reply = tgbotapi.NewMessage(chatId, "Welcome to the team manager bot! Create, Manage and Win!")

			case "quote":
				quote := bs.GetQuote()
				reply = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n\n©%s", quote.Quote, quote.Author))

			case "help":
				reply = tgbotapi.NewMessage(chatId, "I did a list to you here for commands that I can do:\n"+
					"/start - Start the bot\n"+
					"/quote - Get inspired\n"+
					"/help - Show this message\n")

			default:
				reply = tgbotapi.NewMessage(chatId, "Unknown command, use one of the commands in the menu.")
			}

			return reply, request, nil
		} else if message.Text != "" {
			reply = tgbotapi.NewMessage(chatId, "Welcome to the Almaty Metro bot! Please, use the commands in the menu.")
			return reply, request, nil
		}
	} else if update.CallbackQuery != nil {
		//// Respond to the callback query, telling Telegram to show the user
		//// a message with the data received.
		//callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		//if _, err := bs.bot.Request(callback); err != nil {
		//	panic(err)
		//}
		//
		//// And finally, send a message containing the data received.
		//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
		//if _, err := bs.bot.Send(msg); err != nil {
		//	panic(err)
		//}
	}

	return reply, request, nil
}

func (bs *BotService) RegisterUser(ctx context.Context, tgUser *tgbotapi.User) {
	_, err := bs.userBiz.CreateUser(ctx, tgUser)
	if err != nil {
		bs.log.Errorf("CreateUser error: %v", err)
		return
	}
}

type Quote struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

func (bs *BotService) GetQuote() *Quote {
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	quote := []Quote{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil
	}

	return &quote[0]
}
