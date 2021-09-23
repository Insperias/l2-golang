package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
)

var timeoutArg = flag.String("timeout", "10s", "timeourgt in seconds")

func main() {
	flag.Parse()

	//Если аргументов не 2, то команда используется неправильно
	if flag.NArg() != 2 {
		log.Printf("usage %s host port\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	//получаем аргументы host и port
	host := flag.Arg(0)
	port := flag.Arg(1)

	//устанавливаем время ожидания
	timeoutStr := strings.Replace(*timeoutArg, "s", "", 1)
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Fatalf("timeout argument: %v\n", err)
	}

	//порождаем контекст с временем ожидания
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancelTimeout()

	//устанавливаем дозвонщика, который пытается подключиться к серверу
	var d net.Dialer
	conn, err := d.DialContext(ctxTimeout, "tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("connection: %v\n", err)
	}
	defer conn.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)

	//создаем конктекст для считывания данных
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		defer cancel()

		//считываем данные с коммандной строки и передаем их серверу
		for {
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				_, err = conn.Write(scanner.Bytes())

				if err != nil {
					if errors.Is(err, syscall.EPIPE) {
						log.Println("connection closed")
					} else {
						log.Printf("connection write: %v\n", err)
					}
					return
				}
			} else {
				log.Println("stopped")
				return
			}
		}
	}()

	ch := make(chan string)

	//считываем данные, полученные от сервера
	go func() {
		for {
			scanner := bufio.NewScanner(conn)
			if scanner.Scan() {
				ch <- scanner.Text()
			}
		}
	}()

	go func() {
		defer wg.Done()
		// получаем данные от сервера, пока передаем ему какие-то данные
		for {
			select {
			case text := <-ch:
				fmt.Println(text)
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
}
