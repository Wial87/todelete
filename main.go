package main

import (
	"deleter/bot"
)

func main() {
	/*
		Вот пример как вбивать аккаунты:
			accounts := map[string]string{
			"token": "keyWord",
			"token": "keyWord",
		}
	*/

	accounts := map[string]string{
		// "token": "keyWord",
		"60d4f347d1b323a3b11b52e47022c294e329f79809f4cce06d60a1f752ba6e55f106011225e433a952d3e": "дд", // Это пример, поля не должны быть пустыми!!!
	}

	// Функция запуска аккаунтов
	bot.StartAccounts(accounts)
}
