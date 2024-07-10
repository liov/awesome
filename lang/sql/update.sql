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