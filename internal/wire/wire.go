package wire

import (
	"dxj/internal/repository"
	"dxj/internal/repository/impl"

	"github.com/google/wire"
)

// ProviderSet is repository providers.
func Initialize() *repository.CategoryRepo {
	wire.Build(impl.NewCategoryRepoImpl)
	return new(repository.CategoryRepo)
}
