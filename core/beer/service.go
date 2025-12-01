package beer

type UseCase interface {
	//Retorna um array que dentro de cada posicção existe a posição da memoria de cada item do array
	GetAll() ([]*Beer, error)
	GetByID(Id int) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(Id int) error
}

type Service struct{}

func NewService() *Service {
	// & retorna o endereço de memória do objeto Service
	return &Service{}
}

func (s *Service) GetAll() ([]*Beer, error) {
	//  
	return nil, nil
}

func (s *Service) GetByID(Id int) (*Beer, error) {
	return nil, nil
}

func (s *Service) Store(b *Beer) error {
	return nil
}

func (s *Service) Update(b *Beer) error {
	return nil
}

func (s *Service) Remove(Id int) error {
	return nil
}
