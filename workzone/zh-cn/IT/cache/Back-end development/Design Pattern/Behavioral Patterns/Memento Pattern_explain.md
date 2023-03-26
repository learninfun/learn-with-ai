

Memento Pattern是一种行为型设计模式，它允许将物件的状态储存起来，以后可以回复到先前的状态。它通常用在需要回复先前状态的应用程式，例如文字编辑器、游戏、以及档案管理系统等。

Memento Pattern的主要元素包含三个：Originator、Memento、以及Caretaker。

- Originator：负责产生需要储存的状态，并允许存取、回复先前状态。
- Memento：负责储存原始物件的状态。
- Caretaker：负责管理Memento的储存以及回复，但它不应该存取或修改Memento。

下面举个例子：假设有一个文字编辑器，我们想要实现撤销（Undo）和重做（Redo）的功能：

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

在这个例子中，TextEditor是Originator，它负责在字串尾巴添加新字元，并且可以保存状态到Memento中。TextEditorMemento是Memento，它负责储存TextEditor的状态。TextEditorCaretaker是Caretaker，它负责管理Memento的储存以及回复，在撤销时，会储存当前状态到undos列表中，并且将上一个状态从mementos列表中拿出来并回复到TextEditor中，在重做时，会将undos列表中的状态拿出来，调用TextEditor的save_to_memento方法储存到mementos列表中，再回复到TextEditor中。

这样我们就可以实现文字编辑器的撤销和重做功能了。