SELECT name 
FROM (
    SELECT e1.id, e1.name, count(*) FROM Employee e1
    LEFT JOIN Employee e2 ON e2.managerId = e1.id
    GROUP BY e1.id, e1.name
) WHERE count >=5;