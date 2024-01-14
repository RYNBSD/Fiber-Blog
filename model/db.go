package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `sql:"VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4()" json:"id"`
	Username  string    `sql:"VARCHAR(50) NOT NULL" json:"username"`
	Email     string    `sql:"VARCHAR(50) NOT NULL UNIQUE" json:"email"`
	Password  string    `sql:"VARCHAR(70) NOT NULL" json:"password"`
	Picture   string    `sql:"VARCHAR(70) NOT NULL" json:"picture"`
	CreatedAt time.Time `sql:"DATETIME NOT NULL DEFAULT NOW()" json:"createdAt"`
	UpdatedAt time.Time `sql:"DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()" json:"updatedAt"`
}

type Blog struct {
	Id          uuid.UUID `sql:"VARCHAR(36) PRIMARY KEY DEFAULT uuid_generate_v4()" json:"id"`
	Title       string    `sql:"VARCHAR(255) NOT NULL" json:"title"`
	Description string    `sql:"VARCHAR(1000) NULL" json:"description"`
	BloggerId   string    `sql:"VARCHAR(36) NOT NULL REFERENCES user(id) ON DELETE CASCADE" json:"bloggerId"`
	CreatedAt   time.Time `sql:"DATETIME NOT NULL DEFAULT NOW()" json:"createdAt"`
	UpdatedAt   time.Time `sql:"DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()" json:"updatedAt"`
}

type BlogImages struct {
	Id        int       `sql:"BIGINT PRIMARY KEY AUTO_INCREMENT" json:"id"`
	Image     string    `sql:"VARCHAR(100) NOT NULL" json:"image"`
	BlogId    uuid.UUID `sql:"VARCHAR(36) NOT NULL REFERENCES blog(id) ON DELETE CASCADE" json:"blogId"`
	CreatedAt time.Time `sql:"DATETIME NOT NULL DEFAULT NOW()" json:"createdAt"`
}

type BlogLikes struct {
	Id        int       `sql:"BIGINT PRIMARY KEY AUTO_INCREMENT" json:"id"`
	BlogId    uuid.UUID `sql:"VARCHAR(36) NOT NULL REFERENCES blog(id) ON DELETE CASCADE" json:"blogId"`
	LikerId   uuid.UUID `sql:"VARCHAR(36) NOT NULL REFERENCES user(id) ON DELETE CASCADE" json:"likerId"`
	CreatedAt time.Time `sql:"DATETIME NOT NULL DEFAULT NOW()" json:"createdAt"`
}

type BlogComments struct {
	Id          int       `sql:"BIGINT PRIMARY KEY AUTO_INCREMENT" json:"id"`
	Comment     string    `sql:"VARCHAR(255) NOT NULL" json:"comment"`
	BlogId      uuid.UUID `sql:"VARCHAR(36) NOT NULL REFERENCES blog(id) ON DELETE CASCADE" json:"blogId"`
	CommenterId uuid.UUID `sql:"VARCHAR(36) NOT NULL REFERENCES user(id) ON DELETE CASCADE" json:"commenterId"`
	CreatedAt   time.Time `sql:"DATETIME NOT NULL DEFAULT NOW()" json:"createdAt"`
	UpdatedAt   time.Time `sql:"DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()" json:"updatedAt"`
}

type Error struct {
	Id        int       `sql:"BIGINT PRIMARY KEY AUTO_INCREMENT" json:"id"`
	Status    int16     `sql:"INT NOT NULL" json:"status"`
	Message   string    `sql:"VARCHAR(255) NOT NULL" json:"message"`
	IsFixed   bool      `sql:"BOOL NOT NULL DEFAULT FALSE" json:"isFixed"`
	CreatedAt time.Time `sql:"DATETIME NOT NULL DEFAULT NOW()" json:"createdAt"`
	UpdatedAt time.Time `sql:"DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()" json:"updatedAt"`
}
