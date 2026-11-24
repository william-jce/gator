-- name: CreateBookmark :one
INSERT INTO bookmarks(user_id, post_id, created_at)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: DeleteBookmark :execrows
DELETE FROM bookmarks WHERE user_id = $1 AND post_id = $2;

-- name: GetBookmarksForUser :many
SELECT posts.*, bookmarks.created_at AS bookmarked_at
FROM bookmarks
INNER JOIN posts ON posts.id = bookmarks.post_id
WHERE bookmarks.user_id = $1
ORDER BY bookmarks.created_at DESC;

-- name: GetBookmarkedPostByIDPrefix :many
SELECT posts.*
FROM bookmarks
INNER JOIN posts on posts.id = bookmarks.post_id
WHERE bookmarks.user_id = sqlc.arg(user_id)
    AND posts.id::text LIKE sqlc.arg(id_prefix)::text || '%'
LIMIT 4;
