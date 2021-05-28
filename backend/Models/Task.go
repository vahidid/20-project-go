package Models

import (
	"vahidid/20-project-go/Config"
)

func GetAllTasks(t *[]Task) (err error) {
	if err = Config.DB.Find(t).Error; err != nil {
		return err
	}

	return nil
}

func SaveTask(t *Task) (err error) {
	if err = Config.DB.Create(t).Error; err != nil {
		return err
	}

	return nil
}

func GetOneTask(b *Task, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneTask(b *Task, id string) (err error) {

	Config.DB.Save(b)
	return nil
}

func DeleteTask(b *Task, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
