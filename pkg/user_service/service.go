package user_service

const (
	dateTimeLayout = "2006-01-02"
)

type UserService interface {
	ParseUsers([]request.User) ([]response.User, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ParseUsers([]request.User) ([]response.User, error) {

}
