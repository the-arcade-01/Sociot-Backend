package utils

// ENV Constants
var PORT = "PORT"
var JWT_SECRET = "JWT_SECRET"
var JWT_ALGO = "HS256"
var DB_DRIVER = "DB_DRIVER"
var DB_URL = "DB_URL"

var USER_ID = "userId"

/*
*   Get post queries
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
