// @Author zhangjiaozhu 2023/10/15 15:10:00
package services

import (
	"FlashSkill/backend/models"
	"FlashSkill/backend/respositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	IsPwdOK(string,string)(*models.User,bool)
	AddUser(user *models.User)(int64,error)
	Select(username string)(*models.User,error)
}

type UserService struct {
	UserResp respositories.IUserResp
}

func (u *UserService) Select(username string) (*models.User, error) {
	//TODO implement me
	return u.UserResp.Select(username)
}

func (u *UserService) IsPwdOK(username string, password string) (*models.User, bool) {
	//TODO implement me
	user, err := u.UserResp.Select(username)
	if err != nil {
		return nil,false
	}
	isOk , err := CheckPassword(password,user.HashPassword)
	if !isOk {
		return nil,false
	}
	return user,true
	
}

func (u *UserService) AddUser(user *models.User) (int64, error) {
	//TODO implement me
	password, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return 0, err
	}
	user.HashPassword = string(password)
	return u.UserResp.Insert(user)
}

func NewUserService() IUserService {
	return &UserService{UserResp: respositories.NewUserResp("user")}
}

func CheckPassword(userpassword string,hashed string) (bool,error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(userpassword));err != nil {
		return false, errors.New("密码比对错误")
	}
	return true,nil
}

func GeneratePassword(password string)([]byte,error)  {
	return bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
}