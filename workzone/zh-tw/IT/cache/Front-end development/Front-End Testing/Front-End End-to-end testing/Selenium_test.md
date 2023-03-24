

1. 使用Selenium WebDriver如何選擇iframe内的元素?
答案：
可以使用switch_to.frame()方法切換到iframe, 然後再使用find_element()方法查找元素

2. 如何在Selenium WebDriver中實現窗口間的切換?
答案：
可以使用switch_to.window()方法切換到指定窗口, 或使用window_handles屬性得到所有窗口句柄, 再切換到目標窗口的句柄

3. 如何在Selenium WebDriver中實現鼠標和鍵盤操作?
答案：
可以使用ActionChains類中的move_to_element(), click(), send_keys()等方法實現鼠標和鍵盤操作

4. 如何實現Selenium WebDriver的等待機制?
答案：
可以使用WebDriverWait類中的until()或until_not()方法實現顯式等待, 或使用implicitly_wait()方法實現隱式等待

5. 如何實現Selenium WebDriver的截圖功能?
答案：
可以使用get_screenshot_as_file(), get_screenshot_as_base64()或get_screenshot_as_png()方法實現截圖功能。具體方法取決於你需要哪種格式的圖片存儲到哪個位置。