原题见这里[原题](https://leetcode.com/problems/print-foobar-alternately/)


思路：
- 两个线程用管道交替打信号即可
- 也可以用sync.Mutex,抢到锁的线程如果发现是该由自己打印，则打印，反之退出，重新抢锁。
- 也可以用semaphore，每个线程打印完，release以通知对方线程打印（当然）