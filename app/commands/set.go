package commands

import (
	"errors"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/lib/commands"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/resp/types"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Set struct {
	Vocabulary map[string]commands.ExpiryValue
	Mu         *sync.RWMutex
}

func (c *Set) Execute(input []string) (string, error) {
	src, value := input[0], input[1]

	c.Mu.Lock()
	c.Vocabulary[src] = commands.ExpiryValue{
		Value:     value,
		IsDurable: true,
	}
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

			c.Mu.Lock()
			c.pxExecute(src, value, v)
			c.Mu.Unlock()
		default:
			return "", errors.New("Unsupported command: " + src)
		}
	}

	var respType resp.Type = &types.SimpleString{Content: "OK"}

	return respType.Encode(), nil
}

func (c *Set) pxExecute(word string, value string, duration int) {
	c.Vocabulary[word] = commands.ExpiryValue{
		Value:      value,
		ValidUntil: time.Now().Add(time.Duration(duration) * time.Millisecond),
		IsDurable:  false,
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
