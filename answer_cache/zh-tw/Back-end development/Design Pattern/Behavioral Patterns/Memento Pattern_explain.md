

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