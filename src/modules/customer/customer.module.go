package customer

type CustomerModule struct {
	Repository *CustomerRepositoryImp
	Controller *CustomerControllerImp
}

func NewModule() *CustomerModule {
	repository := NewRepository(Customer{})
	controller := NewCustomerController(repository)

	return &CustomerModule{
		Repository: repository,
		Controller: controller,
	}
}
