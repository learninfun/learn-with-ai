

1. 有一個字元陣列，要求在其中找到第一個不是英文字母的字元。如果找不到，回傳-1。

Input: {'a', 'b', '+', 'd', 'E', 'f'}
Output: 2

2. 有一個整數陣列，每個數字都出現兩次，只有一個數字出現一次。找到這個數字。

Input: {1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}
Output: 6

3. 有一個正整數陣列，找到其中最小的缺失數字。也就是說，若陣列中有1、3、4，則回傳2。

Input: {1, 3, 5, 6, 7, 9}
Output: 2

4. 有一個整數陣列，找到其中第一個出現超過一次的數字。

Input: {1, 2, 3, 4, 5, 2, 3, 6, 7, 7}
Output: 2

5. 有一個字串陣列，找到其中最長的字串。

Input: {"cat", "dog", "elephant", "bird", "frog"}
Output: "elephant"

答案：

1. 

int findNonAlpha(char[] arr) {
    for(int i=0; i<arr.length; i++) {
        if(!Character.isLetter(arr[i])) {
            return i;
        }
    }
    return -1;
}

2. 

int findUnique(int[] arr) {
    int unique = arr[0];
    for(int i=1; i<arr.length; i++) {
        unique ^= arr[i];
    }
    return unique;
}

3.

int findMissing(int[] arr) {
    int n = arr.length;
    boolean[] present = new boolean[n+1];
    for(int i=0; i<n; i++) {
        if(arr[i] > 0 && arr[i] <= n) {
            present[arr[i]] = true;
        }
    }
    for(int i=1; i<=n; i++) {
        if(!present[i]) {
            return i;
        }
    }
    return n+1;
}

4.

int findDuplicate(int[] arr) {
    for(int i=0; i<arr.length; i++) {
        if(arr[Math.abs(arr[i])] >= 0) {
            arr[Math.abs(arr[i])] = -arr[Math.abs(arr[i])];
        } else {
            return Math.abs(arr[i]);
        }
    }
    return -1;
}

5.

String longestString(String[] arr) {
    int maxLength = 0;
    String longest = "";
    for(int i=0; i<arr.length; i++) {
        if(arr[i].length() > maxLength) {
            maxLength = arr[i].length();
            longest = arr[i];
        }
    }
    return longest;
}