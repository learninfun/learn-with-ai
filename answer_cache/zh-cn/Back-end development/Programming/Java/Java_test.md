

Question 1:
实现一个方法，接收一个整数阵列，并返回一个由所有偶数组成的新阵列。

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
请实现一个反转字串的方法，将给定字串反转，并返回反转后的结果。

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
实现一个方法，接受一个整数阵列和一个数字，将阵列中所有等于该数字的元素删除，并返回一个新阵列。

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
给定一个只包含正整数的阵列 nums，从中找到一个数字，使得其出现了超过一半以上的次数并返回该数字。

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
实现一个方法判断一个字符串是否为回文串，回文串是指正和反读都一样的字符串。

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