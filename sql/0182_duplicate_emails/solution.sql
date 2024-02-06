select distinct p1.email as Email
from person as p1
inner join person p2 on p2.email = p1.email
where p1.id <> p2.id;