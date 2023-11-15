package main

import (
	"fmt"
	"os"
	"reflect"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	_ "github.com/lib/pq"
)

type Vacancy struct {
	Company string
	Salary  string
	URL     string
	Title   string
}

var Vacancies = []Vacancy{}
var Link = "https://hh.ru/search/vacancy?search_field=name&search_field=company_name&search_field=description&enable_snippets=false&L_save_area=true&experience=between1And3&professional_role=160&schedule=remote&text=DevOps"

// для работы с бд

func telegramBot() {
	fmt.Print("start")
	//Создаем бота
	bot, err := tgbotapi.NewBotAPI("6638858172:AAEJ2N-9DTX8Pz31Iu4E-CX8BzWCFe2rPHk")
	if err != nil {
		panic(err)
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//Проверяем что от пользователья пришло именно текстовое сообщение
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

			switch update.Message.Text {
			case "/start":
				fmt.Print("first step")
				//Отправлем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a hh parser bot.")
				bot.Send(msg)

			case "/set_link":

				ans := fmt.Sprintf("Вставьте ссылку, по которой будет идти поиск вакансий")

				//Отправлем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, ans)
				bot.Send(msg)
				fmt.Println(update.Message.Text)
				fmt.Println()

			case "/parse":

				Parse("hh.ru", Link)

				//Отправляем данные в БД
				for _, value := range Vacancies {
					if err := CollectData(value.Title, value.Salary, value.Company, value.URL); err != nil {

						//Отправлем сообщение
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database error, but bot still working.")
						bot.Send(msg)
					}
				}
			}

			//Проходим через срез и отправляем каждый элемент пользователю
			for _, value := range Vacancies {

				//Отправлем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, value.Title+value.Company+value.Salary+value.URL)
				bot.Send(msg)
			}
		}

	}
}

func main() {
	time.Sleep(1 * time.Minute)
	fmt.Println(os.Getenv("CREATE_TABLE"))
	fmt.Println(os.Getenv("HOST"))
	fmt.Println(os.Getenv("test"))
	if os.Getenv("CREATE_TABLE") == "yes" {

		if os.Getenv("DB_SWITCH") == "on" {

			if err := CreateTable(); err != nil {
				fmt.Println("не удалось создать таблицу в бд")
				panic(err)
			}
		}
	}
	//time.Sleep(1 * time.Minute)

	//Вызываем бота
	telegramBot()
	//parse("hh.ru", Link)

}
