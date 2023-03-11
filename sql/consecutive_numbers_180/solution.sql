-- first solution
select distinct num as ConsecutiveNums 
from logs l1
where num = (select num from logs l2 where l2.id = l1.id+1) 
and num = (select num from logs l3 where l3.id = l1.id+2); 

-- second solution
select distinct l1.num as consecutiveNums 
from logs l1
inner join logs l2 on l1.id=l2.id+1
inner join logs l3 on l2.id=l3.id+1
where l1.num=l2.num and l2.num=l3.num