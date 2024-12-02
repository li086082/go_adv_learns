##### 文件操作
+ 获取文件描述符
    + os.Create() 
        + 如果文件已经存在，它将被截断
        + 如果该文件不存在，则以模式0666（在umask之前）创建
        + 如果成功，返回的File上的方法可以用于I/O；关联的文件描述符模式为O_RDWR
        + 如果有一个错误，它的类型将是*PathError

    + os.Open() Open打开指定文件以供阅读
        + 如果成功，返回文件上的方法可用于读取, 关联的文件描述符的模式为O_RDONLY
        + 如果有一个错误，它的类型将是*PathError

    + os.OpenFile() OpenFile是通用的公开调用；大多数用户会使用“打开”或“创建”
        + 如果文件不存在，并且传递了O_CREATE标志，则使用模式perm（在umask之前）创建该文件
        + 如果成功，返回的File上的方法可以用于I/O
        + 如果有一个错误，它的类型将是*PathError

+ 调整文件读/写位置
    + Seek(num, tag) - 在没有指定O_APPEND模式下生效
        + io.SeekStart: 向末尾偏移
        + io.SeekEnd: 向开头偏移
        + io.SeekCurren: 当前位置

+ 读文件
    + 按行读

    + 按字节读(推荐)
        + 注意每次追加都要按读取到的字节数进行追加, 不能整个buf都追加
    
+ 写文件
    + WriteString()
        + 写入指定字符串
    + WriteAt()
        + 按[]byte()字节切片写入指定位置, 可以和Seek()函数搭配


##### 权限说明
+ 每次打开这三个必须要选一个
    + O_RDONLY // 只读
	+ O_WRONLY // 只写
	+ O_RDWR   // 读写

+ 用作补充的权限
	+ O_APPEND // 不清空只追加内容
	+ O_CREAT  // 如何不存在则创建一个新的文件
	+ O_EXCL   // 和O_CREATE一起使用, 只有文件不存在时使用
	+ O_SYNC   // 同步打开I/O
	+ O_TRUNC  // 打开时截断可写文件