package helper_crypto

/*

Android ID 是获取门槛最低的，不需要任何权限，64bit 的取值范围，唯一性算是很好的了。但是不足之处也很明显：

刷机、root、恢复出厂设置等会使得 Android ID 改变；Android 8.0之后，Android ID的规则发生了变化：

对于升级到8.0之前安装的应用，ANDROID_ID会保持不变。如果卸载后重新安装的话，ANDROID_ID将会改变。
对于安装在8.0系统的应用来说，ANDROID_ID根据应用签名和用户的不同而不同。ANDROID_ID的唯一决定于应用签名、用户和设备三者的组合。

两个规则导致的结果就是：

第一，如果用户安装APP设备是8.0以下，后来卸载了，升级到8.0之后又重装了应用，Android ID不一样；

第二，不同签名的APP，获取到的Android ID不一样。

其中第二点可能对于广告联盟之类的有所影响（如果彼此是用Android ID对比数据的话），所以Google文档中说“请使用Advertising ID”，不过大家都知道，Google的服务在国内用不了。


对Android ID做了约束，对隐私保护起到一定作用，并且用来做APP自己的活跃统计也还是没有问题的。

*/