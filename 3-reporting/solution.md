# SQL Task Solution

## 1
```
| questionnaire_name | participants |
| ------------------ | ------------ |
| Hair Colour Survey | Jenny, Alan |
| Arm survey | Alan, Jenny |
| Blood Quantity Survey | Alan, Jenny |
| Hair Qualitity Survey | Jenny, Alan |
| Who has legs? | Alan |
| Leg survey | Alan, Jenny |
| Blood Colour Survey | Jenny, Alan |
| Who has hands? | Alan |
| Blood Type Survey | Alan |
```
```
SELECT 
    `questionnaires`.`name` AS 'questionnaire_name',
    GROUP_CONCAT(`participants`.`name`) AS 'participants'
FROM 
    `scheduled_questionnaires`
LEFT JOIN 
    `questionnaires` ON `scheduled_questionnaires`.`questionnaire_id` = `questionnaires`.`id`
LEFT JOIN 
    `participants` ON `scheduled_questionnaires`.`participant_id` = `participants`.`id`
WHERE 
    `scheduled_questionnaires`.`scheduled_at` < CURRENT_TIMESTAMP()
AND 
    `scheduled_questionnaires`.`status` = 'pending'
GROUP BY 
    `scheduled_questionnaires`.`questionnaire_id`;
```

## 2
```
| calendar_day | scheduled_count | completed_count |
| ------------ | --------------- | --------------- |
| 2021-01-01 | 18 | 9 |
| 2021-01-02 | 1 | 0 |
| 2021-01-03 | 11 | 10 |
| 2021-01-04 | 1 | 1 |
| 2021-01-05 | 3 | 2 |
| 2021-01-06 | 1 | 1 |
| 2021-01-07 | 3 | 3 |
| 2021-01-08 | 4 | 2 |
| 2021-01-09 | 1 | 1 |
| 2021-01-10 | 3 | 3 |
| 2021-01-11 | 1 | 0 |
| 2021-01-13 | 1 | 1 |
| 2021-01-17 | 2 | 2 |
| 2021-01-24 | 2 | 2 |
| 2021-02-01 | 9 | 9 |
| 2021-02-02 | 1 | 1 |
| 2021-02-03 | 2 | 2 |
| 2021-02-04 | 1 | 1 |
| 2021-02-05 | 3 | 3 |
| 2021-02-06 | 2 | 2 |
| 2021-02-07 | 2 | 2 |
| 2021-02-08 | 3 | 3 |
| 2021-02-11 | 1 | 1 |
| 2021-02-15 | 2 | 2 |
| 2021-02-22 | 2 | 2 |
```
```
SELECT
	DATE(`scheduled_questionnaires`.`scheduled_at`) AS 'calendar_day',
	COUNT(`scheduled_questionnaires`.`id`) AS 'scheduled_count',
	COUNT(`questionnaire_results`.`id`) AS 'completed_count'
FROM
	`scheduled_questionnaires`
LEFT JOIN
	`questionnaire_results` 
ON
	`scheduled_questionnaires`.`id` = `questionnaire_results`.`questionnaire_schedule_id`
AND
	DATE(`scheduled_questionnaires`.`scheduled_at`) = DATE(`questionnaire_results`.`completed_at`)
GROUP BY
	DATE(`scheduled_questionnaires`.`scheduled_at`)
ORDER BY
	DATE(`scheduled_questionnaires`.`scheduled_at`) ASC
```

## 3)
```
| participant_name | questionnaire_name | questionnaire_timestamp |
| ---------------- | ------------------ | ----------------------- |
| Limmy | Who has legs? | 2021-01-01 12:00:00 |
| Limmy | Arm survey | 2021-01-03 12:00:00 |
| Limmy | Blood Type Survey | 2021-01-02 12:00:00 |
| Limmy | Hair Qualitity Survey | 2021-01-04 12:00:00 |
```
```
SELECT
	`participants`.`name` AS 'participant_name',
	`questionnaires`.`name` AS 'questionnaire_name',
	`questionnaire_results`.`completed_at` AS 'questionnaire_timestamp'
FROM
	`questionnaire_results`
LEFT JOIN
	`questionnaires` ON `questionnaire_results`.`questionnaire_id` = `questionnaires`.`id`
LEFT JOIN
	`participants` ON `questionnaire_results`.`participant_id` = `participants`.`id`
LEFT JOIN
	`scheduled_questionnaires` ON `participants`.`id` = `scheduled_questionnaires`.`participant_id`
WHERE
	`questionnaire_results`.`questionnaire_schedule_id` IS NULL
AND (
	`scheduled_questionnaires`.`status` = 'pending'
	OR `scheduled_questionnaires`.`status` IS NULL
)
```
