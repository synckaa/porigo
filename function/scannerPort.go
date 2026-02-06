package function

import (
	"fmt"
	"net"
	"porigo/model"
	"strings"
	"sync"
	"time"
)

func ScannerPort(myport []int, mytarget string, download bool) {
	target := strings.ToLower(mytarget)
	timeout := 2 * time.Second
	start := time.Now()
	ports := make(chan int, 100)
	resultData := make(chan model.Result)
	done := make(chan bool)

	var wg sync.WaitGroup
	var finalResult []model.Result

	go func() {
		for res := range resultData {
			finalResult = append(finalResult, res)
		}
		done <- true
	}()

	for w := 0; w < 100; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range ports {
				address := fmt.Sprintf("%s:%d", target, p)
				conn, err := net.DialTimeout("tcp", address, timeout)
				if err != nil {
					continue
				}
				serviceName := getServiceName(p)
				resultData <- model.Result{p, serviceName}
				fmt.Printf("Port [%d] is Open -> %s\n", p, serviceName)
				err = conn.Close()
				if err != nil {
					return
				}

			}
		}()
	}

	for i := myport[0]; i <= myport[1]; i++ {
		ports <- i
	}

	close(ports)
	wg.Wait()
	close(resultData)
	<-done
	timing := time.Since(start)
	if download {
		saveToJson(finalResult)
	}
	fmt.Printf("scanning finished in %s\n", timing.String())

}
