package ports

import "github.com/immersivesky/go-rest-mongo/internal/core/domain"

type DB interface {
  GetChat(int) *domain.Chat
  CreateChat(int) bool
}
