-- You can't specify target table 'product_attr' for update in FROM clause
UPDATE customer_info
SET status     = 1,
    created_at = '2020-08-06 00:00:00',
    updated_at = '2020-08-06 00:00:00',
    succeeded  = 1,
    stage      = 2
WHERE id >= (SELECT id FROM (SELECT id FROM `customer_info` WHERE customer_num = '2020') temp);

# upsert
INSERT INTO user_role(id, role_id) VALUES (1, 1) ON DUPLICATE KEY UPDATE role_id = 1;

-- 去重
UPDATE report_receivers
SET is_deleted = 1
WHERE report_id IN (SELECT report_id
                    FROM (SELECT report_id FROM report_receivers GROUP BY report_id, emp_id HAVING COUNT(*) > 1) a)
  AND id NOT IN (SELECT id FROM (SELECT id FROM report_receivers GROUP BY report_id, emp_id) b)

-- 要根据一个表（假设为 B 表）的数据来更新另一个表（假设为 A 表），并且当 B 表中的某些字段（例如 a, b, c）与 A 表中相应的字段一致时，将 A 表的某个字段（例如 d）设置为 B 表的 id，你可以使用 SQL 的 UPDATE 语句结合 JOIN 来实现。
UPDATE A
SET d = B.id
FROM B
WHERE A.a = B.a
  AND A.b = B.b
  AND A.c = B.c;