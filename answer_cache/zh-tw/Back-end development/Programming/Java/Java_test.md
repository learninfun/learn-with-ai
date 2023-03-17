

Question 1:
實現一個方法，接收一個整數陣列，並返回一個由所有偶數組成的新陣列。

public static int[] getEvenArray(int[] array) {
    int count = 0;
    for (int i = 0; i < array.length; i++) {
        if (array[i] % 2 == 0) {
            count++;
        }
    }
    int[] newArray = new int[count];
    int index = 0;
    for (int i = 0; i < array.length; i++) {
        if (array[i] % 2 == 0) {
            newArray[index] = array[i];
            index++;
        }
    }
    return newArray;
}

Question 2:
請實現一個反轉字串的方法，將給定字串反轉，並返回反轉後的結果。

public static String reverse(String str) {
    char[] chars = str.toCharArray();
    int start = 0;
    int end = chars.length - 1;
    while (start < end) {
        char temp = chars[start];
        chars[start] = chars[end];
        chars[end] = temp;
        start++;
        end--;
    }
    return new String(chars);
}

Question 3:
實現一個方法，接受一個整數陣列和一個數字，將陣列中所有等於該數字的元素刪除，並返回一個新陣列。

public static int[] delete(int[] array, int num) {
    int count = 0;
    for (int i = 0; i < array.length; i++) {
        if (array[i] != num) {
            count++;
        }
    }
    int[] newArray = new int[count];
    int index = 0;
    for (int i = 0; i < array.length; i++) {
        if (array[i] != num) {
            newArray[index] = array[i];
            index++;
        }
    }
    return newArray;
}

Question 4:
給定一個只包含正整數的陣列 nums，從中找到一個數字，使得其出現了超過一半以上的次數並返回該數字。

public static int majorityElement(int[] nums) {
    int count = 1;
    int major = nums[0];
    for (int i = 1; i < nums.length; i++) {
        if (count == 0) {
            major = nums[i];
            count++;
        } else if (major == nums[i]) {
            count++;
        } else {
            count--;
        }
    }
    return major;
}

Question 5:
實現一個方法判斷一個字符串是否為回文串，回文串是指正和反讀都一樣的字符串。

public static boolean isPalindrome(String s) {
    s = s.toLowerCase();
    int start = 0;
    int end = s.length() - 1;
    while (start < end) {
        char c1 = s.charAt(start);
        char c2 = s.charAt(end);
        if (!Character.isLetterOrDigit(c1)) {
            start++;
        } else if (!Character.isLetterOrDigit(c2)) {
            end--;
        } else if (c1 != c2) {
            return false;
        } else {
            start++;
            end--;
        }
    }
    return true;
}