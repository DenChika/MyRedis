package workers

type BackgroundWorker interface {
	Execute()
}

func StartWorkers() {
	vocabularyClearWorker := NewVocabularyClearWorker()
	go vocabularyClearWorker.Execute()
}
