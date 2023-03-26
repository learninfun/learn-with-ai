

1. 有一個背包可容納重量為W的物品。現在有n件物品，每件物品的重量為wi，價值為vi。請問能夠放入背包的物品中，總價值最大是多少？
答案：經典的0/1背包問題

2. 有一個背包可容納重量為W的物品。現在有n件物品，每件物品的重量為wi，價值為vi，但是每件物品只能使用一定次數。請問能夠放入背包的物品中，總價值最大是多少？
答案：有限制的01背包問題

3. 有一個背包可容納重量為W的物品。現在有n件物品，每件物品的重量為wi，價值為vi，但是被放入背包後每件物品的價值會隨時間推移而下降，下降的速率與時間有關，每單位時間價值下降d元。請問如何放置物品，使得背包內物品總價值的下降速率最小？
答案：附帶下降價值的背包問題

4. 有一個背包可容納重量為W的物品。現在有n件物品，每件物品有兩個參數，重量wi和價值vi，且重量和價值均為整數。現在要求選擇若干件物品放入背包中，使得滿足以下條件：
- 價值的總和最大化
- 重量的總和小於等於W
- 剛好選取k件物品
請問如何選擇物品以滿足以上條件？
答案：K背包問題

5. 有一個背包可容納重量為W的物品。現在有n件物品，每件物品都擁有多個屬性，例如體積、重量、價值等等。現在要求選擇若干件物品擺放至背包中，使得滿足以下條件：
- 滿足所有物品的多個屬性限制
- 價值的總和最大化
請問如何選擇物品以滿足以上條件？
答案：多維背包問題