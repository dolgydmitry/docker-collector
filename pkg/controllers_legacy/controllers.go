package controllers

// "docker-collector/pkg/collectors"

// type MainControllerParams struct {
// 	Col    []collectors.CollectorExp
// 	Ticker *time.Ticker
// 	Quit   chan struct{}
// }

// func CollectMetric(param MainControllerParams) {
// 	for {
// 		select {
// 		case <-param.Ticker.C:
// 			for _, col := range param.Col {
// 				col.GetMetricsValue()

//				}
//			case <-param.Quit:
//				param.Ticker.Stop()
//				return
//			}
//		}
//	}
// var (
// 	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
// 		Name: "myapp_processed_ops_total",
// 		Help: "The total number of processed events",
// 	})
// )

// func CollectMetric(collector *collectors.DockerCli) {
// 	var mem runtime.MemStats
// 	go func() {
// 		for {
// 			runtime.ReadMemStats(&mem)
// 			log.Printf("alloc [%v] \t heapAlloc [%v] \n", mem.Alloc, mem.HeapAlloc)
// 			time.Sleep(5 * time.Second)
// 			collector.GetMetricsValue()
// 			// opsProcessed.Inc()
// 		}
// 	}()
// }
