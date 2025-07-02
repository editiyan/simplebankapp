// Gunakan pointer untuk data dinamis (menggunakan heap memory)
// Jangan gunakan pointer untuk data yg konstan
// escape analysis
// go build -gcflags="-m"
// Array itu fix (copy by value) (menggunakan stack) (ga flexible karena ukuran tetap)
// slice itu dinamis (pass by references) (biasanya di heap, tergantung besar kecilnya ukuran data) (flexible dapat ditambah/dikurangin waktu runtime)

// go test -bench=. -benchmem
// go test -coverprofile=coverage.out .
// go tool cover -html=coverage.out -o coverage.html

// Running tool: C:\Program Files (x86)\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkByValue$ exercisealo/unittesting

// goos: windows
// goarch: 386
// pkg: exercisealo/unittesting
// cpu: AMD Ryzen 7 PRO 5850U with Radeon Graphics
// BenchmarkByValue-16    	1000000000	         0.2331 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	exercisealo/unittesting	1.001s

// Running tool: C:\Program Files (x86)\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkByPointer$ exercisealo/unittesting

// goos: windows
// goarch: 386
// pkg: exercisealo/unittesting
// cpu: AMD Ryzen 7 PRO 5850U with Radeon Graphics
// BenchmarkByPointer-16    	1000000000	         0.2443 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	exercisealo/unittesting	0.532s

// Running tool: C:\Program Files (x86)\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkByValue$ exercisealo/unittesting

// goos: windows
// goarch: 386
// pkg: exercisealo/unittesting
// cpu: AMD Ryzen 7 PRO 5850U with Radeon Graphics
// BenchmarkByValue-16    	1000000000	         0.2437 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	exercisealo/unittesting	1.222s

// Running tool: C:\Program Files (x86)\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkByPointer$ exercisealo/unittesting

// goos: windows
// goarch: 386
// pkg: exercisealo/unittesting
// cpu: AMD Ryzen 7 PRO 5850U with Radeon Graphics
// BenchmarkByPointer-16    	1000000000	         0.2659 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	exercisealo/unittesting	0.649s

SIMPLE BANKING APP
|
|-- go.mod
|-- main.go
|-- domain
|
|-- model
|
|-- impl
