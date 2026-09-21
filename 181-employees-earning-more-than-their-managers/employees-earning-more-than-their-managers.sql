# Write your MySQL query statement below
select name as Employee
from Employee as e
inner join (
    select id, salary
    from Employee
) as tmp
on e.managerId = tmp.id
where e.salary > tmp.salary
;