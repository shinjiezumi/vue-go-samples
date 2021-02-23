package auth

import (
	"log"
	"net/http"

	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
)

type RegisterRequest struct {
	Name string `json:"name" validate:"required,lte=255"`
	LoginRequest
}
type registerUseCase struct{}

func NewRegisterUseCase() *registerUseCase {
	return &registerUseCase{}
}

func (s *registerUseCase) Execute(c *gin.Context, r RegisterRequest) {
	v := validator.New()
	if err := v.Struct(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": common.ExtractValidationErrorMsg(err)})
		return
	}

	repo := user.NewRepository(database.Conn)
	u := repo.GetUserByEmail(r.Email)
	if u != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": common.EmailAlreadyExists,
		})
		return
	}

	u = &user.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: common.HashPassword(r.Password),
	}
	user.NewRepository(database.Conn).Create(u)

	// JWTトークン発行
	authMiddleware, err := CreateAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	authMiddleware.LoginHandler(c)
}
