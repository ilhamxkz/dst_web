package models

import (
	"time"
)

// Structs mapped from the provided MySQL dump. Adjust tags/types if you need
// different behaviors (e.g. use string for timestamps instead of time.Time).

// Orders
type Order struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaPemesan string     `gorm:"type:longtext" json:"nama_pemesan"`
	Produk      string     `gorm:"type:longtext" json:"produk"`
	Jumlah      int64      `json:"jumlah"`
	TotalHarga  float64    `json:"total_harga"`
	CreatedAt   *time.Time `json:"created_at"`
}

// Users table (matches `users` in dump)
type User struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string     `gorm:"type:varchar(100);not null" json:"username"`
	Email     string     `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"-"`
	Role_id   string     `gorm:"type:varchar(20);default:'viewer'" json:"role_id"`
	Token     string     `gorm:"type:text" json:"token,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	// relations
	BlogPosts []BlogPost `gorm:"foreignKey:AuthorID" json:"blog_posts,omitempty"`
	Posts     []Post     `gorm:"foreignKey:AuthorID" json:"posts,omitempty"`
}

// Role from `roles` table
type Role struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaRole string `gorm:"type:longtext" json:"nama_role"`
}

// Menu from `menus` table
type Menu struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaMenu string `gorm:"type:longtext" json:"nama_menu"`
	Routes   string `gorm:"type:longtext" json:"routes"`
}

// Akses (akses table)
type Akses struct {
	ID     uint `gorm:"primaryKey;autoIncrement" json:"id"`
	MenuID uint `json:"menu_id"`
	RoleID uint `json:"role_id"`

	Role Role `json:"role" gorm:"foreignKey:RoleID"`
	Menu Menu `json:"menu" gorm:"foreignKey:MenuID"`
}

type ContactMessage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type CompanyProfile struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CompanyName string    `json:"company_name" gorm:"column:company_name"`
	Tagline     string    `json:"tagline" gorm:"column:tagline"`
	Description string    `json:"description" gorm:"column:description"`
	Address     string    `json:"address" gorm:"column:address"`
	Phone       string    `json:"phone" gorm:"column:phone"`
	Email       string    `json:"email" gorm:"column:email"`
	Website     string    `json:"website" gorm:"column:website"`
	Logo        string    `json:"logo" gorm:"column:logo"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// Posts (site posts table `posts`)
type Post struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     *string    `gorm:"type:varchar(255);default:null" json:"title"`
	Content   string     `gorm:"type:longtext;not null" json:"content"`
	AuthorID  *int       `gorm:"column:author_id" json:"author_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Author *User `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author,omitempty"`
}

// BlogPosts (separate table `blog_posts`)
type BlogPost struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string     `gorm:"type:varchar(150);not null" json:"title"`
	Slug       *string    `gorm:"type:varchar(150);uniqueIndex" json:"slug"`
	Content    string     `gorm:"type:text;not null" json:"content"`
	CoverImage *string    `gorm:"type:varchar(255)" json:"cover_image"`
	AuthorID   *uint      `json:"author_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`

	Author   *User         `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author,omitempty"`
	Comments []BlogComment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}

// BlogComment (blog_comments)
type BlogComment struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    *uint      `json:"post_id"`
	Name      *string    `gorm:"type:varchar(100)" json:"name"`
	Email     *string    `gorm:"type:varchar(100)" json:"email"`
	Comment   *string    `gorm:"type:text" json:"comment"`
	CreatedAt *time.Time `json:"created_at"`

	Post *BlogPost `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"post,omitempty"`
}

// Clients
type Client struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(150);not null" json:"name"`
	Logo      string     `gorm:"type:varchar(255)" json:"logo"`
	Website   string     `gorm:"type:varchar(255)" json:"website"`
	CreatedAt *time.Time `json:"created_at"`
}

// Projects
type Project struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `gorm:"type:varchar(150);not null" json:"title"`
	ClientName  *string    `gorm:"type:varchar(100)" json:"client_name"`
	Description *string    `gorm:"type:text" json:"description"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Status      string     `gorm:"type:enum('ongoing','completed','pending');default:'ongoing'" json:"status"`
	CoverImage  *string    `gorm:"type:varchar(255)" json:"cover_image"`
	CreatedAt   *time.Time `json:"created_at"`

	Images []ProjectImage `gorm:"foreignKey:ProjectID" json:"images,omitempty"`
}

// ProjectImages
type ProjectImage struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProjectID *uint   `json:"project_id"`
	ImageURL  *string `gorm:"type:varchar(255)" json:"image_url"`
	Caption   *string `gorm:"type:varchar(255)" json:"caption"`

	Project *Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"project,omitempty"`
}

// Services
type Service struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string  `gorm:"type:varchar(150);not null" json:"title"`
	Description *string `gorm:"type:text" json:"description"`
	Icon        string  `json:"icon"`   // path/logo
	Images      string  `json:"images"` // path/image
	CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
}

// experinces
type Experience struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `gorm:"type:varchar(150);not null" json:"title"`
	ClientName  string     `gorm:"type:varchar(100)" json:"client_name"`
	Description *string    `gorm:"type:text" json:"description"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Status      string     `gorm:"type:enum('ongoing','completed','paused')" json:"status"`
	Year        *int       `gorm:"type:year(4)" json:"year"`
	CoverImage  string     `gorm:"type:varchar(255)" json:"cover_image"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
}

// Team Members
type TeamMember struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	Position  *string    `gorm:"type:varchar(100)" json:"position"`
	Bio       *string    `gorm:"type:text" json:"bio"`
	Photo     *string    `gorm:"type:varchar(255)" json:"photo"`
	LinkedIn  *string    `gorm:"type:varchar(255)" json:"linkedin"`
	GitHub    *string    `gorm:"type:varchar(255)" json:"github"`
	CreatedAt *time.Time `json:"created_at"`
}

// Gallery
type Gallery struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string     `gorm:"type:varchar(150);not null" json:"title"`
	ImageURL  string     `gorm:"type:varchar(255);not null" json:"image_url"`
	Caption   *string    `gorm:"type:varchar(255)" json:"caption"`
	CreatedAt *time.Time `json:"created_at"`
}

type Pertanyaan struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Question string `gorm:"type:text;not null" json:"question"`
	Answer   string `gorm:"type:text;not null" json:"answer"`
}

type Testimoni struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama          string     `gorm:"type:varchar(100);not null" json:"nama"`
	UrlFotoProfil string     `gorm:"type:varchar(255)" json:"url_foto_profil"`
	Perusahaan    string     `gorm:"type:varchar(150)" json:"perusahaan"`
	Pesan         string     `gorm:"type:text;not null" json:"pesan"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// pelatihan
type Pelatihan struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Image     string `gorm:"type:varchar(255);not null" json:"image"`
	Category  string `gorm:"type:varchar(100);not null" json:"category"`
	Jenis     string `gorm:"type:varchar(100);not null" json:"jenis"`
	Deskripsi string `gorm:"type:text;not null" json:"deskripsi"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type HeroBackground struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Image     string    `gorm:"type:varchar(255);not null" json:"image"`
	AltText   string    `gorm:"type:varchar(255)" json:"alt_text"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}


type Statistic struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	Value       uint      `gorm:"type:int unsigned;not null;default:0" json:"value"`
	Unit        string    `gorm:"type:varchar(50);not null;default:''" json:"unit"`
	Icon        string    `gorm:"type:varchar(255);not null;default:''" json:"icon"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}