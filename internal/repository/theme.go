package repository

import "github.com/hipeday/rosen/internal/logging"

// Theme represents the theme entity.
type Theme struct {
	GeneralEntity[int64]
	Name      string `db:"name" json:"name"`           // 主题显示名称
	Defaulted bool   `db:"defaulted" json:"defaulted"` // 当前主题是否为默认主题
	Folder    string `db:"folder" json:"folder"`       // 当前主题在文件系统中的文件夹名称
}

// ThemeRepository represents the theme repository.
type ThemeRepository struct {
	DefaultRepository
}

// NewThemeRepository creates a new theme repository.
func NewThemeRepository() *ThemeRepository {
	logger := logging.Logger()
	repository := ThemeRepository{}
	defaultRepository, err := newRepositoryFactory(&repository)
	if err != nil {
		logger.Fatal(err)
	}
	return defaultRepository.(*ThemeRepository)
}

// SelectDefaultTheme selects the default theme. If no default theme is found, it returns nil.
func (r *ThemeRepository) SelectDefaultTheme(tenant int64) (*Theme, error) {
	theme := &Theme{}
	err := r.db.Get(theme, "SELECT * FROM theme WHERE defaulted = true and tenant_id = ?", tenant)
	if r.checkIsNoRowsErr(err) {
		return nil, nil
	}
	return theme, err
}
