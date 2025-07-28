package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm_learn/models"
	"gorm_learn/repository"
	"gorm_learn/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// 请求模型
// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required,len=11"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// ActivateRequest 用户激活请求
type ActivateRequest struct {
	UserID uint `json:"user_id" binding:"required,min=1"`
}

// 响应模型
// Response 统一API响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 路由处理函数
// registerHandler 用户注册接口
func registerHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 转换请求为User模型
	user := &models.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Email:    req.Email,
		Password: req.Password,
	}

	// 创建仓库和服务实例
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// 调用注册业务逻辑
	if err := userService.Register(user); err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "邮箱已被") || strings.Contains(err.Error(), "手机号已被") {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, Response{
			Code:    statusCode,
			Message: "注册失败: " + err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "注册成功，请等待管理员审核激活",
		Data:    gin.H{"user_id": user.ID},
	})
}

// activateHandler 用户激活接口
func activateHandler(c *gin.Context) {
	var req ActivateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 获取当前用户角色
	roleVal, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    http.StatusUnauthorized,
			Message: "未授权访问",
		})
		return
	}
	role, ok := roleVal.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "用户角色格式错误",
		})
		return
	}

	// 检查是否为管理员
	if role != models.RoleAdmin {
		c.JSON(http.StatusForbidden, Response{
			Code:    http.StatusForbidden,
			Message: "权限不足，仅管理员可执行此操作",
		})
		return
	}

	// 创建仓库和服务实例
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// 调用激活业务逻辑
	if err := userService.ActivateUser(req.UserID); err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "用户不存在") {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, Response{
			Code:    statusCode,
			Message: "激活失败: " + err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "用户激活成功",
	})
}

// LoginRequest 用户登录请求结构
type LoginRequest struct {
	LoginID  string `json:"login_id" binding:"required"` // 邮箱或手机号
	Password string `json:"password" binding:"required"`
}

// Claims JWT声明结构
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   int    `json:"role"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// JWT密钥(测试环境)
// 对于WindowsPowerShell,使用[guid]::NewGuid().ToString().Replace("-","") + [guid]::NewGuid().ToString().Replace("-","")生成
var jwtSecret = []byte("08d42af14d9f4dcabe7f465e6cc915b5ff1f65f94eb146a59b3fafeeec41ff8c")

// GenerateToken 生成JWT令牌
func GenerateToken(userID uint, role int, email string) (string, error) {
	// 设置令牌过期时间为24小时
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "member-management-system",
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名令牌
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    http.StatusUnauthorized,
				Message: "未提供认证令牌",
			})
			c.Abort()
			return
		}

		// 检查令牌格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    http.StatusUnauthorized,
				Message: "认证令牌格式错误",
			})
			c.Abort()
			return
		}

		// 解析令牌
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    http.StatusUnauthorized,
				Message: "无效的认证令牌",
			})
			c.Abort()
			return
		}

		// 将用户信息存储在上下文中
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("email", claims.Email)

		c.Next()
	}
}

// AdminMiddleware 管理员权限中间件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户角色
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    http.StatusUnauthorized,
				Message: "未授权访问",
			})
			c.Abort()
			return
		}

		// 检查是否为管理员
		if role != models.RoleAdmin {
			c.JSON(http.StatusForbidden, Response{
				Code:    http.StatusForbidden,
				Message: "权限不足，仅管理员可执行此操作",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// loginHandler 用户登录接口
func loginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 创建仓库和服务实例
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// 调用登录业务逻辑
	user, err := userService.Login(req.LoginID, req.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		message := "登录失败: " + err.Error()
		if strings.Contains(err.Error(), "用户不存在") {
			statusCode = http.StatusUnauthorized
		} else if strings.Contains(err.Error(), "未激活") {
			statusCode = http.StatusForbidden
		} else if strings.Contains(err.Error(), "密码错误") {
			statusCode = http.StatusUnauthorized
		}
		c.JSON(statusCode, Response{
			Code:    statusCode,
			Message: message,
		})
		return
	}

	// 生成JWT令牌
	token, err := GenerateToken(user.ID, user.Role, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "生成令牌失败: " + err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "登录成功",
		Data: gin.H{
			"token": token,
			"user": gin.H{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
				"role":  user.Role,
			},
		},
	})
}

// getUserHandler 获取单个用户信息接口
func getUserHandler(c *gin.Context) {
	var temp struct {
		ID uint `uri:"id" binding:"required,min=1"`
	}
	if err := c.ShouldBindUri(&temp); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "用户ID参数错误: " + err.Error(),
		})
		return
	}
	userID := temp.ID

	// 创建仓库和服务实例
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// 获取用户信息
	user, err := userService.GetUserProfile(userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "用户不存在") {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, Response{
			Code:    statusCode,
			Message: "获取用户信息失败: " + err.Error(),
		})
		return
	}

	// 隐藏密码字段
	user.Password = ""

	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "获取用户信息成功",
		Data:    user,
	})
}

// getAllUsersHandler 获取所有用户信息接口
func getAllUsersHandler(c *gin.Context) {
	var page, pageSize int
	var err error

	// 获取分页参数
	page, err = strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err = strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 创建仓库和服务实例
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// 获取用户列表
	users, total, err := userService.GetAllUsers(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "获取用户列表失败: " + err.Error(),
		})
		return
	}

	// 隐藏密码字段
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "获取用户列表成功",
		Data: gin.H{
			"users":       users,
			"total":       total,
			"page":        page,
			"page_size":   pageSize,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

// deleteUserHandler 删除用户接口
func deleteUserHandler(c *gin.Context) {
	var temp struct {
		ID uint `uri:"id" binding:"required,min=1"`
	}
	if err := c.ShouldBindUri(&temp); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "用户ID参数错误: " + err.Error(),
		})
		return
	}
	userID := temp.ID

	currentUserID, _ := c.Get("userID")
	if userID == currentUserID.(uint) {
		c.JSON(http.StatusBadRequest, Response{Message: "不能删除自己"})
		return
	}
	// 创建仓库和服务实例
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// 调用删除业务逻辑
	if err := userService.DeleteUser(userID); err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "用户不存在") {
			statusCode = http.StatusBadRequest
		} else if strings.Contains(err.Error(), "不能删除管理员") {
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, Response{
			Code:    statusCode,
			Message: "删除用户失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "删除用户成功",
	})
}

// UpdateProfileRequest 用户更新个人信息请求结构
type UpdateProfileRequest struct {
	Name  *string `json:"name,omitempty" binding:"omitempty,min=2,max=63"`
	Phone *string `json:"phone,omitempty" binding:"omitempty,len=11"`
	Email *string `json:"email,omitempty" binding:"omitempty,email"`
}

// updateProfileHandler 更新个人信息接口
func updateProfileHandler(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    http.StatusUnauthorized,
			Message: "未授权访问",
		})
		return
	}

	// 绑定请求参数
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 创建仓库和服务实例
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// 获取用户信息
	user, err := userService.GetUserProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: "获取用户信息失败: " + err.Error(),
		})
		return
	}

	// 准备更新数据
	updateData := make(map[string]interface{})
	if req.Name != nil && *req.Name != user.Name {
		updateData["name"] = *req.Name
	}
	if req.Phone != nil && *req.Phone != user.Phone {
		updateData["phone"] = *req.Phone
	}
	if req.Email != nil && *req.Email != user.Email {
		updateData["email"] = *req.Email
	}

	// 调用更新业务逻辑
	updatedUser, err := userService.UpdateProfile(userID.(uint), updateData)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "邮箱已被") || strings.Contains(err.Error(), "手机号已被") {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, Response{
			Code:    statusCode,
			Message: "更新失败: " + err.Error(),
		})
		return
	}

	// 隐藏密码字段
	updatedUser.Password = ""

	// 返回更新后的用户信息
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "更新成功",
		Data:    updatedUser,
	})

}

func main() {
	// 初始化Gin引擎
	r := gin.Default()

	// 注册路由
	api := r.Group("/api/v1")
	{
		api.POST("/register", registerHandler)
		api.POST("/login", loginHandler)

		// 需要认证的路由
		auth := api.Group("")
		auth.Use(AuthMiddleware())
		{
			// 普通用户路由
			auth.PUT("/profile", updateProfileHandler)

			// 管理员路由
			admin := auth.Group("")
			admin.Use(AdminMiddleware())
			{
				admin.POST("/activate", activateHandler)
				admin.GET("/users/:id", getUserHandler)
				admin.GET("/users", getAllUsersHandler)
				admin.DELETE("/users/:id", deleteUserHandler)
			}
		}
	}

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		panic("启动服务器失败: " + err.Error())
	}
}
