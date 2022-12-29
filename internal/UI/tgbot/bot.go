package tgbot

import (
	"errors"
	"log"
	"math"
	"strconv"
	"tg/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	botAPI   *tgbotapi.BotAPI
	timeout  int
	services *service.Service
}

func NewTgBot(api string, debug bool, timeout int, services *service.Service) (*TgBot, error) {
	bot, err := tgbotapi.NewBotAPI(api)
	if err != nil {
		return nil, err
	}
	bot.Debug = debug

	return &TgBot{
		botAPI:   bot,
		timeout:  timeout,
		services: services,
	}, nil
}

func (tg *TgBot) GetUpdates() {

	//	tg.botAPI.Debug = true

	log.Printf("Authorized on account %s", tg.botAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.botAPI.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /create, /delete and /status."
		case "create":
			trigger, err := triggerValidator(&update)
			if err != nil {
				msg.Text = err.Error()
				break
			}

			if err := tg.CreateTrigger(update.Message.Chat.ID, trigger); err != nil {
				msg.Text = err.Error()
				break
			}
			msg.Text = "Уведомление успешно установлено - ждите оповещения!"
		case "status":
			res, err := tg.services.GetTriggers(update.Message.Chat.ID)
			if err != nil {
				log.Println(err)
				msg.Text = "Что-то пошло не так!"
				break
			}
			msg.Text = "Список уведомлений:\n"
			for i := 0; i < len(res); i++ {
				msg.Text += strconv.FormatFloat(res[i], 'f', 2, 64) + "\n"
			}
			tg.services.GetKeys()
		case "delete":
			trigger, err := triggerValidator(&update)
			if err != nil {
				msg.Text = err.Error()
				break
			}

			if err := tg.DeleteTrigger(update.Message.Chat.ID, trigger); err != nil {
				msg.Text = err.Error()
				break
			}
			msg.Text = "Уведомление успешно удалено!"
		default:
			msg.Text = "Я не знаю такой команды"
		}

		if _, err := tg.botAPI.Send(msg); err != nil {
			log.Panic(err)
		}
		// // инициализируем объект планировщика
		// s := gocron.NewScheduler(time.UTC)
		// // добавляем одну задачу на каждую минуту
		// s.Cron("* * * * *").Do(tg.SendNotify())
		// // запускаем планировщик с блокировкой текущего потока
		// s.StartBlocking()
		tg.SendNotify()
	}
}

func triggerValidator(update *tgbotapi.Update) (float64, error) {
	value := update.Message.CommandArguments()

	trigger, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, errors.New("неверное значение")
	}
	return math.Round(trigger*100) / 100, nil
}
