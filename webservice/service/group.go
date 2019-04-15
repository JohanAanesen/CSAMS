package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// GroupService struct service
type GroupService struct {
	groupRepo *repository.GroupRepository
	userRepo  *repository.UserRepository
}

// NewGroupService func
func NewGroupService(db *sql.DB) *GroupService {
	return &GroupService{
		groupRepo: repository.NewGroupRepository(db),
		userRepo:  repository.NewUserRepository(db),
	}
}

func (s *GroupService) fetchUsers(grp *model.Group) error {
	userIDs, err := s.groupRepo.FetchUsers(grp.ID)
	if err != nil {
		return err
	}

	for _, userID := range userIDs {
		user, err := s.userRepo.Fetch(userID)
		if err != nil {
			return err
		}

		grp.Users = append(grp.Users, *user)
	}

	return nil
}

// Fetch group from assignment
func (s *GroupService) Fetch(groupID, assignmentID int) (*model.Group, error) {
	grp, err := s.groupRepo.Fetch(groupID, assignmentID)
	if err != nil {
		return nil, err
	}

	err = s.fetchUsers(grp)
	if err != nil {
		return nil, err
	}

	return grp, nil
}

// FetchAll groups from an assignment
func (s *GroupService) FetchAll(assignmentID int) ([]*model.Group, error) {
	groups, err := s.groupRepo.FetchAll(assignmentID)
	if err != nil {
		return nil, err
	}

	for _, grp := range groups {
		err = s.fetchUsers(grp)
		if err != nil {
			return nil, err
		}
	}

	return groups, nil
}

// Insert group to database
func (s *GroupService) Insert(group model.Group) (int64, error) {
	return s.groupRepo.Insert(group)
}

// Update group in database
func (s *GroupService) Update(group model.Group) error {
	return s.groupRepo.Update(group)
}

// AddUser to group
func (s *GroupService) AddUser(groupID, userID int) error {
	return s.groupRepo.AddUser(groupID, userID)
}

// RemoveUser from group
func (s *GroupService) RemoveUser(groupID, userID int) error {
	return s.groupRepo.RemoveUser(groupID, userID)
}
