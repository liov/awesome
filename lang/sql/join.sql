--INNER JOIN（内连接,或等值连接）：获取两个表中字段匹配关系的记录。
--LEFT JOIN（左连接）：获取左表所有记录，即使右表没有对应匹配的记录。
--RIGHT JOIN（右连接）： 与 LEFT JOIN 相反，用于获取右表所有记录，即使左表没有对应匹配的记录。
--full join:外连接，返回两个表中的行：left join + right join。
--cross join:结果是笛卡尔积，就是第一个表的行数乘以第二个表的行数。

-- 两个表在，join时，首先做一个笛卡尔积，on后面的条件是对这个笛卡尔积做一个过滤形成一张临时表，如果没有where就直接返回结果，如果有where就对上一步的临时表再进行过滤。
--
-- 在使用left  jion时，on和where条件的区别如下：
--
-- 1、on条件是在生成临时表时使用的条件，它不管on中的条件是否为真，都会返回左边表中的记录。
--
-- 2、where条件是在临时表生成好后，再对临时表进行过滤的条件。这时已经没有left  join的含义（必须返回左边表的记录）了，条件不为真的就全部过滤掉


-- join on:on后边写条件，以后边的条件为准生成一个临时表存储数据。

-- on和where的区别：
--
-- 1. 对于left join，不管on后面跟什么条件，左表的数据全部查出来，因此要想过滤需把条件放到where后面
--
-- 2. 对于inner join，满足on后面的条件表的数据才能查出，可以起到过滤作用。也可以把条件放到where后面。


SELECT count(*)                                                as t_count,
       count(case when field1 - field2 <= field3 then '1' end) as w_count,
       count(case when field1 = field2 then '1' end)           as f_count
FROM `table2`
         left join table1 on table2.table2_id = table1.id
    AND table1.status = 0 AND table1.field4 = 0;

SELECT count(*)                                                as t_count,
       count(case when field1 - field2 <= field3 then '1' end) as w_count,
       count(case when field1 = field2 then '1' end)           as f_count
FROM `table2`
         right join table1 on table2.table2_id = table1.id
    AND table1.status = 0 AND table1.field4 = 0;

SELECT count(*) as t_count;
SELECT count(case when table1.field4 = 0 then '1' end) as t_count;
