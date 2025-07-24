package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

// CONSTRUCTOR
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}


// INTERFACE
func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	scriptSql := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repository.DB.ExecContext(ctx, scriptSql, comment.Email, comment.Comment)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	comment.Id = int32(id)

	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	scriptSql := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, scriptSql, id)
	comment := entity.Comment{}
	
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada / id tidak ditemukan
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	scriptSql := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, scriptSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
// END INTERFACE