package core

import pb "com/aliyun/serverless/nodeservice/proto"

var DefaultMaxUsedCount int64 = 1 //Container实例的默认最大连接数
var CollectionMaxCapacity = 1     //集合最大容量

//表示一个函数实例
//存放container信息
type Container struct {
	ContainerId string  //容器id
	TotalMem    int64   //容器总内存
	UsageMem    int64   //容器使用内存
	CpuUsagePct float64 //容器使用百分比

	FuncName         string //函数名字
	UseCount         int64  //使用数量
	ConcurrencyCount int64  //支持并发数量

	node *Node //所属node
}

func (c *Container) UpdateContainerStats(stats *pb.ContainerStats) {
	if stats == nil {
		return
	}

	c.TotalMem = stats.TotalMemoryInBytes
	c.UsageMem = stats.MemoryUsageInBytes
	c.CpuUsagePct = stats.CpuUsagePct
}


//
////向集合中添加一个Container
//func (cs *Collection) AddContainer(container *Container) {
//	cs.Containers = append(cs.Containers, container)
//}
//
////判断集合是否还有空间装container实例
//func (cs *Collection) Lack() bool {
//	return int64(len(cs.Containers)) < cs.Capacity
//}
//
////判断节点是否满足container的要求,和这个collection的使用人数
//func (cs *Collection) Satisfy(reqMem int64) (bool, int64) {
//	//判断集合中是否有容器
//	if len(cs.Containers) <= 0 {
//		return false, 0
//	}
//	//如果集合的使用人数，小于集合的最大使用人数，就数名满足需要
//	bool := cs.UsedCount < int64(len(cs.Containers))*cs.MaxUsedCount
//	return bool, cs.UsedCount
//}
//
////获取container
//func (cs *Collection) Acquire(reqMem int64) *Container {
//	//判断集合中是否有容器
//	if len(cs.Containers) <= 0 {
//		return nil
//	}
//
//	//获取一个使用人数最少的容器
//	container := cs.Containers[0]
//	for _, c := range cs.Containers {
//		if c.UsedCount < container.UsedCount {
//			container = c
//		}
//	}
//
//	cs.UsedCount++
//	container.UsedCount++
//	return container
//}
//
////归还container实例
//func (cs *Collection) Return(container *Container, actualUseMem int64) {
//	cs.UsedCount--
//	container.UsedCount--
//	if actualUseMem == 0 {
//		actualUseMem = 1 * 1024 * 1024
//	}
//	cs.MaxUsedCount = cs.MaxUsedMem / actualUseMem
//	container.MaxUsedCount = container.MaxUsedMem / actualUseMem
//	cs.UsedMem = actualUseMem
//	container.UsedMem = actualUseMem
//}
//
////将collection转换为字符串，打印日志的时候需要，
//func (cs *Collection) ToString() string {
//	info := "{" + cs.FuncName + ", " + str(cs.UsedCount) + "/" + str(int64(len(cs.Containers))*cs.MaxUsedCount) +
//		", " + str(cs.UsedMem/1024/1024) + "/" + str(cs.MaxUsedMem/1024/1024) + ", "
//
//	for _, c := range cs.Containers {
//		//info += "[" + c.FuncName + ", " + str(c.UsedCount) + ", " + str(c.UsedMem/1024/1024) + "], "
//		info += "[" + str(c.UsedCount) + "/" + str(c.MaxUsedCount) + ", " + str(c.UsedMem/1024/1024) + "], "
//	}
//	info += "}"
//	return info
//}
//
//func str(i int64) string {
//	return strconv.FormatInt(i, 10)
//}
