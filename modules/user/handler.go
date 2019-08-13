package user

import (
	"github.com/kjh123/huiGo/util/response"
	"net/http"
)

func (s *Service) getUsers(w http.ResponseWriter, r *http.Request) {
	// TODO 返回用户列表
	response.WriteJSON(w, map[string]interface{}{
		"data": s.userList,
		"message": "user list",
		"code" : "200",
	}, 200)
}

func (s *Service) createUser(w http.ResponseWriter, r *http.Request) {
	// TODO 新增用户
	response.WriteJSON(w, map[string]interface{}{
		"message": "create user",
		"code": "200",
	}, 200)
}

func (s *Service) updateUser(w http.ResponseWriter, r *http.Request) {
	// TODO 更新用户信息
	response.WriteJSON(w, map[string]interface{}{
		"message": "user update",
		"code": "200",
	}, 200)
}

func (s *Service) deleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO 删除用户信息
	response.WriteJSON(w, map[string]interface{}{
		"message": "drop user",
		"code": "200",
	}, 200)
}
