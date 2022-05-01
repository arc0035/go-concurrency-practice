原题见这里[原题](https://leetcode.com/problems/print-in-order/)


要点：
- 先让三个goroutine用sleep功能，按指定睡眠时间sleep，达到我们手动控制调度的效果
- 使用管道，第1个routine 通知第2个routine，第2个routine再通知第三个，等等
