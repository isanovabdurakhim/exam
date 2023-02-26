package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type branchRepo struct {
	fileName string
}

func NewBranchRepo(fileName string) *branchRepo {
	return &branchRepo{
		fileName: fileName,
	}
}

func (b *branchRepo) Read() ([]models.Branch, error) {
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return []models.Branch{}, err
	}

	var branches []models.Branch
	err = json.Unmarshal(data, &branches)
	if err != nil {
		return []models.Branch{}, err
	}
	return branches, nil
}

func (b *branchRepo) Create(req *models.CreateBranch) (string, error) {
	branches, err := b.Read()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	branches = append(branches, models.Branch{
		Id:      uuid,
		Name:    req.Name,
	})

	body, err := json.MarshalIndent(branches, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (b *branchRepo) Delete(req *models.BranchPrimaryKey) error {
	branches, err := b.Read()
	if err != nil {
		return err
	}
	flag := true
	for i, v := range branches {
		if v.Id == req.Id {
			branches = append(branches[:i], branches[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("There is no branch with this id")
	}

	body, err := json.MarshalIndent(branches, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (b *branchRepo) Update(req *models.UpdateBranch, branchId string) error {
	branches, err := b.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range branches {
		if v.Id == branchId {
			branches[i].Name = req.Name
			flag = false
		}
	}

	if flag {
		return errors.New("There is no branch with this id")
	}

	body, err := json.MarshalIndent(branches, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (b *branchRepo) GetByID(req *models.BranchPrimaryKey) (models.Branch, error) {
	branches, err := b.Read()
	if err != nil {
		return models.Branch{}, err
	}

	for _, val := range branches {
		if val.Id == req.Id {
			return val, nil
		}
	}

	return models.Branch{}, errors.New("There is no branch with this id")
}

func (b *branchRepo) GetAll(req *models.GetBranchListRequest) (models.GetBranchListResponse, error) {
	users, err := b.Read()
	if err != nil {
		return models.GetBranchListResponse{}, err
	}

	if req.Limit+req.Offset > len(users) {
		return models.GetBranchListResponse{}, errors.New("out of range")
	}

	fBranches := []models.Branch{}
	for i := req.Offset; i < req.Offset+req.Limit; i++ {
		fBranches = append(fBranches, users[i])
	}
	return models.GetBranchListResponse{
		Branches: fBranches,
		Count: len(fBranches),
	}, nil
}