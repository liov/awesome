INSERT INTO `customer_extra_info`(customer_id, last_visit_time);
SELECT customer_id,
       MAX(last_visit_time)
FROM (
         SELECT id         AS customer_id,
                created_at AS last_visit_time
         FROM `customer_info`
         UNION ALL
         SELECT customer_id,
                visit_time AS last_visit_time
         FROM `customer_visit`
         UNION ALL
         SELECT customer_id,
                sign_time AS last_visit_time
         FROM `d_crm_sales`.`sign_info`
     ) a
GROUP BY customer_id;



