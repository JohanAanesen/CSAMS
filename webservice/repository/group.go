package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
)

// GroupRepository struct
type GroupRepository struct {
	db *sql.DB
}

// NewGroupRepository func
func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

// Fetch single group
func (repo *GroupRepository) Fetch(groupID, assignmentID int) (*model.Group, error) {
	result := model.Group{}

	query := "SELECT id, name FROM groups WHERE group_id = ? AND assignment_id = ?"

	rows, err := repo.db.Query(query, groupID, assignmentID)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			return nil, err
		}

		result.AssignmentID = assignmentID
	}

	return &result, nil
}

// FetchAll groups from an assignment
func (repo *GroupRepository) FetchAll(assignmentID int) ([]*model.Group, error) {
	result := make([]*model.Group, 0)

	query := "SELECT id, name FROM groups WHERE assignment_id = ?"

	rows, err := repo.db.Query(query, assignmentID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := model.Group{}

		err = rows.Scan(&temp.ID, &temp.Name)
		if err != nil {
			return nil, err
		}

		temp.AssignmentID = assignmentID

		result = append(result, &temp)
	}

	return result, nil
}

// FetchUsers gets all users in a group
func (repo *GroupRepository) FetchUsers(groupID int) ([]int, error) {
	result := make([]int, 0)

	query := "SELECT user_id FROM user_groups WHERE group_id = ?"

	rows, err := repo.db.Query(query, groupID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := 0

		err = rows.Scan(&temp)
		if err != nil {
			return nil, err
		}

		result = append(result, temp)
	}

	return result, nil
}

// Insert group into database
func (repo *GroupRepository) Insert(group model.Group) (int64, error) {
	var id int64

	query := "INSERT INTO groups (assignment_id, name, user_id) VALUES (?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return id, err
	}

	rows, err := tx.Exec(query, group.AssignmentID, group.Name, group.Creator)
	if err != nil {
		tx.Rollback()
		return id, err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return id, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return id, err
	}

	return id, nil
}

// Update group in database
func (repo *GroupRepository) Update(group model.Group) error {
	query := "UPDATE groups SET name = ? WHERE id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, group.Name, group.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// AddUser to a group
func (repo *GroupRepository) AddUser(groupID, userID int) error {
	query := "INSERT INTO user_groups (group_id, user_id) VALUES (?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, groupID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// RemoveUser to a group
func (repo *GroupRepository) RemoveUser(groupID, userID int) error {
	query := "DELETE FROM user_groups WHERE group_id = ? AND user_id = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, groupID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// FetchUsersInGroups fetches all users who are in group in a given assigment
func (repo *GroupRepository) FetchUsersInGroups(assignmentID int) ([]int, error) {
	result := make([]int, 0)

	query := "SELECT ug.user_id FROM user_groups AS ug INNER JOIN groups AS g ON ug.group_id = g.id WHERE g.assignment_id = ?"

	rows, err := repo.db.Query(query, assignmentID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp int

		err = rows.Scan(&temp)
		if err != nil {
			return nil, err
		}

		result = append(result, temp)
	}

	return result, nil
}
