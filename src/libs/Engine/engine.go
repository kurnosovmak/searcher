package Engine

import (
	"fmt"
	"kurnosovmak/search/src/libs/Engine/Storages"
	"strconv"
	"strings"
)

type Storage interface {
	Get(text string) *[]int
	Set(id int, text string)
}

type Engine struct {
	storage Storage
}

type HandleParams struct {
	Word string
}

type HandleResult struct {
	Items *[]int
}

func NewEngine() *Engine {
	return &Engine{
		storage: Storages.NewMemoryStorage(100),
	}
}

func (engine *Engine) Start() {
	defaultStrings := []string{"test word", "abc der", "test abs", "test word abc der test abs"}
	for index, text := range defaultStrings {
		fmt.Println("Add text - <" + text + "> to id - " + strconv.Itoa(index+1))
		for _, word := range strings.Split(text, " ") {
			engine.storage.Set(index+1, word)
		}
	}
}

func (engine *Engine) Handle(params HandleParams) HandleResult {

	return HandleResult{
		Items: engine.storage.Get(params.Word),
	}
}
