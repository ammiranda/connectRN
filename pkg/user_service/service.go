package user_service

import (
	"fmt"
	"github.com/ammiranda/connectRN/pkg/rest_api/models/request"
	"github.com/ammiranda/connectRN/pkg/rest_api/models/response"
	"time"
)

const (
	dateTimeLayout = "2006-01-02"
)

type UserService interface {
	ParseUsers([]request.User) ([]response.User, error)
}

type service struct{}

func NewService() UserService {
	return &service{}
}

func (s *service) ParseUsers(users []request.User) ([]response.User, error) {
	resp := []response.User{}
	for _, user := range users {
		u, err := s.parseUser(user)
		if err != nil {
			return []response.User{}, err
		}
		resp = append(resp, u)
	}

	return resp, nil
}

func (s *service) parseUser(user request.User) (response.User, error) {
	var userResp response.User
	userResp.UserID = user.UserID
	userResp.Name = user.Name

	dobDayOfWeek, err := determineDOBDayOfWeek(user.DOB)
	if err != nil {
		return response.User{}, err
	}

	userResp.DOBDayOfWeek = dobDayOfWeek

	createdOn, err := parseCreatedOn(user.CreatedOn)
	if err != nil {
		return response.User{}, err
	}

	userResp.CreatedOn = createdOn

	return userResp, nil
}

func determineDOBDayOfWeek(t string) (string, error) {
	timeObj, err := time.Parse(dateTimeLayout, t)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", timeObj.Weekday()), nil
}

func parseCreatedOn(t int64) (string, error) {
	loc, err := time.LoadLocation("EST")
	if err != nil {
		return "", err
	}
	timeObj := time.Unix(t, 0)
	return timeObj.In(loc).Format(time.RFC3339), nil
}
