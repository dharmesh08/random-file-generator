package main

import (
	"flag"
	"github.com/Pallinder/go-randomdata"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

//$env:GOPATH= "C:\Users\dpatel\golang"
//$Env:GOPATH=`godep path`
// go run main.go -dir=C:\Users\dpatel\Desktop\Temp\randgenerator -howManyFiles=1000 -total_lines=10000 -remove-files=false
// godep go run main.go -dir=C:\Users\dpatel\Desktop\Temp\randgenerator -howManyFiles=100 -total-lines=10000 -remove-files=true -wait-before-remove=10s
func main() {

	cwdir, _ := os.Getwd()
	dir := flag.String("dir", cwdir, "Directory")
	howManyFiles := flag.Int("how-many-files", 10, "how-many-files default 10")
	totalLines := flag.Int("total-lines", 10, "totallines default 1000")
	removeFiles := flag.Bool("remove-files", true, "removeFile is false")
	WaitBeforeRemove := flag.Duration("wait-before-remove", 1*time.Second, "wait-before-remove 1s")

	flag.Parse()
	log.Println("dir:", *dir)
	log.Println("how-many-files:", *howManyFiles)
	log.Println("total_lines:", *totalLines)
	log.Println("remove-files:", *removeFiles)

	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		log.Println("dir is require")
		os.Exit(0)
	}
	hostname := ""
	if hs, herr := os.Hostname(); herr == nil {
		hostname = hs
	}
	var wg sync.WaitGroup
	files := []string{}

	var mux sync.Mutex

	for i := 0; i < *howManyFiles; i++ {
		wg.Add(1)
		go func(fwg *sync.WaitGroup) {

			defer fwg.Done()
			file := path.Join(*dir, hostname+"_"+strconv.Itoa(random(100, 5000))+"_"+randomdata.FirstName(randomdata.Male)+randomdata.StringNumberExt(2, "-", 4)+".txt")
			mux.Lock()
			rand.Seed(time.Now().Unix())
			files = append(files, file)
			mux.Unlock()
			f, err := os.Create(file)
			if err != nil {
				log.Println(err)
				//break
				return
			}
			defer f.Close()
			for j := 0; j < *totalLines; j++ {
				data := randomdata.FirstName(randomdata.Male) + "," + randomdata.LastName() + "," + randomdata.Address() + "," + randomdata.City() + "," + randomdata.State(randomdata.Large) + "," + randomdata.PostalCode("US") + "," + randomdata.Email()
				f.WriteString(data)
			}
			f.Sync()

			return

		}(&wg)

	}

	log.Printf("Waiting ")
	wg.Wait()
	log.Printf("Done Waiting ")

	log.Printf("sleeping %f seconds", WaitBeforeRemove.Seconds())

	time.Sleep(*WaitBeforeRemove)
	if *removeFiles {
		log.Printf("Removing Files ")
		for _, file := range files {
			os.Remove(file)
		}
	}

	os.Exit(0)

}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
