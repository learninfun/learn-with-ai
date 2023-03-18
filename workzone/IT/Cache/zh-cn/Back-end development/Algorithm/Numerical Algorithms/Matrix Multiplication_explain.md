

矩陣乘法 (Matrix Multiplication) 是一種數學運算，用於將兩個矩陣進行運算得到一個新的矩陣。矩陣乘法的步驟如下：

1. 確認兩個矩陣相乘時是否符合維度要求，即第一個矩陣的列數必須等於第二個矩陣的行數。
2. 從第一個矩陣中選取一行，從第二個矩陣中選取一列，將兩者對應位置的數字相乘，得到新的矩陣中的一個元素。
3. 重複步驟2，直到得到新矩陣的所有元素。

例如，若有以下兩個矩陣：

矩陣A:
[ 1, 2 ]
[ 3, 4 ]

矩陣B:
[ 5, 6 ]
[ 7, 8 ]

則 A 與 B 的乘積可表示為：
[ 1*5 + 2*7, 1*6 + 2*8 ]
[ 3*5 + 4*7, 3*6 + 4*8 ]

即：
[ 19, 22 ]
[ 43, 50 ]