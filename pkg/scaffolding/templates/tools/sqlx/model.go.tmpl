package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel contains common columns for all tables
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// User represents a user in the system
type User struct {
	BaseModel
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"` // Never return password in JSON
	Active   bool   `json:"active" gorm:"default:true"`

	// Relationships (example)
	Profile *UserProfile `json:"profile,omitempty" gorm:"foreignKey:UserID"`
	Posts   []Post       `json:"posts,omitempty" gorm:"foreignKey:UserID"`
}

// UserProfile represents additional user information
type UserProfile struct {
	BaseModel
	UserID      uint   `json:"user_id" gorm:"not null;uniqueIndex"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth *time.Time `json:"date_of_birth"`

	// Relationship
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// Post represents a blog post or article
type Post struct {
	BaseModel
	UserID      uint   `json:"user_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Content     string `json:"content" gorm:"type:text"`
	Slug        string `json:"slug" gorm:"uniqueIndex;not null"`
	Published   bool   `json:"published" gorm:"default:false"`
	PublishedAt *time.Time `json:"published_at"`

	// Relationships
	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Tags     []Tag     `json:"tags,omitempty" gorm:"many2many:post_tags;"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}

// Tag represents a post tag
type Tag struct {
	BaseModel
	Name        string `json:"name" gorm:"uniqueIndex;not null"`
	Slug        string `json:"slug" gorm:"uniqueIndex;not null"`
	Description string `json:"description"`
	Color       string `json:"color" gorm:"default:#007bff"`

	// Relationships
	Posts []Post `json:"posts,omitempty" gorm:"many2many:post_tags;"`
}

// Comment represents a comment on a post
type Comment struct {
	BaseModel
	PostID   uint   `json:"post_id" gorm:"not null"`
	UserID   uint   `json:"user_id" gorm:"not null"`
	Content  string `json:"content" gorm:"type:text;not null"`
	Approved bool   `json:"approved" gorm:"default:false"`

	// Self-referencing for nested comments
	ParentID *uint      `json:"parent_id"`
	Parent   *Comment   `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies  []Comment  `json:"replies,omitempty" gorm:"foreignKey:ParentID"`

	// Relationships
	Post Post `json:"post,omitempty" gorm:"foreignKey:PostID"`
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// Category represents a post category
type Category struct {
	BaseModel
	Name        string `json:"name" gorm:"uniqueIndex;not null"`
	Slug        string `json:"slug" gorm:"uniqueIndex;not null"`
	Description string `json:"description"`

	// Self-referencing for nested categories
	ParentID   *uint      `json:"parent_id"`
	Parent     *Category  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children   []Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`
}

// Setting represents application settings
type Setting struct {
	Key   string `json:"key" gorm:"primarykey"`
	Value string `json:"value" gorm:"type:text"`
	Type  string `json:"type" gorm:"default:string"` // string, int, bool, json

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuditLog represents audit trail for important actions
type AuditLog struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Action    string    `json:"action" gorm:"not null"`
	Resource  string    `json:"resource" gorm:"not null"`
	ResourceID uint     `json:"resource_id"`
	OldValues string    `json:"old_values,omitempty" gorm:"type:text"`
	NewValues string    `json:"new_values,omitempty" gorm:"type:text"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	CreatedAt time.Time `json:"created_at"`

	// Relationship
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName methods (optional - to specify custom table names)

// TableName overrides the table name for User model
func (User) TableName() string {
	return "users"
}

// TableName overrides the table name for UserProfile model
func (UserProfile) TableName() string {
	return "user_profiles"
}

// TableName overrides the table name for Post model
func (Post) TableName() string {
	return "posts"
}

// TableName overrides the table name for Tag model
func (Tag) TableName() string {
	return "tags"
}

// TableName overrides the table name for Comment model
func (Comment) TableName() string {
	return "comments"
}

// TableName overrides the table name for Category model
func (Category) TableName() string {
	return "categories"
}

// TableName overrides the table name for Setting model
func (Setting) TableName() string {
	return "settings"
}

// TableName overrides the table name for AuditLog model
func (AuditLog) TableName() string {
	return "audit_logs"
}

// Model validation methods (examples)

// BeforeCreate is a GORM hook that runs before creating a record
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Example: Hash password, generate UUID, etc.
	return nil
}

// BeforeUpdate is a GORM hook that runs before updating a record
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	// Example: Update slug, validate fields, etc.
	return nil
}

// AfterCreate is a GORM hook that runs after creating a record
func (u *User) AfterCreate(tx *gorm.DB) error {
	// Example: Send welcome email, create default profile, etc.
	return nil
}

// Helper methods

// IsActive checks if user is active
func (u *User) IsActive() bool {
	return u.Active && u.DeletedAt.Time.IsZero()
}

// GetFullName returns user's full name from profile
func (u *User) GetFullName() string {
	if u.Profile != nil {
		return u.Profile.FirstName + " " + u.Profile.LastName
	}
	return u.Name
}

// IsPublished checks if post is published
func (p *Post) IsPublished() bool {
	return p.Published && p.PublishedAt != nil
}

// GetExcerpt returns first 100 characters of content
func (p *Post) GetExcerpt() string {
	if len(p.Content) > 100 {
		return p.Content[:100] + "..."
	}
	return p.Content
}

// IsApproved checks if comment is approved
func (c *Comment) IsApproved() bool {
	return c.Approved
}

// HasParent checks if comment has a parent (is a reply)
func (c *Comment) HasParent() bool {
	return c.ParentID != nil
}

// GetValue returns setting value based on type
func (s *Setting) GetValue() interface{} {
	switch s.Type {
	case "int":
		// Convert to int
		return s.Value
	case "bool":
		return s.Value == "true"
	case "json":
		// Parse JSON
		return s.Value
	default:
		return s.Value
	}
}
