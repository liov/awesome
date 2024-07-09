use `db`;
delimiter $$　　#将语句的结束符号从分号;临时改为两个$$(可以是自定义)
CREATE PROCEDURE insert_sfm(IN s_id INTEGER,c_id INTEGER,,value VARCHAR)
BEGIN
    DECLARE id int unsigned default 0;
    INSERT INTO `db`.`table1`(`category_id`,`supplier_id`,`product_type`,`field_type`,`field_id`) VALUES(@s_id,@c_id,8,13,7);
    SET id = select MAX(id) FROM sp_field_mapper;
    INSERT INTO `db`.`table2`(`field_mapper_id`,`field_value`,`field_display_value`) VALUES(@id,19,@value);
END$$
delimiter;
CALL insert_sfm(1,2,'');