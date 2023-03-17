

1. 問題：請寫一個方法，接收一個整數作為參數，若該參數小於0，則拋出一個IllegalArgumentException，否則返回該參數的平方。
答案：

public int getSquare(int num) {
    if (num < 0) { //檢查參數是否小於0
        throw new IllegalArgumentException("The input number should be greater than or equal to 0.");
    }
    return num * num;
}

2. 問題：請寫一個方法，接收一個字串作為參數，若該字串為空或null，則拋出一個NullPointerException，否則返回該字串的長度。
答案：

public int getLength(String str) {
    if (str == null || str.isEmpty()) { //檢查字串是否為空或null
        throw new NullPointerException("The input string cannot be null or empty.");
    }
    return str.length();
}

3. 問題：請寫一個方法，接收一個整數陣列和一個索引值作為參數，若索引值小於0或大於等於陣列長度，則拋出一個IndexOutOfBoundsException，否則返回該索引所對應的陣列元素。
答案：

public int getElement(int[] arr, int index) {
    if (index < 0 || index >= arr.length) { //檢查索引值是否越界
        throw new IndexOutOfBoundsException("The index is out of bounds.");
    }
    return arr[index];
}

4. 問題：請寫一個方法，接收一個字串和一個正整數作為參數，若該字串長度小於等於該正整數，則拋出一個RuntimeException，否則返回該字串的前n個字元。
答案：

public String getSubstring(String str, int n) {
    if (str.length() <= n) { //檢查字串長度是否足夠
        throw new RuntimeException("The length of the input string is not greater than " + n + ".");
    }
    return str.substring(0, n);
}

5. 問題：請寫一個方法，接收一個整數作為參數，若該整數為奇數，則拋出一個Checked Exception，否則返回該整數的平方。
答案：

public int getSquareEven(int num) throws Exception {
    if (num % 2 != 0) { //檢查整數是否為奇數
        throw new Exception("The input number should be even.");
    }
    return num * num;
}