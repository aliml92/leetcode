CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
   DECLARE limitv INT;
   SET limitv = N - 1; 
   RETURN (
        SELECT IFNULL(
            (SELECT DISTINCT salary from employee 
            ORDER BY salary DESC
            LIMIT limitv, 1),
            NULL 
        ) AS getNthHighestSalary
   ); 
END    