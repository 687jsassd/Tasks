package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm_learn/models"
)

type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetByID 根据ID查询用户
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, result.Error
	}
	return &user, nil
}

// GetByCredentials 根据邮箱或手机号查询用户
func (r *UserRepository) GetByCredentials(loginID string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ? OR phone = ?", loginID, loginID).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, result.Error
	}
	return &user, nil
}

// Create 创建新用户
func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

// Update 更新用户信息
func (r *UserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

// Delete 删除用户
func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}

// GetAll 获取用户列表（分页）
func (r *UserRepository) GetAll(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	if err := r.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.DB.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CheckPhoneExists 检查手机号是否已存在
func (r *UserRepository) CheckPhoneExists(phone string, excludeID ...uint) bool {
	var count int64
	query := r.DB.Model(&models.User{}).Where("phone = ?", phone)

	if len(excludeID) > 0 && excludeID[0] > 0 {
		query = query.Where("id != ?", excludeID[0])
	}

	query.Count(&count)
	return count > 0
}

// CheckEmailExists 检查邮箱是否已存在
func (r *UserRepository) CheckEmailExists(email string, excludeID ...uint) bool {
	var count int64
	query := r.DB.Model(&models.User{}).Where("email = ?", email)

	if len(excludeID) > 0 && excludeID[0] > 0 {
		query = query.Where("id != ?", excludeID[0])
	}

	query.Count(&count)
	return count > 0
}

// ActivateUser 激活用户
func (r *UserRepository) ActivateUser(id uint) error {
	return r.DB.Model(&models.User{}).Where("id = ?", id).Update("is_active", true).Error
}
