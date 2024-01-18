-- Active: 1700445944039@@159.75.155.236@3306@cgo

select *
from user_room as ur
    inner join users as u on ur.user_id = u.id
where ur.room_id = 1

select * from msg