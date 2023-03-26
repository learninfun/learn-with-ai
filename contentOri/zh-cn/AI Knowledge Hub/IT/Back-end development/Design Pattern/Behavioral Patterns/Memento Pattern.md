+++
title = "備忘錄模式"
weight = "6"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Memento Pattern的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Memento Pattern？
2. Memento Pattern的使用情境是什麼？
3. Memento Pattern中的三個角色分別是誰？
4. Memento Pattern與其他模式有什麼區別？
5. Memento Pattern的缺點是什麼？該如何解決？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Memento Pattern並舉例
{{< /ask_chatgpt >}}



Memento Pattern是一種行為型設計模式，它允許將物件的狀態儲存起來，以後可以回復到先前的狀態。它通常用在需要回復先前狀態的應用程式，例如文字編輯器、遊戲、以及檔案管理系統等。

Memento Pattern的主要元素包含三個：Originator、Memento、以及Caretaker。

- Originator：負責產生需要儲存的狀態，並允許存取、回復先前狀態。
- Memento：負責儲存原始物件的狀態。
- Caretaker：負責管理Memento的儲存以及回復，但它不應該存取或修改Memento。

下面舉個例子：假設有一個文字編輯器，我們想要實現撤銷（Undo）和重做（Redo）的功能：

```python
class TextEditor:
    def __init__(self):
        self.content = ""
        self.mementos = []

    def add_content(self, text):
        self.content += text

    def save_to_memento(self):
        self.mementos.append(TextEditorMemento(self.content))

    def restore_from_memento(self, memento):
        self.content = memento.get_saved_content()

class TextEditorMemento:
    def __init__(self, content):
        self.saved_content = content

    def get_saved_content(self):
        return self.saved_content

class TextEditorCaretaker:
    def __init__(self, editor):
        self.editor = editor
        self.undos = []

    def save(self):
        self.editor.save_to_memento()

    def undo(self):
        if self.mementos:
            undo_memento = self.mementos.pop()
            self.undos.append(undo_memento)
            self.editor.restore_from_memento(undo_memento)

    def redo(self):
        if self.undos:
            redo_memento = self.undos.pop()
            self.mementos.append(redo_memento)
            self.editor.restore_from_memento(redo_memento)
```

在這個例子中，TextEditor是Originator，它負責在字串尾巴添加新字元，並且可以保存狀態到Memento中。TextEditorMemento是Memento，它負責儲存TextEditor的狀態。TextEditorCaretaker是Caretaker，它負責管理Memento的儲存以及回復，在撤銷時，會儲存當前狀態到undos列表中，並且將上一個狀態從mementos列表中拿出來並回復到TextEditor中，在重做時，會將undos列表中的狀態拿出來，調用TextEditor的save_to_memento方法儲存到mementos列表中，再回復到TextEditor中。

這樣我們就可以實現文字編輯器的撤銷和重做功能了。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Memento Pattern的重點
{{< /ask_chatgpt >}}



1. Memento Pattern 主要是用來捕捉並儲存對象的內部狀態，並允許在未來恢復到之前的狀態。
2. Memento，即是備忘錄的意思，代表儲存對象的狀態的物件。
3. Originator，即是原始對象，指的是需要被儲存狀態的對象。
4. Caretaker，即是負責管理備忘錄的物件，可用於儲存、恢復 Originator 的狀態。
5. Memento Pattern 的優點是允許對像在不破壞其封裝性的前提下，儲存和恢復內部狀態，可以有效地解耦合作的對象，並簡化了設計和實現的複雜性。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Memento Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 如何利用Memento Pattern來保存遊戲的進度？
答案：創建一個Memento Class，包含玩家目前的遊戲進度。然後在Game Class中定義createMemento()和restoreFromMemento()方法，分別用於創建進度快照和從快照恢復進度。

2. 如何使用Memento Pattern來保存編輯器文檔的撤銷和重做操作？
答案：創建一個Memento Class，存儲文檔當前的狀態和操作歷史。在Editor Class中定義createMemento()和restoreFromMemento()方法，用於創建和恢復編輯器文檔的狀態。並且在Editor Class中實現撤銷和重做操作。

3. 如何利用Memento Pattern來保存狀態機的狀態？
答案：創建一個Memento Class，包含狀態機當前的狀態。在StateMachine Class中定義createMemento()和restoreFromMemento()方法，用於創建和恢復狀態機的狀態。並且在StateMachine Class中實現狀態轉換操作。

4. 如何使用Memento Pattern來保存文檔的多個版本？
答案：創建一個Memento Class，存儲文檔的狀態和版本號。在Document Class中定義createMemento()和restoreFromMemento()方法，用於創建和恢復文檔的狀態。在DocumentHistory Class中維護一個Memento List，存儲所有的文檔版本。

5. 如何使用Memento Pattern來保存繪圖軟件的繪圖步驟？
答案：創建一個Memento Class，存儲當前的繪圖狀態。在Painter Class中定義createMemento()和restoreFromMemento()方法，用於創建和恢復繪圖步驟。在History Class中維護一個Memento List，存儲所有的繪圖步驟，實現撤銷和重做操作。   

