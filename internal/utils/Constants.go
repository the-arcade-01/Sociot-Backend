package utils

/*
*   env Constants
 */
var (
	PORT       = "PORT"
	JWT_SECRET = "JWT_SECRET"
	JWT_ALGO   = "HS256"
	DB_DRIVER  = "DB_DRIVER"
	DB_URL     = "DB_URL"
)

var USER_ID = "userId"

/*
*   Get posts queries
*   Note: Don't make changes without discussion
 */

const GET_POSTS_QUERY_TAGS = `
SELECT 
    u.userId,
    u.userName,
    p.postId,
    p.title,
    p.content,
    (
        SELECT GROUP_CONCAT(t.tag) 
        FROM tags t 
        JOIN post_tags pt ON t.tagId = pt.tagId 
        WHERE pt.postId = p.postId
    ) AS tags,
    v.views,
    p.createdAt,
    p.updatedAt
FROM 
    posts p
LEFT JOIN 
    post_views v ON p.postId = v.postId
LEFT JOIN 
    users u ON p.userId = u.userId
HAVING
    tags LIKE '%v' 
ORDER BY 
    %v DESC;
`

const GET_POSTS_QUERY = `
SELECT 
    u.userId,
    u.userName,
    p.postId,
    p.title,
    p.content,
    (
        SELECT GROUP_CONCAT(t.tag) 
        FROM tags t 
        JOIN post_tags pt ON t.tagId = pt.tagId 
        WHERE pt.postId = p.postId
    ) AS tags,
    v.views,
    p.createdAt,
    p.updatedAt
FROM 
    posts p
LEFT JOIN 
    post_views v ON p.postId = v.postId
LEFT JOIN 
    users u ON p.userId = u.userId
ORDER BY 
    %v DESC;
`

/*
*   Search queries
*   Note: Don't make changes without discussion
 */

const SEARCH_USER = `
SELECT
    u.userId,
    u.userName,
    (
        SELECT COUNT(postId)
        FROM posts p
        WHERE p.userId = u.userId
    ) AS postCount,
    COALESCE((
        SELECT SUM(views)
        FROM post_views pv
        WHERE pv.postId IN 
        (
            SELECT postId FROM posts p
            WHERE p.userId = u.userId
        )
    ), 0) AS viewCount,
    u.createdAt
FROM
    users u
HAVING
    u.userName LIKE '%v'
ORDER BY
    postCount DESC;
`

const SEARCH_POST = `
SELECT 
    u.userId,
    u.userName,
    p.postId,
    p.title,
    p.content,
    (
        SELECT GROUP_CONCAT(t.tag) 
        FROM tags t 
        JOIN post_tags pt ON t.tagId = pt.tagId 
        WHERE pt.postId = p.postId
    ) AS tags,
    v.views,
    p.createdAt,
    p.updatedAt
FROM 
    posts p
LEFT JOIN 
    post_views v ON p.postId = v.postId
LEFT JOIN 
    users u ON p.userId = u.userId
HAVING
    p.title LIKE '%v' 
ORDER BY 
    v.views DESC;
`

/*
*   GetUserStats queries
*   Note: Don't make changes without discussion
 */

const GET_USER_STATS = `
SELECT
    u.userId,
    u.userName,
    (
        SELECT COUNT(postId)
        FROM posts p
        WHERE p.userId = u.userId
    ) AS postCount,
    COALESCE((
        SELECT SUM(views)
        FROM post_views pv
        WHERE pv.postId IN 
        (
            SELECT postId FROM posts p
            WHERE p.userId = u.userId
        )
    ), 0) AS viewCount,
    u.createdAt
FROM
    users u
WHERE
    u.userId = ?;
`
