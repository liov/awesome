-- 根据日期
-- 1. 找到当前表中的日期列，并且其转换成需要排序的年月格式便可，并且取出对应的字符长度。

-- 2. 如下，我需要将金额数据按照月度汇总，那么我需要做的就是把当前日期先转换成年月格式的日期，然后按照分组。

-- 3. 需要注意的是，需要将group后的日期字段和查询列的字段都转换为年月格式的字符。如 2019-05 。


select id,CONVERT(varchar(7),date,120) group by CONVERT(varchar(7),date,120)


--  4. 常用的日期格式有以下：


select CONVERT(varchar(12) , getdate(), 101 );
--   05/19/2019

select CONVERT(varchar(12) , getdate(), 102 );
--  2019.05.19

select CONVERT(varchar(12) , getdate(), 103 );
--  19/05/2019

select CONVERT(varchar(12) , getdate(), 104 );
--  19.05.2019

select CONVERT(varchar(12) , getdate(), 105 );
--  19-05-2019

select CONVERT(varchar(12) , getdate(), 106 );
--  19 05 2019
--  ---------------------------------------------
select CONVERT(varchar(12) , getdate(), 107 );
-- 05 19, 2019

select CONVERT(varchar(12) , getdate(), 108 );
--  09:15:33

select CONVERT(varchar(12) , getdate(), 109 );
--  05 19 2019

select CONVERT(varchar(12) , getdate(), 110 );
-- 05-19-2019

select CONVERT(varchar(12) , getdate(), 111 );
-- 2019/05/19

select CONVERT(varchar(12) , getdate(), 112 );
-- 20190519

select CONVERT(varchar(12) , getdate(), 113 );
--  19 05 2019 0

select CONVERT(varchar(12) , getdate(), 114 );
--  09:16:06:747


-- 5. 如果需要去日期中相关的值，有以下方法
#
#     YEAR('2018-05-17 00:00:00'); -- 年
#     MONTH('2018-05-15 00:00:00'); -- 月
#     DAY('2008-05-15 00:00:00'); -- 日
#     DATEPART ( datepart , date );
#     DATEPART(MM,'2018-05-15 00:00:00');
#    年份 yy、yyyy
#    季度 qq、q
#    月份 mm、m
#    每年的某一日 dy、y
#    日期 dd、d
#    星期 wk、ww
#    工作日 dw
#    小时 hh
#    分钟 mi、n
#    秒 ss、s
#    毫秒 ms


template<T:DATE_FORMAT(create_time,'%Y-%m-%d')|WEEK(create_time)|MONTH(create_time)>
SELECT count(*),T AS dateTime
FROM `customer` a
WHERE a.level = 3 OR a.`level` = 2
GROUP BY T;

SELECT a.id,sum(b.contract_number),DATE_FORMAT(create_time,'%Y-%m-%d') AS dateTime FROM `trade` a,`trade_contract` b WHERE a.id = b.trade_id GROUP BY DATE_FORMAT(create_time,'%Y-%m-%d');


SELECT
    COUNT( customerNum ) / ( SELECT count( id ) FROM `customer` WHERE `level` = 4 AND create_time BETWEEN "2017-11-30T16:00:00.000Z" AND "2018-12-13T16:00:00.000Z" ) AS rate
FROM
    (
    SELECT
    count( customer_id ) AS customerNum
    FROM
    `follow` a LEFT JOIN `customer` b
    WHERE
    a.customer_id IN ( SELECT id FROM `customer` d WHERE d.`level` = 4 AND d.create_time BETWEEN "2017-11-30T16:00:00.000Z" AND "2018-12-13T16:00:00.000Z" )
    AND a.customer_level < 4
    AND DATEDIFF(b.create_time,a.create_time) <=15
    GROUP BY a.customer_id
    ) c;
