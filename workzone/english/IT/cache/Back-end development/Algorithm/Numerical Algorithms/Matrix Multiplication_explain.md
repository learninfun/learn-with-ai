

Matrix multiplication is an arithmetic operation that takes two matrices as input and produces a new matrix as output. It involves iterating through the rows and columns of the two matrices and multiplying matching elements, then summing the resulting products.

For example, let's say we have two matrices:

```
A = [[1, 2],
     [3, 4]]

B = [[5, 6],
     [7, 8]]
```

To calculate the product of these matrices, we first need to determine the dimensions of the resulting matrix. In general, the number of columns in the first matrix must equal the number of rows in the second matrix to perform matrix multiplication.

In this case, A is a 2x2 matrix and B is also a 2x2 matrix. Since the number of columns in A (2) equals the number of rows in B (2), we can calculate their product.

To perform the calculation, we iterate through the rows of A and the columns of B, multiplying corresponding elements and summing the results. This gives us the following:

```
C = [[1*5 + 2*7, 1*6 + 2*8],
     [3*5 + 4*7, 3*6 + 4*8]]

C = [[19, 22],
     [43, 50]]
```

The resulting matrix C is also a 2x2 matrix, with each element calculated as the sum of products of the corresponding elements in A and B.