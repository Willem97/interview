# Docker 面试题
<!-- TOC -->

- [Docker 面试题](#docker-面试题)
    - [1.](#1)
    - [2.](#2)
    - [3.Docker资源隔离](#3docker资源隔离)

<!-- /TOC -->
---

## 1.

## 2.

## 3.Docker资源隔离

Docker的隔离性主要运用`Namespace 技术，其中资源隔离还是用了cgroup，可以限制被namespace隔离的资源。

传统上Linux中的PID是唯一且独立的，在正常情况下，用户不会看见重复的PID。然而在Docker采用了Namespace，从而令相同的PID可于不同的Namespace中独立存在。如，A Container 之中PID=1是A程序，而B Container之中的PID=1同样可以是A程序。虽然Docker可透过Namespace的方式分隔出看似是独立的空间，然而Linux内核（Kernel）却不能Namespace，所以即使有多个Container，所有的system call其实都是通过主机的内核处理，这便为Docker留下了不可否认的安全问题。

`Namespace（命名空间）` 可以隔离哪些

- 文件系统需要是被隔离的
- 网络也是需要被隔离的
- 进程间的通信也要被隔离
- 针对权限，用户和用户组也需要隔离
- 进程内的PID也需要与宿主机中的PID进行隔离
- 容器也要有自己的主机名
- 有了以上的隔离，我们认为一个容器可以与宿主机和其他容器是隔离开的。

`Namespace` 隔离缺点（隔离不彻底）

- 容器是运行在宿主机的特殊进程，不同容器之间实用的还是同一个宿主机的操作系统内核
- Linux内核中，有些资源和对象不能对namespace化，如时间，修改容器时间，宿主机的时间也会被更改。

参考链接

> [1] https://zhuanlan.zhihu.com/p/106009423

> [2] https://www.cnblogs.com/--smile/p/11810027.html