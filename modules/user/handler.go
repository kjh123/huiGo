package user

import (
	"fmt"
	"github.com/kjh123/huiGo/log"
	"net/http"
)

func (s *Service) getUsers(w http.ResponseWriter, r *http.Request) {
	// TODO 返回用户列表
	log.INFO.Println("用户列表")
	fmt.Errorf("create users")
}

func (s *Service) createUser(w http.ResponseWriter, r *http.Request) {
	// TODO 新增用户
	fmt.Errorf("create users")
}

func (s *Service) updateUser(w http.ResponseWriter, r *http.Request) {
	// TODO 更新用户信息

}
