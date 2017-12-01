package main

//v3
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var Usage = func() {
	fmt.Println("Usage: ./dns zone_file_name")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func RemoveDuplicates(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func master() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	file := args[1]
	f, err := ioutil.ReadFile(file)
	CheckErr(err)
	str := string(f)
	array := strings.Split(str, " ")

	//获取主机区域
	filename := filepath.Base(file)
	zone := strings.Split(filename, `.zone`)

	var hostlist []string //获取完整主机名
	var service []string  //业务
	var project []string  //项目
	for _, name := range array {
		//获取主机
		//hre := regexp.MustCompile(`(?m:^[a-z]+)[\d]+[a-z]+.*$`)
		hre := regexp.MustCompile(`([[:alpha:]])+([\d])+([a-z])?.*$`)
		host := hre.FindAllString(name, -1)
		if host != nil {
			hostname := host[0] + "." + zone[0]
			hostlist = append(hostlist, hostname)
		}
		//获取业务
		sre := regexp.MustCompile(`(?m:^[a-z]+)[\d]+[a-z]?`)
		regex := sre.FindAllString(name, -1)
		if regex != nil {
			rmint := regexp.MustCompile(`([a-z]+)([\d]+)(.*)`)
			result := rmint.ReplaceAllString(regex[0], "$1")
			service = append(service, result)
		}
		//获取项目
		preg := regexp.MustCompile(`[a-z]+$`)
		item := preg.FindAllString(name, -1)
		if item != nil {
			project = append(project, item[0])
		}
	}

	sort.Sort(sort.StringSlice(service))
	ser := RemoveDuplicates(service)
	z := make(map[string]map[string][]string)
	sort.Sort(sort.StringSlice(project))
	pro := RemoveDuplicates(project)
	sort.Sort(sort.StringSlice(hostlist))
	//reb := regexp.MustCompile(`([a-z]+)([\d]+)([a-z]?)`)
	reb := regexp.MustCompile(`([a-z]+)([\d]+)([a-z]?)`)
	for _, p := range pro {
		m := make(map[string][]string)
		for _, hosts := range hostlist {
			for _, s := range ser {
				prefix := strings.Split(hosts, `.`)
				hprefix := reb.ReplaceAllString(prefix[0], "$1")
				if prefix[1] == p && hprefix == s {
					m[s] = append(m[s], hosts)
					z[p] = m
				}
			}
		}
	}
	hostjson, _ := json.Marshal(z)
	fmt.Println(string(hostjson))
}

func main() {
	master()
}
