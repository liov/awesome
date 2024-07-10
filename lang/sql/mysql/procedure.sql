use `db`;
delimiter //　　-- 将语句的结束符号从分号;临时改为两个$$(可以是自定义)
CREATE PROCEDURE insert_sfm(IN s_id INTEGER,c_id INTEGER,value VARCHAR)
BEGIN
    DECLARE id int unsigned default 0;
    INSERT INTO `db`.`table1`(`category_id`,`supplier_id`,`product_type`,`field_type`,`field_id`) VALUES(@s_id,@c_id,8,13,7);
    SET id = select MAX(id) FROM sp_field_mapper;
    INSERT INTO `db`.`table2`(`field_mapper_id`,`field_value`,`field_display_value`) VALUES(@id,19,@value);
END //
delimiter;
CALL insert_sfm(1,2,'');

DELIMITER //
CREATE
    DEFINER = `web`@`%` PROCEDURE `insert`()
BEGIN
    declare i int;
    set i = 6001;
    while i < 7001
    do
        insert into customer_erptask(customer_id) values (i);
        insert into customer_extra_info(customer_id) values (i);
    set i = i + 1;
    end while;

END //
DELIMITER ;

        -- 您提供的SQL语句定义了一个名为`insert`的存储过程，该存储过程属于`web`用户，并可以在任何主机上运行
--
--                                   以下是存储过程的逐行解释：
--
--                                   1. `CREATE DEFINER = `web`@`%` PROCEDURE `insert`()`：创建一个名为`insert`的存储过程，定义者为`web`用户，可以在任何主机上运行（`%`表示任何主机）。
-- 2. `BEGIN`：存储过程的开始。
-- 3. `declare i int;`：声明一个名为`i`的整数变量。
-- 4. `set i = 6001;`：将变量`i`的值设置为6001。
-- 5. `while i < 7001`：开始一个`WHILE`循环，条件是`i`小于7001。
-- 6. `do`：`WHILE`循环的主体开始。
-- 7. `insert into customer_erptask(customer_id) values (i);`：向`customer_erptask`表插入一行数据，其中`customer_id`字段的值为当前的`i`值。
-- 8. `insert into customer_extra_info(customer_id) values (i);`：向`customer_extra_info`表插入一行数据，其中`customer_id`字段的值为当前的`i`值。
-- 9. `set i = i + 1;`：将变量`i`的值增加1。
-- 10. `end while;`：`WHILE`循环的结束。
-- 11. `END;`：存储过程的结束。
--
-- 当调用这个存储过程时，它将在`customer_erptask`和`customer_extra_info`表中插入数据，`customer_id`的值从6001到7000。