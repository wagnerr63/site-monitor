package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorimentos = 3
const delay = 5

func main() {
	handleHello()

	for {
		handleMenu()

		comando := handleComand()

		switch comando {

		case 1:
			handleWatch()
		case 2:
			fmt.Println("Exibindo logs...")
			showLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando!")
			os.Exit(-1)

		}
	}
}

func handleHello() {
	nome := "Wagner"
	versao := 1.2
	fmt.Println("Olá sr.", nome, "!")
	fmt.Println("Programa na versão:", versao)
}

func handleMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")
}

func handleComand() int {
	var readComand int

	fmt.Scan(&readComand)
	fmt.Println("Opção: ", readComand)
	println("")

	return readComand
}

func handleWatch() {
	fmt.Println("Monitorando...")
	sites := readSiteFile()

	for i := 0; i < monitorimentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ": ", site)
			handleTests(site)

		}
		println("")
		time.Sleep(delay * time.Second)
	}
	println("")
}

func handleTests(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Houve um erro. Mensagem:(", err, ")")
		os.Exit(-1)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLogs(site, true)
	} else {
		fmt.Println("Site:", site, "teve algum erro em seu carregamento. Status Code:", resp.StatusCode)
		registerLogs(site, false)
	}
}

func readSiteFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Houve um erro ao tentar abrir o arquivo. Mensagem:(", err, ")")
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)

	for {
		linha, err := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLogs(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
