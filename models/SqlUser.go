package models

type SqlUser struct {
	Uid           int    `gorm:"column:uid;primary_key;AUTO_INCREMENT"`
	Username      string `gorm:"column:username;unique;NOT NULL"`
	DisplayName   string `gorm:"column:display_name"`
	Email         string `gorm:"column:email"`
	Quota         string `gorm:"column:quota"`
	Home          string `gorm:"column:home"`
	Password      string `gorm:"column:password;NOT NULL"`
	Active        int    `gorm:"column:active;default:1;NOT NULL"`
	Disabled      int    `gorm:"column:disabled;default:0;NOT NULL"`
	ProvideAvatar int    `gorm:"column:provide_avatar;default:0;NOT NULL"`
	Salt          string `gorm:"column:salt"`
}

func (m *SqlUser) TableName() string {
	return "sql_user"
}
