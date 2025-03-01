// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: post.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts
(
    user_id,
    title,
    content,
    description,
    is_draft,
    is_private,
    thumbnail_img_id
)
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_id, title, description, content, created_at, deleted_at, last_updated_at, last_updated_user_id, is_draft, is_private, thumbnail_img_id
`

type CreatePostParams struct {
	UserID         int32          `json:"user_id"`
	Title          string         `json:"title"`
	Content        string         `json:"content"`
	Description    sql.NullString `json:"description"`
	IsDraft        bool           `json:"is_draft"`
	IsPrivate      bool           `json:"is_private"`
	ThumbnailImgID sql.NullInt32  `json:"thumbnail_img_id"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.UserID,
		arg.Title,
		arg.Content,
		arg.Description,
		arg.IsDraft,
		arg.IsPrivate,
		arg.ThumbnailImgID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.LastUpdatedAt,
		&i.LastUpdatedUserID,
		&i.IsDraft,
		&i.IsPrivate,
		&i.ThumbnailImgID,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
UPDATE posts
SET
    deleted_at = now()
WHERE
    id = $1
`

func (q *Queries) DeletePost(ctx context.Context, postID int32) error {
	_, err := q.db.ExecContext(ctx, deletePost, postID)
	return err
}

const getPostById = `-- name: GetPostById :one
SELECT
    id, user_id, title, description, content, created_at, deleted_at, last_updated_at, last_updated_user_id, is_draft, is_private, thumbnail_img_id
FROM
    posts
WHERE
    id = $1
`

func (q *Queries) GetPostById(ctx context.Context, postID int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostById, postID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.LastUpdatedAt,
		&i.LastUpdatedUserID,
		&i.IsDraft,
		&i.IsPrivate,
		&i.ThumbnailImgID,
	)
	return i, err
}

const getPostHistoryById = `-- name: GetPostHistoryById :one
SELECT
    id as post_history_id,
    post_id,
    user_id,
    title,
    content,
    created_at,
    deleted_at,
    last_updated_at,
    last_updated_user_id,
    is_draft,
    is_private,
    description,
    thumbnail_img_id
FROM post_history WHERE id = $1
`

type GetPostHistoryByIdRow struct {
	PostHistoryID     int32          `json:"post_history_id"`
	PostID            int32          `json:"post_id"`
	UserID            int32          `json:"user_id"`
	Title             string         `json:"title"`
	Content           string         `json:"content"`
	CreatedAt         time.Time      `json:"created_at"`
	DeletedAt         sql.NullTime   `json:"deleted_at"`
	LastUpdatedAt     sql.NullTime   `json:"last_updated_at"`
	LastUpdatedUserID sql.NullInt32  `json:"last_updated_user_id"`
	IsDraft           bool           `json:"is_draft"`
	IsPrivate         bool           `json:"is_private"`
	Description       sql.NullString `json:"description"`
	ThumbnailImgID    sql.NullInt32  `json:"thumbnail_img_id"`
}

func (q *Queries) GetPostHistoryById(ctx context.Context, postHistoryID int32) (GetPostHistoryByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getPostHistoryById, postHistoryID)
	var i GetPostHistoryByIdRow
	err := row.Scan(
		&i.PostHistoryID,
		&i.PostID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.LastUpdatedAt,
		&i.LastUpdatedUserID,
		&i.IsDraft,
		&i.IsPrivate,
		&i.Description,
		&i.ThumbnailImgID,
	)
	return i, err
}

const getPostHistoryByPostId = `-- name: GetPostHistoryByPostId :many
SELECT
    id as post_history_id,
    post_id,
    user_id,
    title,
    created_at,
    deleted_at,
    last_updated_at,
    last_updated_user_id,
    is_draft,
    is_private,
    description,
    thumbnail_img_id
FROM post_history WHERE post_id = $1 ORDER BY created_at DESC
`

type GetPostHistoryByPostIdRow struct {
	PostHistoryID     int32          `json:"post_history_id"`
	PostID            int32          `json:"post_id"`
	UserID            int32          `json:"user_id"`
	Title             string         `json:"title"`
	CreatedAt         time.Time      `json:"created_at"`
	DeletedAt         sql.NullTime   `json:"deleted_at"`
	LastUpdatedAt     sql.NullTime   `json:"last_updated_at"`
	LastUpdatedUserID sql.NullInt32  `json:"last_updated_user_id"`
	IsDraft           bool           `json:"is_draft"`
	IsPrivate         bool           `json:"is_private"`
	Description       sql.NullString `json:"description"`
	ThumbnailImgID    sql.NullInt32  `json:"thumbnail_img_id"`
}

func (q *Queries) GetPostHistoryByPostId(ctx context.Context, postID int32) ([]GetPostHistoryByPostIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostHistoryByPostId, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostHistoryByPostIdRow{}
	for rows.Next() {
		var i GetPostHistoryByPostIdRow
		if err := rows.Scan(
			&i.PostHistoryID,
			&i.PostID,
			&i.UserID,
			&i.Title,
			&i.CreatedAt,
			&i.DeletedAt,
			&i.LastUpdatedAt,
			&i.LastUpdatedUserID,
			&i.IsDraft,
			&i.IsPrivate,
			&i.Description,
			&i.ThumbnailImgID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPosts = `-- name: GetPosts :many
WITH cte AS (
    SELECT
        id, user_id, title, description, content, created_at, deleted_at, last_updated_at, last_updated_user_id, is_draft, is_private, thumbnail_img_id, thumbnail_img_url, entity_id, module_id, module_type, module_type_id, tags
    FROM get_posts($3, $4, $5::int[], $6, $7, $8, $9, $10, 0, 0)
)
SELECT
    CAST((SELECT count(*) FROM cte) as integer) as total_count,
    cte.id, cte.user_id, cte.title, cte.description, cte.content, cte.created_at, cte.deleted_at, cte.last_updated_at, cte.last_updated_user_id, cte.is_draft, cte.is_private, cte.thumbnail_img_id, cte.thumbnail_img_url, cte.entity_id, cte.module_id, cte.module_type, cte.module_type_id, cte.tags
FROM cte
LIMIT $2
OFFSET $1
`

type GetPostsParams struct {
	PageOffset     int32          `json:"page_offset"`
	PageLimit      int32          `json:"page_limit"`
	IsPrivate      sql.NullBool   `json:"is_private"`
	IsDraft        sql.NullBool   `json:"is_draft"`
	Tags           []int32        `json:"tags"`
	UserID         sql.NullInt32  `json:"user_id"`
	ModuleID       sql.NullInt32  `json:"module_id"`
	ModuleType     NullModuleType `json:"module_type"`
	OrderBy        sql.NullString `json:"order_by"`
	OrderDirection sql.NullString `json:"order_direction"`
}

type GetPostsRow struct {
	TotalCount        int32          `json:"total_count"`
	ID                int32          `json:"id"`
	UserID            int32          `json:"user_id"`
	Title             string         `json:"title"`
	Description       sql.NullString `json:"description"`
	Content           string         `json:"content"`
	CreatedAt         time.Time      `json:"created_at"`
	DeletedAt         sql.NullTime   `json:"deleted_at"`
	LastUpdatedAt     sql.NullTime   `json:"last_updated_at"`
	LastUpdatedUserID sql.NullInt32  `json:"last_updated_user_id"`
	IsDraft           bool           `json:"is_draft"`
	IsPrivate         bool           `json:"is_private"`
	ThumbnailImgID    sql.NullInt32  `json:"thumbnail_img_id"`
	ThumbnailImgUrl   sql.NullString `json:"thumbnail_img_url"`
	EntityID          sql.NullInt32  `json:"entity_id"`
	ModuleID          sql.NullInt32  `json:"module_id"`
	ModuleType        NullModuleType `json:"module_type"`
	ModuleTypeID      sql.NullInt32  `json:"module_type_id"`
	Tags              []int32        `json:"tags"`
}

func (q *Queries) GetPosts(ctx context.Context, arg GetPostsParams) ([]GetPostsRow, error) {
	rows, err := q.db.QueryContext(ctx, getPosts,
		arg.PageOffset,
		arg.PageLimit,
		arg.IsPrivate,
		arg.IsDraft,
		pq.Array(arg.Tags),
		arg.UserID,
		arg.ModuleID,
		arg.ModuleType,
		arg.OrderBy,
		arg.OrderDirection,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsRow{}
	for rows.Next() {
		var i GetPostsRow
		if err := rows.Scan(
			&i.TotalCount,
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.CreatedAt,
			&i.DeletedAt,
			&i.LastUpdatedAt,
			&i.LastUpdatedUserID,
			&i.IsDraft,
			&i.IsPrivate,
			&i.ThumbnailImgID,
			&i.ThumbnailImgUrl,
			&i.EntityID,
			&i.ModuleID,
			&i.ModuleType,
			&i.ModuleTypeID,
			pq.Array(&i.Tags),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByIDs = `-- name: GetPostsByIDs :many
SELECT id, user_id, title, description, content, created_at, deleted_at, last_updated_at, last_updated_user_id, is_draft, is_private, thumbnail_img_id FROM posts WHERE id = ANY($1::int[])
`

func (q *Queries) GetPostsByIDs(ctx context.Context, postIds []int32) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByIDs, pq.Array(postIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.CreatedAt,
			&i.DeletedAt,
			&i.LastUpdatedAt,
			&i.LastUpdatedUserID,
			&i.IsDraft,
			&i.IsPrivate,
			&i.ThumbnailImgID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByUserId = `-- name: GetPostsByUserId :many
SELECT
    id, user_id, title, description, content, created_at, deleted_at, last_updated_at, last_updated_user_id, is_draft, is_private, thumbnail_img_id, thumbnail_img_url, entity_id, module_id, module_type, module_type_id, tags
FROM
    view_posts
WHERE
    user_id = $1 AND
    deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $3
OFFSET $2
`

type GetPostsByUserIdParams struct {
	UserID     int32 `json:"user_id"`
	PageOffset int32 `json:"page_offset"`
	PageLimit  int32 `json:"page_limit"`
}

func (q *Queries) GetPostsByUserId(ctx context.Context, arg GetPostsByUserIdParams) ([]ViewPost, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByUserId, arg.UserID, arg.PageOffset, arg.PageLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ViewPost{}
	for rows.Next() {
		var i ViewPost
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.CreatedAt,
			&i.DeletedAt,
			&i.LastUpdatedAt,
			&i.LastUpdatedUserID,
			&i.IsDraft,
			&i.IsPrivate,
			&i.ThumbnailImgID,
			&i.ThumbnailImgUrl,
			&i.EntityID,
			&i.ModuleID,
			&i.ModuleType,
			&i.ModuleTypeID,
			pq.Array(&i.Tags),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertPostHistory = `-- name: InsertPostHistory :one
INSERT INTO post_history (
    post_id,
    user_id,
    title,
    content,
    created_at,
    deleted_at,
    last_updated_at,
    last_updated_user_id,
    is_draft,
    is_private,
    description,
    thumbnail_img_id
)
SELECT
    id,
    user_id,
    title,
    content,
    created_at,
    deleted_at,
    last_updated_at,
    last_updated_user_id,
    is_draft,
    is_private,
    description,
    thumbnail_img_id
FROM
    posts
WHERE
    posts.id = $1
RETURNING id, post_id, user_id, title, description, content, created_at, deleted_at, last_updated_at, last_updated_user_id, is_draft, is_private, thumbnail_img_id
`

func (q *Queries) InsertPostHistory(ctx context.Context, postID int32) (PostHistory, error) {
	row := q.db.QueryRowContext(ctx, insertPostHistory, postID)
	var i PostHistory
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.LastUpdatedAt,
		&i.LastUpdatedUserID,
		&i.IsDraft,
		&i.IsPrivate,
		&i.ThumbnailImgID,
	)
	return i, err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET
    title = COALESCE($1, title),
    content = COALESCE($2, content),
    description = COALESCE($3, description),
    is_draft = COALESCE($4, is_draft),
    is_private = COALESCE($5, is_private),
    last_updated_user_id = $6,
    last_updated_at = now(),
    thumbnail_img_id = CASE WHEN $7 = 0 THEN NULL ELSE COALESCE($7, thumbnail_img_id) END
WHERE
    posts.id = $8
RETURNING id, user_id, title, description, content, created_at, deleted_at, last_updated_at, last_updated_user_id, is_draft, is_private, thumbnail_img_id
`

type UpdatePostParams struct {
	Title             sql.NullString `json:"title"`
	Content           sql.NullString `json:"content"`
	Description       sql.NullString `json:"description"`
	IsDraft           sql.NullBool   `json:"is_draft"`
	IsPrivate         sql.NullBool   `json:"is_private"`
	LastUpdatedUserID sql.NullInt32  `json:"last_updated_user_id"`
	ThumbnailImgID    interface{}    `json:"thumbnail_img_id"`
	PostID            int32          `json:"post_id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.Title,
		arg.Content,
		arg.Description,
		arg.IsDraft,
		arg.IsPrivate,
		arg.LastUpdatedUserID,
		arg.ThumbnailImgID,
		arg.PostID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.LastUpdatedAt,
		&i.LastUpdatedUserID,
		&i.IsDraft,
		&i.IsPrivate,
		&i.ThumbnailImgID,
	)
	return i, err
}
