-- first version
SELECT 
    user_id, 
    ROUND((count(user_id) FILTER (WHERE action = 'confirmed')) / (count(user_id) FILTER (WHERE action IS NOT NULL))::decimal, 2)  AS confirmation_rate  
FROM (
    SELECT s.user_id, COALESCE(c.action, 'timeout') as action   
    FROM Signups s
    LEFT JOIN Confirmations c ON c.user_id = s.user_id
) GROUP BY user_id


-- second version
SELECT 
    user_id, 
    ROUND((count(user_id) FILTER (WHERE action = 'confirmed'))/(count(user_id))::decimal, 2)  AS confirmation_rate  FROM (
    SELECT s.user_id, c.action
    FROM Signups s
    LEFT JOIN Confirmations c ON c.user_id = s.user_id
) GROUP BY user_id