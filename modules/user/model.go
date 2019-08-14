package user

// user list (GET /v1/user/users)
func (s *Service) userList() ([]UserModel, error) {
    var users []UserModel
    if err := s.db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
