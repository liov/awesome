SELECT name, COUNT(*)
FROM employee_tbl
GROUP BY name;

SELECT name, SUM(singin) as singin_count
FROM employee_tbl
GROUP BY name
WITH ROLLUP;


--  EXTRACT (component_name, FROM {datetime | interval})
-- GROUP BY GROUP_CONCAT 拼某一列
SELECT t.sid,t.name,t.sex,GROUP_CONCAT(t.num) from distinct_concat t GROUP BY t.sid,t.name,t.sex;

-- order by 在 group by之后执行，要保留第一行要做子查询 实测mysql 要加LIMIT bignum
SELECT * FROM (SELECT * FROM tsp_settle_info WHERE effect_date <= now() ORDER BY effect_date DESC LIMIT 10000000) GROUP BY tsp_id
