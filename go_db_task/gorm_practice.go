package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/**
进阶gorm
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： UserBlob （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 UserBlob 、 Post 和 Comment 模型，其中 UserBlob 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
// UserBlob 模型（用户）
type UserBlob struct {
	//ID        uint           `gorm:"primarykey"` // 主键ID
	//CreatedAt time.Time      `gorm:"index"`      // 创建时间（索引）
	//UpdatedAt time.Time      // 更新时间
	//DeletedAt gorm.DeletedAt `gorm:"index"`                             // 软删除标记（索引）
	gorm.Model
	Username  string `gorm:"type:varchar(50);unique;not null"`  // 用户名（唯一、非空）
	Email     string `gorm:"type:varchar(100);unique;not null"` // 邮箱（唯一、非空）
	Password  string `gorm:"type:varchar(100);not null"`        // 密码（哈希存储）
	PostCount int    `gorm:"default:0"`                         // 新增：文章数量统计
	Posts     []Post `gorm:"foreignKey:UserID"`                 // 关联的文章（一对多）
}

// Post 模型（文章）
type Post struct {
	//ID        uint           `gorm:"primarykey"` // 主键ID
	//CreatedAt time.Time      `gorm:"index"`      // 创建时间（索引）
	//UpdatedAt time.Time      // 更新时间
	//DeletedAt gorm.DeletedAt `gorm:"index"`                      // 软删除标记（索引）
	gorm.Model
	Title         string    `gorm:"type:varchar(200);not null"`     // 文章标题（非空）
	Content       string    `gorm:"type:text;not null"`             // 文章内容（非空）
	UserID        uint      `gorm:"not null;index"`                 // 关联的用户ID（外键、索引）
	CommentStatus string    `gorm:"type:varchar(20);default:'有评论'"` // 新增：评论状态
	User          UserBlob  `gorm:"foreignKey:UserID"`              // 关联的用户（多对一）
	Comments      []Comment `gorm:"foreignKey:PostID"`              // 关联的评论（一对多）
}



// Comment 模型（评论）
type Comment struct {
	//ID        uint           `gorm:"primarykey"` // 主键ID
	//CreatedAt time.Time      `gorm:"index"`      // 创建时间（索引）
	//UpdatedAt time.Time      // 更新时间
	//DeletedAt gorm.DeletedAt `gorm:"index"`              // 软删除标记（索引）
	gorm.Model
	Content string   `gorm:"type:text;not null"` // 评论内容（非空）
	PostID  uint     `gorm:"not null;index"`     // 关联的文章ID（外键、索引）
	Post    Post     `gorm:"foreignKey:PostID"`  // 关联的文章（多对一）
	UserID  uint     `gorm:"not null;index"`     // 评论者ID（外键、索引）
	User    UserBlob `gorm:"foreignKey:UserID"`  // 关联的用户（多对一）
}

const dsn = "root:JI-109385147-ch@tcp(47.92.123.15:3306)/go_lesson_db?charset=utf8mb4&parseTime=True"

func createTableByGorm() {
	// 数据库连接配置（替换为实际连接信息）
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 启用日志模式（开发环境推荐，生产环境可关闭）
		Logger: logger.Default.LogMode(logger.Info),
	})

	// 连接数据库
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 自动迁移创建表（会根据模型自动创建或更新表结构）
	err = db.AutoMigrate(
		&UserBlob{},
		&Post{},
		&Comment{},
	)
	if err != nil {
		panic("表创建失败: " + err.Error())
	}

	println("博客系统表创建成功")
}

// GetUserPostsWithComments 查询指定用户的所有文章及对应的评论
func GetUserPostsWithComments(userID uint, db *gorm.DB) ([]Post, error) {
	var user UserBlob

	// 1. 先查询用户，并预加载文章
	if err := db.Preload("Posts").First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}

	// 2. 为每篇文章预加载评论
	for i := range user.Posts {
		if err := db.Model(&user.Posts[i]).Preload("Comments").Error; err != nil {
			return nil, fmt.Errorf("查询文章评论失败: %v", err)
		}
	}

	return user.Posts, nil
}

// GetMostCommentedPost 查询评论数量最多的文章
func GetMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post

	// 使用子查询统计每篇文章的评论数，并按评论数降序排序
	err := db.Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("left join comments on posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count desc").
		First(&post).Error

	if err != nil {
		return post, fmt.Errorf("查询失败: %v", err)
	}

	return post, nil
}

// BeforeCreate Post的创建前钩子：在文章创建前更新用户的文章数量
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 1. 验证UserID有效性
	if p.UserID == 0 {
		return fmt.Errorf("用户ID不能为空")
	}

	// 2. 更新用户的文章数量（+1）
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1")).
		Error
}

// AfterDelete Comment的删除后钩子：评论删除后检查文章评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 1. 查询当前文章的剩余评论数量
	var commentCount int64
	if err := tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&commentCount).Error; err != nil {
		return err
	}

	// 2. 如果评论数量为0，更新文章的评论状态
	if commentCount == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").
			Error
	}

	return nil
}
