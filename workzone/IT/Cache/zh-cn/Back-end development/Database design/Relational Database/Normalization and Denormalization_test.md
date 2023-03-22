

1. 请问何谓Normalization(正规化)？列出三种正规化的形式。

答案：Normalization 是一种关系型资料库设计的技术，目的是为了减少重复资料并提高资料库设计的效率和可维护性。三种正规化形式分别是1NF、2NF 和 3NF。

2. 如果资料库中某张表格存在大量重复资料，应该考虑使用哪些正规化形式来解决这个问题？

答案：这表明这张表格不符合3NF限制，可以使用2NF 或 3NF去除重复资料。

3. 请列出两个常见的反规范化(Deormalization)的例子。

答案：1.在资料重复较多的查询上添加冗余栏位，以提高查询效率。2. 藉由合并多张表格，以减少查询中的表格连接数量和提高效率。

4. 什么是冗余栏位，给一个例子说明。

答案：冗余栏位就是多余的资料栏位。比如，在一个订单表中，有一个包含订单总价金额的栏位，如果再在该订单表中添加一个新的栏位，可以计算出每个产品单价的总价金额，那这个新栏位就是一个冗余栏位。

5. 当为了提高系统效率而进行反规范化时，也会带来一些问题，请列出至少两个反规范化可能带来的问题。

答案：1.增加了资料库的复杂度和体积，提高了系统维护成本。2. 如果不小心产生错误或栏位重复，可能会导致数据不一致。