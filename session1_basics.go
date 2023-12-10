package main

import "fmt"

type human struct {
	height int
	weight float32
	name   string
}

func arrayAndSlices() {
	arr1 := [2]int{1, 2}
	fmt.Println(arr1)

	// sparse initialization of array
	arr2 := [5]string{"Hello", 3: "swecha"}
	fmt.Println(len(arr2[3]))

	// uncomment this to see error
	//fmt.Println(arr1 == arr2)

	// Slice, length is not fixed and will not give the error above
	humans := []human{{height: 10, weight: float32(166.3), name: "test 1"}, {height: 21, weight: float32(324), name: "hello`"}}
	fmt.Println(humans)

	// Backend of Slice is an Array and resizing of it will happen in background if required
	capitals := make([]string, 2, 5)
	// value of index 0 and 1 is default value ""
	fmt.Println("len:", len(capitals), "cap: ", cap(capitals))
	// pushed in index 2, 3, 4
	capitals = append(capitals, "Hyderabad", "Bengaluru", "Chennai")
	fmt.Println("len:", len(capitals), "cap: ", cap(capitals))

	// size increased of the slice, new array created with prev value and bigger size (size allocation done in batches [Buddy System])
	// old array will be deleted by the garbage collector
	capitals = append(capitals, "Mumbai")
	fmt.Println("len:", len(capitals), "cap: ", cap(capitals))
	capitals = append(capitals, "Delhi", "Bhubaneswar")
	fmt.Println("len:", len(capitals), "cap: ", cap(capitals))

	// Zero Length Slices
	// will apply 5 length array during program starting to avoid runtime resizing of slices in beginning
	names := make([]string, 0, 5)
	names = append(names, "Rajesh", "Ramesh", "Suresh")

	// Slicing of Slices (is a reference only)
	n := names[1:3]
	fmt.Println("n:", n, "Len:", len(n), "Cap:", cap(n))

	// slices are a pointer to actual array in backend, slicing of slices will make slices with name backend data
	n[1] = "RAJESH"
	fmt.Println("names:", names)
	fmt.Println("n:", n)

	// strings are immutable, inplace changes are not allowed
	s := "SWECHA"
	//s[1] = 'W'
	fmt.Println("s:", s, "Len:", len(s))
	fmt.Println("Accessing three characters:", s[0:3])
}

func maps() {
	var testmap map[string]string
	fmt.Println(testmap)

	fmt.Println(`testmap ["AP"]`, testmap["AP"])
	//testmap["AP"] = "Amaravati" // Error: this is a setter only

	//-------------------------
	capitalMap := map[string]string{
		"AP": "Amaravati",
		"TS": "Hyderabad",
	}
	fmt.Println("nmap", capitalMap, "Len:", len(capitalMap))
	fmt.Println(`capitalMap["AP"]`, capitalMap[`AP`])

	//-------------
	capitalMapMake := make(map[string]string)
	capitalMapMake[`AP`] = `Ama`

	var capital, ok = capitalMapMake[`AP`]

	fmt.Println(`capital`, capital, `ok`, ok)

}

func structs() {
	var h1 human
	h1.name = "a"
	h1.weight = 1
	h1.height = 1

	var h2 human
	h2.name = "a"
	h2.weight = 1
	h2.height = 1

	fmt.Println(`h1 == h2`, h1 == h2)
}

func controlStatement() {
	// n is defined only inside the block of if
	if n := 10; n%2 == 0 {
		fmt.Println(`n is even`)
	} else {
		fmt.Println(`n is odd`)
	}

	m := map[string]int{
		"one": 1,
	}

	// _ same as python
	if _, ok := m["two"]; !ok {
		fmt.Println(`two not found`)
	}

	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	names := []string{"John", "Paul", "George", "Ringo"}
	for i, name := range names {
		fmt.Println(i, name)
	}

	nmap := map[string]int{
		`John`: 1,
		`Paul`: 2,
		`Cow`:  3,
	}

	for key, value := range nmap {
		fmt.Println(key, value)
		switch size := len(key); size {
		case 1, 2, 3:
			fmt.Println(key, `is a small word`)
		case 4:
			keyLen := len(key)
			fmt.Println(key, `is exactly right length`, keyLen)
		default:
			fmt.Println(key, `is a big word`)
		}
	}

}

func add(x int, y int) int {
	return x + y
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, v := range vals {
		out = append(out, base+v)
	}

	return out
}

func functions() {
	fmt.Println(add(1, 3))
	fmt.Println(addTo(1, 2, 3, 4))

	foo := func() {
		fmt.Println("Anonymous function")
	}

	//foo() // blocking function
	defer foo() // will add this function execution after the execution of this function NOTE: will cause deadlock if something goes wrong
	fmt.Println("Hello, playground")

	// Go is always Call by Value always

}

func main() {
	fmt.Println("Hello World")

	a := 1000000000
	fmt.Println(a)

	//arrayAndSlices()
	//maps()
	//structs()
	//controlStatement()
	functions()
}
