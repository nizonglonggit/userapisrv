package define

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    NickName        string `json:"nick_name"`
    PassWord        string `json:"pass_word"`
    Gender          string `json:"gender"`
    Phone           string `json:"phone"`
    Email           string `json:"email"`
    HeadPortraitUrl string `json:"head_portrait_url"`
    RegDate         string `json:"reg_date"`
    Description     string `json:"description"`
    CreatedAt       int64
    DeletedAt       int64
    UpdatedAt       int64
}

type Role struct {
    ID          uint8  `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type Permission struct {
    ID          uint8  `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

type RolePermission struct {
    RoleID       uint8 `gorm:"primaryKey" json:"role_id"`
    PermissionID uint8 `gorm:"primaryKey" json:"permission_id"`
    CreatedAt    int64
    DeletedAt    int64
    UpdatedAt    int64
}

type UserRole struct {
    UserID    uint  `gorm:"primaryKey" json:"user_id"`
    RoleID    uint8 `gorm:"primaryKey" json:"role_id"`
    CreatedAt int64
    DeletedAt int64
    UpdatedAt int64
}
