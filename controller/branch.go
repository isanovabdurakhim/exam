package controller

import (
	"app/models"
)

func (c *Controller) CreateBranch(req *models.CreateBranch) (string, error) {
	id, err := c.store.Branch().Create(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) DeleteBranch(req *models.BranchPrimaryKey) error {
	err := c.store.Branch().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateBranch(req *models.UpdateBranch, branchId string) error {
	err := c.store.Branch().Update(req, branchId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetByIdBranch(req *models.BranchPrimaryKey) (models.Branch, error) {
	branch, err := c.store.Branch().GetByID(req)
	if err != nil {
		return models.Branch{}, err
	}
	return branch, nil
}

func (c *Controller) GetAllBranch(req *models.GetBranchListRequest) (models.GetBranchListResponse, error) {
	braches, err := c.store.Branch().GetAll(req)
	if err != nil {
		return models.GetBranchListResponse{}, err
	}
	return braches, nil
}
