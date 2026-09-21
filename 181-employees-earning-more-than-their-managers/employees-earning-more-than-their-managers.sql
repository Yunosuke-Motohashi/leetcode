# Write your MySQL query statement below
-- select e.name as Employee
-- from Employee as e
-- join Employee as tmp
-- on e.managerId = tmp.id
-- where e.salary > tmp.salary
-- ;

select e.name as Employee
from Employee as e
where e.salary > (
    select m.salary
    from Employee m
    where e.managerId = m.id
)
;