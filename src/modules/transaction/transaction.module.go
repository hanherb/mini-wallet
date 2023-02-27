package transaction

type TransactionModule struct {
	Repository *TransactionRepositoryImp
	Controller *TransactionControllerImp
}

func NewModule() *TransactionModule {
	repository := NewRepository(Transaction{})
	controller := NewTransactionController(repository)

	return &TransactionModule{
		Repository: repository,
		Controller: controller,
	}
}
