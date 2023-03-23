+++
title = "备忘录模式"
weight = "6"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Memento Pattern的问题
{{< /ask_chatgpt >}}



1. 什麼是Memento Pattern？
2. Memento Pattern的使用情境是什麼？
3. Memento Pattern中的三個角色分別是誰？
4. Memento Pattern與其他模式有什麼區別？
5. Memento Pattern的缺點是什麼？該如何解決？   

## 说明知识
{{< ask_chatgpt >}}
说明Memento Pattern并举例
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

## 汇总重点
{{< ask_chatgpt >}}
条列Memento Pattern的重点
{{< /ask_chatgpt >}}



1. Memento Pattern 主要是用来捕捉并储存对象的内部状态，并允许在未来恢复到之前的状态。
2. Memento，即是备忘录的意思，代表储存对象的状态的物件。
3. Originator，即是原始对象，指的是需要被储存状态的对象。
4. Caretaker，即是负责管理备忘录的物件，可用于储存、恢复 Originator 的状态。
5. Memento Pattern 的优点是允许对像在不破坏其封装性的前提下，储存和恢复内部状态，可以有效地解耦合作的对象，并简化了设计和实现的复杂性。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Memento Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 如何利用Memento Pattern来保存游戏的进度？
答案：创建一个Memento Class，包含玩家目前的游戏进度。然后在Game Class中定义createMemento()和restoreFromMemento()方法，分别用于创建进度快照和从快照恢复进度。

2. 如何使用Memento Pattern来保存编辑器文档的撤销和重做操作？
答案：创建一个Memento Class，存储文档当前的状态和操作历史。在Editor Class中定义createMemento()和restoreFromMemento()方法，用于创建和恢复编辑器文档的状态。并且在Editor Class中实现撤销和重做操作。

3. 如何利用Memento Pattern来保存状态机的状态？
答案：创建一个Memento Class，包含状态机当前的状态。在StateMachine Class中定义createMemento()和restoreFromMemento()方法，用于创建和恢复状态机的状态。并且在StateMachine Class中实现状态转换操作。

4. 如何使用Memento Pattern来保存文档的多个版本？
答案：创建一个Memento Class，存储文档的状态和版本号。在Document Class中定义createMemento()和restoreFromMemento()方法，用于创建和恢复文档的状态。在DocumentHistory Class中维护一个Memento List，存储所有的文档版本。

5. 如何使用Memento Pattern来保存绘图软件的绘图步骤？
答案：创建一个Memento Class，存储当前的绘图状态。在Painter Class中定义createMemento()和restoreFromMemento()方法，用于创建和恢复绘图步骤。在History Class中维护一个Memento List，存储所有的绘图步骤，实现撤销和重做操作。   

