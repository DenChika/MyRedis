package workers

import (
	"github.com/codecrafters-io/redis-starter-go/app/lib/commands"
	"sync"
	"time"
)

const vocabularyClearingDuration int = 5000

type VocabularyClearWorker struct {
	vocabulary map[string]commands.ExpiryValue
	mu         *sync.RWMutex
}

func NewVocabularyClearWorker() *VocabularyClearWorker {
	return &VocabularyClearWorker{vocabulary: commands.Storage.Vocabulary, mu: commands.Storage.Mu}
}

func (w *VocabularyClearWorker) Execute() {
	ticker := time.NewTicker(time.Duration(vocabularyClearingDuration) * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			w.mu.Lock()
			w.clearVocabulary()
			w.mu.Unlock()
			continue
		default:
			continue
		}
	}
}

func (w *VocabularyClearWorker) clearVocabulary() {
	for key, value := range w.vocabulary {
		if time.Now().After(value.ValidUntil) {
			delete(w.vocabulary, key)
		}
	}
}
