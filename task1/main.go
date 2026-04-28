package main

import (
	"fmt"
	"sync"
	"time"
)

// ฟังก์ชัน Worker
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n ", id, job)
		time.Sleep(1 * time.Second) // จำลองการทำงาน
		wg.Done()
	}
}

// ฟังก์ชันหลักตามโจทย์
func RunWorkers(numWorkers int, numJobs int) {
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// สร้าง workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, &wg)
	}

	// เพิ่มจำนวนงานเข้า WaitGroup //chanel
	for j := 1; j <= numJobs; j++ {
		wg.Add(1)
		jobs <- j
	}

	close(jobs) // ปิด channel เมื่อส่งงานครบ

	wg.Wait() // รอจนงานทั้งหมดเสร็จ
	fmt.Println("All jobs completed")
}

func main() {
	RunWorkers(3, 10) // ตัวอย่าง: 3 workers, 10 jobs
}