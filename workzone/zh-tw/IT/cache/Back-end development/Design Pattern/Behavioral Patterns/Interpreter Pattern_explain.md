

Interpreter Pattern 是一種設計模式，它定義了一種方式來解釋語言的語法規則，用於創建一個解釋器，以解析和執行語言的表達式。這種模式的核心思想是將一個特定的語言解釋器分離出來，使得它可以獨立於整個語言系統的其他部分進行開發和維護。Interpreter Pattern 主要由以下兩個組成部分組成：

1.文法（Grammar）：定義了語言中可以使用的合法語法結構。
2.解釋器（Interpreter）：用於對文法表示的語言語句進行解釋和執行。

例如，我們可以創建一個簡單的解釋器來解析和執行 SQL 語句，如下所示：

1. 文法：定義了 SQL 語句的語法結構和語義。

```
<SQL Query> ::= SELECT <Columns> FROM <Table> WHERE <Conditions>
<Columns> ::= <Column Name> | <Column Name>, <Columns>
<Conditions> ::= <Field> <Operator> <Value>
<Table> ::= <Table Name>
<Column Name> ::= <String>
<Operator> ::= = | > | < | >= | <= | <>
<Value> ::= <String>
<Field> ::= <String>
```

2. 解釋器：用於解釋和執行 SQL 查詢語句並生成結果。

```python
class SQLInterpreter:
    def __init__(self, query: str):
        self.tokens = query.split()
        self.columns = []
        self.table = ""
        self.conditions = {}

    def interpret(self):
        if "select" in self.tokens:
            for i in range(self.tokens.index("select") + 1, self.tokens.index("from")):
                if self.tokens[i] != ",":
                    self.columns.append(self.tokens[i])
        if "from" in self.tokens:
            self.table = self.tokens[self.tokens.index("from") + 1]
        if "where" in self.tokens:
            i = self.tokens.index("where") + 1
            while i < len(self.tokens):
                self.conditions[self.tokens[i]] = self.tokens[i + 2]
                i += 3
        # 執行 SQL 查詢語句並生成結果
        ...

query = "select id, name from customers where age > 18"
interpreter = SQLInterpreter(query)
interpreter.interpret()
```

在以上示例中，我們創建了一個簡單的 SQL 解釋器，它可以解釋和執行 SQL 查詢語句，並返回查詢結果。使用這種模式可以將解釋器與整個系統的其他部分解耦，從而更好地管理和維護系統。