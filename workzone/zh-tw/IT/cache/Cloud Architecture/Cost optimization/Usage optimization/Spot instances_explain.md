Spot instances 是 Amazon EC2 的一種選項，它是一種基於競價的雲端運算。透過Spot instances，您可以向Amazon EC2（Elastic Compute Cloud）提出較低的出價，以獲取它們未被使用的計算容量。

當您使用Spot instances時，您需要指定最高可接受的出價。如果您的出價高於當前市場價格，那麼您的Spot instances就會被啟動。但是，如果您的出價低於當前市場價格，那麼您的Spot instances就會被暫停或中止。

Spot instances 的優點是價格比其他選項更低，因為您只需支付市場價格。另外一個優點是彈性性，您可以使用Spot instances來執行在市場價格波動幅度較大的應用。但是，使用Spot instances的缺點是，當您的出價低於市場價格時，您的Spot instances就有可能被中止，這可能會影響您的應用程序的穩定性。

以下是使用Spot instances的示例：

假設您正在開發一個需要大量運算資源的應用程序，而您需要在一定的時間內（例如幾個小時或幾天）完成。如果您使用On-Demand instances，那麼您需要支付相對較高的成本，這可能對您的預算造成不良影響。但是，如果您使用Spot instances，您可以根據當前市場價格提出低價出價，這樣就有可能取得所有所需計算容量，並在時間內完成開發工作，同時也更加符合預算要求。