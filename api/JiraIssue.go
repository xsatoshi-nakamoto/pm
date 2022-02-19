package api

import (
	"github.com/xsatoshi-nakamoto/pm/db"
	"github.com/xsatoshi-nakamoto/pm/jiramodels"
	"gorm.io/gorm"
)

func Issue(page int) (tx *gorm.DB) {
	return db.DB.Find(&jiramodels.Project{})
}
