

Linear Queue是指一種基於先進先出(FIFO)原則的資料結構。如同一列在銀行排隊的方式，最先進入排隊的人最先被處理，後進入排隊的人就要等候前面的人處理完畢以後才能進行下一步操作。

在Linear Queue中，資料是線性排列的，並且在做入隊(Enqueue)和出隊(Dequeue)操作時，資料只能在頭尾兩端進行。一般來說，Linear Queue是用Array或Linked List實現的。

以下是Linear Queue的範例：

當一列人在銀行排隊時，最先進入排隊的人(ID: 001)，會成為第一個進入Queue的元素。之後，第二個人(ID: 002)進入排隊，成為Enqueue的元素。當第一個人(ID: 001)完成作業後，他成為Dequeue的元素，並且由第二個人(ID: 002)取代成為Head元素。

就像這個例子，當資料依照順序進入Queue後，會依照先進先出的原則完成操作。所有在Queue中等候的元素會在適當時間被處理完畢。