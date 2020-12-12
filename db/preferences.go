package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type (
	Option struct {
		ID          uint   `json:"id"`
		DisplayText string `json:"Display_text"`
	}
	UserPreferencesData struct {
		ID          uint     `json:"-"`
		DisplayName string   `json:"Display_name"`
		DisplayText string   `json:"Display_text,omitempty"`
		Options     []Option `json:"Options"`
	}
)

func (r *ReaderDB) GetUserPreferences(userID interface{}) (map[uint]*UserPreferencesData, error) {
	rows, err := r.gormDB.Table("user_preferences").Where("user_preferences.user_id = ?", userID).
		Joins("INNER Join preference_category_options on user_preferences.pref_category_option_id = preference_category_options.id").
		Joins("INNER Join preference_categories on preference_category_options.pref_category_id = preference_categories.id").
		Select("preference_categories.id,preference_categories.category_name,preference_categories.category_help_text," +
			"preference_category_options.id,preference_category_options.option_name").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make(map[uint]*UserPreferencesData)

	for rows.Next() {
		pref := &UserPreferencesData{}
		op := Option{}

		err = rows.Scan(&pref.ID, &pref.DisplayName, &pref.DisplayText, &op.ID, &op.DisplayText)
		if err != nil {
			return nil, err
		}

		p, ok := res[pref.ID]
		if !ok {
			pref.Options = append(pref.Options, op)
			res[pref.ID] = pref
			continue
		}
		p.Options = append(p.Options, op)
	}

	return res, nil
}

func (w *WriterDB) CreateUserPrefrences(userID interface{}, optIds []uint) error {

	stmt, valueArgs := buildUserPrefValuesQuery(userID, optIds)
	if err := w.gormDB.Exec(stmt, valueArgs...).Error; err != nil {
		return err
	}

	return nil
}

func (w *WriterDB) UpdateUserPrefrences(ctx context.Context, userID interface{}, optIds []uint) (err error) {

	err = transaction(ctx, w.gormDB, func(tx *gorm.DB) error {
		err := tx.Where("user_id = ?", userID).Delete(UserPreferences{}).Error
		if err != nil {
			return err
		}

		stmt, valueArgs := buildUserPrefValuesQuery(userID, optIds)
		if err := tx.Exec(stmt, valueArgs...).Error; err != nil {
			return err
		}
		return nil
	})
	return
}

func buildUserPrefValuesQuery(userID interface{}, optIds []uint) (stmt string, valueArgs []interface{}) {
	valueStrings := []string{}
	for _, optID := range optIds {
		valueStrings = append(valueStrings, "(?, ?)")

		valueArgs = append(valueArgs, userID)
		valueArgs = append(valueArgs, optID)
	}

	stmt = `INSERT INTO user_preferences(user_id, pref_category_option_id) VALUES %s`
	stmt = fmt.Sprintf(stmt, strings.Join(valueStrings, ","))

	return
}
