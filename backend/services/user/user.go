package user

import (
	"context"
	"errors"

	"react-go-practice/models"
	userRepo "react-go-practice/repository/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo userRepo.UserRepository
}

// NewUserService ユーザーサービスのインスタンスを作成
func NewUserService(userRepo userRepo.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Register 新規ユーザー登録
func (s *UserService) Register(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	// ユーザー名の重複チェック
	existingUser, _ := s.userRepo.GetByUsername(ctx, req.Username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	createReq := models.CreateUserRequest{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	user, err := s.userRepo.Create(ctx, createReq)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login ユーザーのログイン
func (s *UserService) Login(ctx context.Context, req models.LoginRequest) (*models.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

// GetByID ユーザーのIDでユーザーを取得
func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

// GetByUsername ユーザーのユーザー名でユーザーを取得
func (s *UserService) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.userRepo.GetByUsername(ctx, username)
}

// List ユーザーの一覧を取得
func (s *UserService) List(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.List(ctx)
}

// ChangePassword ユーザーのパスワードを変更
func (s *UserService) ChangePassword(ctx context.Context, id uuid.UUID, currentPassword, newPassword string) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(currentPassword))
	if err != nil {
		return errors.New("current password is incorrect")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.userRepo.UpdatePassword(ctx, id, string(hashedPassword))
	return err
}

// Delete ユーザーを削除
func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.userRepo.Delete(ctx, id)
}
