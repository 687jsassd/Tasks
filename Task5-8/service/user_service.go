package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm_learn/models"
	"gorm_learn/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Register 用户注册业务逻辑
func (s *UserService) Register(user *models.User) error {
	// 检查邮箱是否已存在
	if s.Repo.CheckEmailExists(user.Email) {
		return errors.New("邮箱已被其他账号绑定")
	}

	// 检查手机号是否已存在
	if s.Repo.CheckPhoneExists(user.Phone) {
		return errors.New("手机号已被其他账号绑定")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败: " + err.Error())
	}
	user.Password = string(hashedPassword)

	// 初始化账户状态
	if user.Role == 0 {
		user.Role = models.RoleUser
	}
	user.IsActive = false

	// 保存用户
	return s.Repo.Create(user)
}

// Login 用户登录业务逻辑
func (s *UserService) Login(loginID, password string) (*models.User, error) {
	// 查询用户
	user, err := s.Repo.GetByCredentials(loginID)
	if err != nil {
		return nil, err
	}

	// 检查用户是否已激活
	if !user.IsActive {
		return nil, errors.New("用户未激活，请联系管理员")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	return user, nil
}

// GetUserProfile 获取用户个人资料
func (s *UserService) GetUserProfile(userID uint) (*models.User, error) {
	return s.Repo.GetByID(userID)
}

// UpdateProfile 更新用户个人资料
func (s *UserService) UpdateProfile(userID uint, updateData map[string]interface{}) (*models.User, error) {
	// 获取用户信息
	user, err := s.Repo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	// 更新字段
	if name, ok := updateData["name"].(string); ok {
		user.Name = name
	}
	if phone, ok := updateData["phone"].(string); ok {
		// 检查手机号唯一性
		if s.Repo.CheckPhoneExists(phone, userID) {
			return nil, errors.New("手机号已被其他账号绑定")
		}
		user.Phone = phone
	}
	if email, ok := updateData["email"].(string); ok {
		// 检查邮箱唯一性
		if s.Repo.CheckEmailExists(email, userID) {
			return nil, errors.New("邮箱已被其他账号绑定")
		}
		user.Email = email
	}

	// 保存更新
	if err := s.Repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// ActivateUser 激活用户
func (s *UserService) ActivateUser(userID uint) error {
	return s.Repo.ActivateUser(userID)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(userID uint) error {
	// 检查用户是否存在
	user, err := s.Repo.GetByID(userID)
	if err != nil {
		return err
	}

	// 防止删除管理员
	if user.Role == models.RoleAdmin {
		return errors.New("不能删除管理员用户")
	}

	return s.Repo.Delete(userID)
}

// GetAllUsers 获取用户列表（分页）
func (s *UserService) GetAllUsers(page, pageSize int) ([]models.User, int64, error) {
	return s.Repo.GetAll(page, pageSize)
}
