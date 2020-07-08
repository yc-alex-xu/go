# CAP theorem
The CAP Theorem states that, in a distributed system (a collection of interconnected nodes that share data.), you can only have two out of the following three guarantees across a write/read pair: Consistency, Availability, and Partition Tolerance - one of them must be sacrificed. However, as you will see below, you don't have as many options here as you might think.

![overview](images/CAP-overview.png)

states that it is impossible for a distributed data store to simultaneously provide more than two out of the following three guarantees:

* Consistency: Every read receives the most recent write or an error
* Availability: Every request receives a (non-error) response – without the guarantee that it contains the most recent write
* Partition tolerance: The system continues to operate despite an arbitrary number of messages being dropped (or delayed) by the network between nodes

In particular, the CAP theorem implies that in the presence of a network partition, one has to choose between consistency and availability. Note that consistency as defined in the CAP theorem is quite different from the consistency guaranteed in ACID database transactions.(Consistency ensures that a transaction can only bring the database from one valid state to another, in ACID)

##  CP - Consistency/Partition Tolerance -
Wait for a response from the partitioned node which could result in a timeout error. The system can also choose to return an error, depending on the scenario you desire. Choose Consistency over Availability when your business requirements dictate atomic reads and writes.

## AP - Availability/Partition Tolerance - 
Return the most recent version of the data you have, which could be stale. This system state will also accept writes that can be processed later when the partition is resolved. Choose Availability over Consistency when your business requirements allow for some flexibility around when the data in the system synchronizes. Availability is also a compelling option when the system needs to continue to function in spite of external errors (shopping carts, etc.)


# consensus algorithm Vs distributed lock

https://en.wikipedia.org/wiki/Distributed_lock_manager

Etcd is open-source software, developed at CoreOS under the Apache License.It can be used to perform distributed locks as well.

lock 提供的高可用性肯定比直接提供一致性协议的库要来的低, 相当于通过lock 实现强一致和直接通过一致性协议提供强一致的业务逻辑

* 有些场景很难改成通过一致性协议的场景
* 分布式锁的使用方式和之前单机是使用Lock 的方式一样的直观, 对使用人员的要求无成本
* 有时候用户只有少量的机器, 比如只有两个机器, 就无法通过一致性协议提供强一致, 但是通过外部的锁服务可以

在分布式锁里面有两个要求 (alex:跟deadlock avoidance算法非常类似)
* safety :任意时刻只能有一个 node 去获得这个lock
* liveness: 这个lock 是在某一个时刻是会终止的

其实这两点是互相矛盾的, 存在这个问题的本质原因是因为在分布式系统里面我们无法判断一个节点
* 真的挂掉还是
* 一会这个网络又会恢复.

为了解决这个问题.那么大体的实现就是给在节点lock 的时候, 对于这个lock 操作有一个lease的限制, 如果在这个租约的时间内, 这个节点没有来续租, 那么就认为这个操作是超时的, 一般情况为了实现的可靠性保证, 会在这个租约失效前就提前续租, 比如租约的时间是 10s, 我在0s 的时候就获得这个lock, 那么我在6s 的时候就必须去做这个续租操作, 如果没有执行成功的话, 那么我就认为你这个lease 失效了, 其他的节点可以在6s 时刻就获得这个lock, 但是只能在10s以后提供服务, 新的节点的租约时间是10s~20s. 那么从6s~10s 这段时间即使新节点获得了lock, 但是也无法提供服务的, 这个是典型的CAP 场景里面系统availability 换取consistenty 的例子.