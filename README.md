# routree

```go	
r := Router{}
for i, pattern := range map[int]string{
	0: ".*",
	1: "7495123.*",
	2: "7(49[5|9]).......*",
	3: "7(49[5|9])......*",
	4: "7(49[5|9]).......",
	5: "1(72[0-3|4-7|8|9],5[5|7].)......*",
	6: "7495#123.*",
} {
	patterns, err := ParsePattern(pattern)
	if err != nil {
		return
	}
	r.Add(patterns, fmt.Sprintf("%d:%q", i, pattern))
}
for _, number := range []string{
	"74951234567",
	"74981234567",
	"74991234567",
	"749512345678",
	"7495#1234567",
	"17211234567",
	"15555555555",
} {
	phone, err := ParsePhone(number)
	if err != nil {
		return
	}
	fmt.Printf("%-12s -> %v\n", number, r.Match(phone))
}
// Output:
// 74951234567  -> [1:"7495123.*" 4:"7(49[5|9])......." 2:"7(49[5|9]).......*" 3:"7(49[5|9])......*" 0:".*"]
// 74981234567  -> [0:".*"]
// 74991234567  -> [4:"7(49[5|9])......." 2:"7(49[5|9]).......*" 3:"7(49[5|9])......*" 0:".*"]
// 749512345678 -> [1:"7495123.*" 2:"7(49[5|9]).......*" 3:"7(49[5|9])......*" 0:".*"]
// 7495#1234567 -> [6:"7495#123.*"]
// 17211234567  -> [5:"1(72[0-3|4-7|8|9],5[5|7].)......*" 0:".*"]
// 15555555555  -> [5:"1(72[0-3|4-7|8|9],5[5|7].)......*" 0:".*"]

```

---

```
goos: linux
goarch: amd64
pkg: github.com/pshvedko/routree
cpu: 11th Gen Intel(R) Core(TM) i7-11700F @ 2.50GHz
BenchmarkRouter_Add-16      	 9059997	       197.0 ns/op
BenchmarkRouter_Match-16    	  577052	      1740 ns/op
PASS
goos: freebsd
goarch: amd64
pkg: github.com/pshvedko/routree
cpu: Intel(R) Core(TM) i3-6100T CPU @ 3.20GHz
BenchmarkRouter_Add-4     	 5739992	       197.3 ns/op
BenchmarkRouter_Match-4   	 1700040	       701.3 ns/op
PASS
```
