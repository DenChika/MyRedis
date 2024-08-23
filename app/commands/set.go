package commands

import (
	"errors"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp/types"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Set struct {
	Vocabulary map[string]string
	Mu         *sync.Mutex
}

func (c *Set) Execute(input []string) (string, error) {
	src, value := input[0], input[1]

	c.Mu.Lock()
	c.Vocabulary[src] = value
	c.Mu.Unlock()

	if len(input) > 2 {
		normalized := strings.ToLower(input[2])
		if err := ensureThatParamIsNotMissing(input[2:]); err != nil {
			return "", err
		}
		switch normalized {
		case "px":
			v, err := ensureThatIntParamIsValid(input[3])
			if err != nil {
				return "", err
			}
			go c.pxExecute(input[2], v)
		default:
			return "", errors.New("Unsupported command: " + src)
		}
	}

	var respType resp.Type = &types.SimpleString{Content: "OK"}

	return respType.Encode(), nil
}

func (c *Set) pxExecute(word string, ms int) {
	ticker := time.NewTicker(time.Duration(ms) * time.Millisecond)
	fmt.Println("tick")
	for {
		select {
		case <-ticker.C:
			fmt.Println("boom")
			delete(c.Vocabulary, word)
			break
		default:
			continue
		}
	}

}

func ensureThatParamIsNotMissing(params []string) error {
	if len(params) == 1 {
		return errors.New(fmt.Sprintf("Param value of %s flag is missing", params[0]))
	}

	return nil
}

func ensureThatIntParamIsValid(param string) (int, error) {
	v, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}

	if v <= 0 {
		return 0, errors.New("expiry parameter must be greater than zero")
	}

	return v, nil
}
