package main

import (
	"fmt"
	"log"
)

//фасад для работы с кошельком

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

//создание фасада
func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Начато создание аккаунта")
	walletFacade := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Аккаунт создан")
	return walletFacade
}

//добавление средств в кошелек
func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Добавляем деньги в кошелек")

	err := w.account.checkAccount(accountID) //проверить аккаунт на корректность
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode) //проверить код безопасности для проведения операции
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)                  //зачислить средства в кошелек
	w.notification.sendWalletCreditNotification()   //отправить уведомление о зачислении средств
	w.ledger.makeEntry(accountID, "credit", amount) //сохранить данные об операции
	return nil
}

//списание средств с кошелька
func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Считываем деньги с кошелька")

	err := w.account.checkAccount(accountID) //проверить аккаунт на корректность
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode) //проверить код безопасности для проведения операции
	if err != nil {
		return err
	}

	err = w.wallet.debitBalance(amount) //списать средства
	if err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()   //отправить уведомление о списании средств
	w.ledger.makeEntry(accountID, "debit", amount) //сохранить данные об операции
	return nil
}

type account struct {
	name string
}

//создание аккаунта
func newAccount(accountName string) *account {
	return &account{name: accountName}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Некорректное имя аккаунта")
	}
	fmt.Println("Аккаунт проверен")
	return nil
}

type securityCode struct {
	code int
}

//создание кода безопасности
func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Код безопасности некорректен")
	}
	fmt.Println("Код безопасности проверен")
	return nil
}

type wallet struct {
	balance int
}

//создание кошелька
func newWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Средства успешно начислены на кошелек")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("На счету недостаточно средств")
	}
	fmt.Println("На счету достаточно средств для операции")
	w.balance = w.balance - amount
	return nil
}

//список совершенных операций с кошельком
type ledger struct {
}

func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Сохранены данные об операции для accountId %s с типом операции %s на сумму %d\n", accountID, txnType, amount)
}

//уведомления об операциях
type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Отправка уведомления о пополнении баланса кошелька")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Отправка уведомления о списании средств с кошелька")
}

func main() {
	walletFacade := newWalletFacade("aaa", 1111)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("aaa", 1111, 3)
	if err != nil {
		log.Fatalf("Ошибка: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("aaa", 1111, 1)
	if err != nil {
		log.Fatalf("Ошибка: %s\n", err.Error())
	}
	err = walletFacade.addMoneyToWallet("aaa", 123, 5)
	if err != nil {
		log.Fatalf("Ошибка: %s\n", err.Error())
	}
}
