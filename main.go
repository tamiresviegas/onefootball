package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Tarefa struct {
	ID    int
	Texto string
}

type Pessoa struct {
	Name      string `json:"name"`
	Sobrenome string `json:"sobrenome"`
}

func main() {

	//goroutine
	canal := make(chan string)
	var wg sync.WaitGroup

	// Inicia duas goroutines concorrentes
	for i := 1; i <= 2; i++ {
		wg.Add(1) // Incrementa o contador do WaitGroup
		go tarefa(i, &wg, canal)
	}

	// Função anônima para fechar o canal quando todas as goroutines terminarem
	go func() {
		wg.Wait()    // Aguarda até que todas as goroutines tenham terminado
		close(canal) // Fecha o canal quando todas as goroutines terminam
	}()

	// Loop para receber mensagens do canal
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	//Json
	//decode
	pessoa := Pessoa{Name: "Tamires", Sobrenome: "Viegas"}

	jsonData, err := json.Marshal(pessoa)
	if err != nil {
		return
	}

	fmt.Println("Json codificado", string(jsonData))

	//ummarshal
	var pessoaUn Pessoa
	pessoa2 := []byte(`{"name": "Tamires", "sobrenome": "Viegas"}`)

	err = json.Unmarshal(pessoa2, &pessoaUn)
	if err != nil {
		return
	}

	fmt.Printf("Nome: %s, Sobrenome: %s\n", pessoaUn.Name, pessoaUn.Sobrenome)

	//Soma de 2 arrays
	array1 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	array2 := [9]int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	if len(array1) > 4 && len(array2) > 8 {
		resultado := array1[4] + array2[8]
		fmt.Println("Resultado da soma:", resultado)
	} else {
		fmt.Println("Posições desejadas não existem nos arrays.")
	}

	//Encontre o elemento no array
	numeros := [5]int{1, 2, 3, 4, 5}
	elementoAlvo := 5

	result := encontrarElemento(numeros[:], elementoAlvo)

	if result {
		fmt.Printf("O elemento %d está presente no array.\n", elementoAlvo)
	} else {
		fmt.Printf("O elemento %d não está presente no array.\n", elementoAlvo)
	}

	//Contar palava
	palavra := "Teste para Onefootball"
	resultado := contarPalavra(palavra)
	desejada := "Onefootball"
	//Encontrar a palavra
	acheipalavra := encontratPalavra(palavra, desejada)
	if acheipalavra {
		fmt.Printf("A palavra  %s está presente no texto.\n", desejada)
	} else {
		fmt.Printf("A palavra  %s não está presente no texto.\n", desejada)
	}

	fmt.Println("Quantidade de vezes das palavras", resultado)

	//create my slice with the type
	var listaTarefa []Tarefa

	listaTarefa = listaTarefas(listaTarefa, "Teste")
	listaTarefa = listaTarefas(listaTarefa, "Teste 2")

	fmt.Println(listaTarefa)

	for _, tarefa := range listaTarefa {
		fmt.Printf("ID: %d, Tarefa: %s\n", tarefa.ID, tarefa.Texto)
	}

	//create service
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

}

func tarefa(id int, wg *sync.WaitGroup, canal chan string) {
	defer wg.Done() // Decrementa o contador do WaitGroup quando a goroutine termina

	// Simula algum trabalho
	time.Sleep(time.Second)

	// Envia uma mensagem para o canal indicando a conclusão da tarefa
	canal <- fmt.Sprintf("Tarefa %d concluída", id)
}

func encontratPalavra(palava string, desejada string) bool {

	palavras := strings.Fields(palava)

	for _, word := range palavras {
		if word == desejada {
			return true
		}
	}

	return false
}

func encontrarElemento(els []int, alvo int) bool {

	for _, elemento := range els {
		if elemento == alvo {
			return true
		}

	}
	return false
}

func contarPalavra(palavraInput string) map[string]int {

	contagemPalavra := make(map[string]int)
	palavras := strings.Fields(palavraInput)
	for _, palavra := range palavras {
		contagemPalavra[palavra]++
	}

	return contagemPalavra
}

func listaTarefas(tarefa []Tarefa, texto string) []Tarefa {

	novaTarefa := Tarefa{ID: len(tarefa) + 1, Texto: texto}
	return append(tarefa, novaTarefa)
}

/*
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hello World!")
	return
}*/
