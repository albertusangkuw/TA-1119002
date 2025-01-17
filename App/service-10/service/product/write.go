package product

import (
	"errors"
	"service-10/model"
	"service-10/utils"
	"time"
)

func IsExist(tag model.ProductTag) bool {
	result := utils.Database().First(&tag)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func Write(id uint, color uint, name string, uid model.ResUser, targetUpdate []string) error {
	tag := model.ProductTag{
		ID:           id,
		Name:         name,
		Color:        color,
		WriteResUser: uid,
		WriteAt:      time.Now(),
	}

	result := utils.Database().First(&tag)
	if result.RowsAffected == 0 {
		return errors.New("Not Found")
	}

	tx := utils.Database().Begin()
	CreateUserIfNotFound(tx, uid)
	tag.WriteUID = &uid.UID

	targetUpdate = append(targetUpdate, "write_uid", "write_at")
	res := tx.Model(&tag).Select(targetUpdate).Updates(tag)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	tx.Commit()
	return nil
}
