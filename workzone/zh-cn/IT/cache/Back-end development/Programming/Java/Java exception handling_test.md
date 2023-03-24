

1. 问题：请写一个方法，接收一个整数作为参数，若该参数小于0，则抛出一个IllegalArgumentException，否则返回该参数的平方。
答案：

public int getSquare(int num) {
    if (num < 0) { //检查参数是否小于0
        throw new IllegalArgumentException("The input number should be greater than or equal to 0.");
    }
    return num * num;
}

2. 问题：请写一个方法，接收一个字串作为参数，若该字串为空或null，则抛出一个NullPointerException，否则返回该字串的长度。
答案：

public int getLength(String str) {
    if (str == null || str.isEmpty()) { //检查字串是否为空或null
        throw new NullPointerException("The input string cannot be null or empty.");
    }
    return str.length();
}

3. 问题：请写一个方法，接收一个整数阵列和一个索引值作为参数，若索引值小于0或大于等于阵列长度，则抛出一个IndexOutOfBoundsException，否则返回该索引所对应的阵列元素。
答案：

public int getElement(int[] arr, int index) {
    if (index < 0 || index >= arr.length) { //检查索引值是否越界
        throw new IndexOutOfBoundsException("The index is out of bounds.");
    }
    return arr[index];
}

4. 问题：请写一个方法，接收一个字串和一个正整数作为参数，若该字串长度小于等于该正整数，则抛出一个RuntimeException，否则返回该字串的前n个字元。
答案：

public String getSubstring(String str, int n) {
    if (str.length() <= n) { //检查字串长度是否足够
        throw new RuntimeException("The length of the input string is not greater than " + n + ".");
    }
    return str.substring(0, n);
}

5. 问题：请写一个方法，接收一个整数作为参数，若该整数为奇数，则抛出一个Checked Exception，否则返回该整数的平方。
答案：

public int getSquareEven(int num) throws Exception {
    if (num % 2 != 0) { //检查整数是否为奇数
        throw new Exception("The input number should be even.");
    }
    return num * num;
}