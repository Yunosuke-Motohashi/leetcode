# Write your MySQL query statement below
select e.name as Employee
from Employee as e
join Employee as tmp
on e.managerId = tmp.id
where e.salary > tmp.salary
;