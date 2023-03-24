

1. 請說明什麼是Iterator Pattern以及其特點？

答：Iterator Pattern是一種軟體設計模式，用於提供一種方式來順序訪問一個物件的元素，而不暴露該物件的實踐細節。Iterator Pattern的主要特點是：

- 提供統一的接口去訪問容器內的元素。
- 使容器與其內部元素的訪問被分離，從而增加容器的可重複使用性。
- 在不改變容器結構的前提下，支援對內部元素的多種訪問方式。

2. Iterator Pattern中Iterator的三個基本操作是什麼？

答：Iterator Pattern中，Iterator通常需要支援三個基本操作：

- next()：返回下一個元素。
- hasNext()：檢查是否還有下一個元素。
- remove()：從容器中移除目前所指向的元素。

3. 請說明如何實現一個外部迭代器？

答：外部迭代器是指它自身代表了一個遍歷過程，從而能夠讓使用者隨意跳過、刪除與替換集合中的元素。實現一個外部迭代器通常需要以下步驟：

- 創建一個迭代器類，實現基本的Iterator介面。
- 在迭代器類中加入指向具體容器的遊標（例如，ArrayList中的index）。
- 實現基本的next()、hasNext()、remove()等方法。在remove()方法中，需要先取出目前選中的元素，同時刪除該元素，並將遊標指向下一個元素。

4. 在Java中，如何實現一個尋找vector中的最大元素的例子？

答：

```
import java.util.*;

public class VectorIteratorExample {
   public static void main(String[] args) {
      Vector<Integer> nums = new Vector<Integer>(Arrays.asList(1, 2, 3, 4, 5));
      Iterator<Integer> itr = nums.iterator();
      int max = nums.get(0);
      while (itr.hasNext()) {
         int num = itr.next();
         if (num > max) {
            max = num;
         }
      }
      System.out.println("The maximum element in the vector is: " + max);
   }
}
```

5. 如何在C++中實現一個迭代器類？

答：

```
#include <iostream> 

using namespace std; 

template <class T> 

class MyIterator 
{ 
   T* p; 

public: 
   MyIterator(T* x = NULL) { p = x; } 

   T& operator*() { return *p; } 

   T* operator->() { return p; } 

   MyIterator operator++() { p++; return *this; } 

   MyIterator operator++(int) { MyIterator tmp(*this); operator++(); return tmp; } 

   bool operator==(const MyIterator& rhs) const { return p == rhs.p; } 

   bool operator!=(const MyIterator& rhs) const { return p != rhs.p; } 
}; 

int main() 
{ 
   int arr[] = { 1, 2, 3, 4, 5 }; 

   MyIterator<int> i; 
   for (i = arr; i != (arr + 5); i++) { 
      cout << *i << " "; 
   } 

   return 0; 
}
```