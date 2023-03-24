

1) 有兩個整數列表，找出兩個列表中共同出現的最小值。
   答案：SELECT MIN(shard1.val) FROM shard1 JOIN shard2 ON shard1.val=shard2.val;

2) 在給定的列表中找到前k個最大的元素(假設k小於列表的長度)。
   答案：SELECT val FROM shard ORDER BY val DESC LIMIT k;

3) 有一個包含重複元素的列表，找出存在超過n次的元素。
   答案：SELECT val FROM (SELECT val, count(*) as count FROM shard GROUP BY val) as val_count WHERE count>n;

4) 找出包含最多元素的重複子串(子串不需要連續)。
   答案：SELECT SUBSTRING(shard.str, start, length) as sub_str FROM (SELECT str, SUBSTRING_INDEX(SUBSTRING_INDEX(str, ',', numbers.n), ',', -1) as val, LENGTH(SUBSTRING_INDEX(SUBSTRING_INDEX(str, ',', numbers.n), ',', -1)) as length, LOCATE(SUBSTRING_INDEX(SUBSTRING_INDEX(str, ',', numbers.n), ',', -1), str)+1 as start FROM shard,(SELECT @row := @row + 1 as n FROM (select 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3) t1 CROSS JOIN (select 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3) t2 CROSS JOIN (select @row:=0) t3) numbers WHERE CHAR_LENGTH(str)-CHAR_LENGTH(REPLACE(str,',',''))>=numbers.n-1) as shards WHERE LENGTH(sub_str)>=2 GROUP BY sub_str ORDER BY COUNT(*) DESC LIMIT 1;

5) 找出前k個出現最多的單詞。
   答案：SELECT word, COUNT(*) as count FROM shard GROUP BY word ORDER BY count DESC LIMIT k;