

以下是MVVM架构重点：

1. Model：应用程式的资料模型，处理资料逻辑与存取资料的方法，并且不会与View或ViewModel互动。

2. View：使用者介面，可看成是呈现Model的内容，不直接向Model或任何物件发出指令。

3. ViewModel：连结View与Model的重要角色，担负着两方面的转换：资料转换(activity to model)及操作转换(view to model)。例如，当View有任何事件(按钮按下、输入资料等)并通知ViewModel时，ViewModel便可依照事件种类将请求转换为Model可处理的资讯，并返回处理结果到View上，以此牵引整个MVVM的工作流程。

4. 两者的绑定：绑定是让View和ViewModel之间传递资料的方法之一，它是一种方式，让变数状态随时同步，而不需撰写其他逻辑程式。例如，当model资料修改时，View会被通知资料变动，并自动更新自己所显示的内容。

5. 重用性：因为View和ViewModel解耦且彼此分离，所以显示模版和显示资讯都容易复用。

6. 测试性：虽然View和ViewModel都是需要人工测试，但是Model则是可以使用单元测试学派的测试方式，让整体架构更为健全。

7. 编写方式：在MVVM中，View主要是需要XAML语言进行拖拉式设计，也可直接刻画视界设计，ViewModel更多是以程式码的方式撰写，让封装、逻辑分层与非同步处理更为容易。