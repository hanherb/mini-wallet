package wallet

type WalletModule struct {
	Repository *WalletRepositoryImp
	Controller *WalletControllerImp
}

func NewModule() *WalletModule {
	repository := NewRepository(Wallet{})
	controller := NewWalletController(repository)

	return &WalletModule{
		Repository: repository,
		Controller: controller,
	}
}
