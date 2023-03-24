

1. Radix Sort是一種非比較性排序算法。 

2. 排序的主要方式是通過比較數字的每一位，以確定它們的位置和順序。 

3. 數字的每一位都被視為一個鍵(key)，每一個鍵都要排序。 

4. Radix Sort可以使用LSD(Least Significant Digit)和MSD(Most Significant Digit)兩種方式進行排序。 

5. LSD表示最不顯著位排序，MSD表示最顯著位排序。 

6. MSD排序通常使用遞迴算法實現，LSD排序通常使用迭代算法實現。 

7. 數字的每個鍵可以使用桶排序(bucket sort)或計數排序(counting sort)進行排序。 

8. 實現Radix Sort，需要適當的預處理，以使得數據可以按照位數進行排序。 

9. Radix Sort的時間複雜度為O(nd)，其中n是排序數列的元素個數，d是數字的最大位數。 

10. Radix Sort通常用於排序長度固定的數字序列，如IP地址等。