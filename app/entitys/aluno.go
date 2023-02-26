package entitys

import "time"

type AlunoE struct {
	Id        int       `json:"id"        gorm:"primary_key -all"`
	Name      string    `json:"name"      gorm:"-all"`
	Nickname  string    `json:"nickname"  gorm:"-all"`
	Email     string    `json:"email"     gorm:"-all"`
	Password  string    `json:"password"  gorm:"-all"`
	Cellphone string    `json:"cellphone" gorm:"-all"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

func (u AlunoE) GetId() int {
	return u.Id
}

func (u AlunoE) GetName() string {
	return u.Name
}

func (u AlunoE) GetNickname() string {
	return u.Nickname
}

func (u AlunoE) GetEmail() string {
	return u.Email
}
func (u AlunoE) GetPassword() string {
	return u.Password
}
func (u AlunoE) GetCellphone() string {
	return u.Cellphone
}
func (u AlunoE) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u AlunoE) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
