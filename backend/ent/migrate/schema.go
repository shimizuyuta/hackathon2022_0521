// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserSkillsColumns holds the columns for the "user_skills" table.
	UserSkillsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "skill_id", Type: field.TypeInt},
	}
	// UserSkillsTable holds the schema information for the "user_skills" table.
	UserSkillsTable = &schema.Table{
		Name:       "user_skills",
		Columns:    UserSkillsColumns,
		PrimaryKey: []*schema.Column{UserSkillsColumns[0], UserSkillsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_skills_user_id",
				Columns:    []*schema.Column{UserSkillsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_skills_skill_id",
				Columns:    []*schema.Column{UserSkillsColumns[1]},
				RefColumns: []*schema.Column{SkillsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PostsTable,
		SkillsTable,
		UsersTable,
		UserSkillsTable,
	}
)

func init() {
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	UserSkillsTable.ForeignKeys[0].RefTable = UsersTable
	UserSkillsTable.ForeignKeys[1].RefTable = SkillsTable
}
