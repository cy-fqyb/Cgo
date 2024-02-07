-- Active: 1700445944039@@159.75.155.236@3306@cgo

SELECT
    u.id AS friend_id,
    u.name AS friend_name,
    msg.msg AS content,
    msg.create_time AS time,
    u.avatar AS avatar
FROM users u
    LEFT JOIN msg ON msg.from_id = u.id
WHERE
    msg.create_time = (
        SELECT MAX(create_time)
        FROM msg
    )
    AND msg.to_id = 2;

SELECT
    u.id AS friend_id,
    u.name AS friend_name,
    COUNT(msg.id) AS message_count
FROM users u
    LEFT JOIN msg ON msg.from_id = u.id
WHERE
    DATE(msg.create_time) = CURDATE()
GROUP BY
    u.id,
    u.name;

select
    u.id as friend_id,
    u.name as friend_name,
    msg.msg as content,
    msg.create_time AS time,
    u.avatar as avatar,
    daily_message_counts.message_count as 'msg_count'
from
    users u
    left join msg on msg.from_id = u.id
    LEFT JOIN (
        SELECT from_id, COUNT(*) AS message_count
        FROM msg
        WHERE
            DATE(create_time) = CURDATE()
        GROUP BY
            from_id
    ) AS daily_message_counts ON u.id = daily_message_counts.from_id
WHERE
    msg.to_id = 2
    AND DATE(msg.create_time) = CURDATE() -- 仅包括今天发送的消息
    AND msg.create_time = (
        SELECT MAX(create_time)
        FROM msg
        WHERE
            msg.from_id = u.id
    );

;