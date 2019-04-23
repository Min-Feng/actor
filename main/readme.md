# 模擬 erlang 的 訊息傳遞方式

實現兩個特性

1. 對特定actor對象，發送訊息
2. 當父級actor結束時，子級actor也會一起結束

可用main.go做不同的情況的測試，執行如下圖

**Files:**

1. [main.go](main.go)
2. [package actor](../actor.go)

![](result.png)