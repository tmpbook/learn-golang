package main

import "fmt"
import "unicode/utf8"

func main() {
	s := "hello, world"
	fmt.Println(len(s))                     // "12"
	fmt.Println(s[0], s[7])                 // "104 119" 这里的数字是字节，下面的切片也是
	fmt.Println(s[0:5])                     // "hello" 和Python一样，不包含s[5]
	fmt.Println(s[:])                       // "hello, world" 和Python一样
	fmt.Println([]rune(s))                  // "[104 101 108 108 111 44 32 119 111 114 108 100]"
	fmt.Println([]byte(s))                  // "[104 101 108 108 111 44 32 119 111 114 108 100]"
	fmt.Println([]rune("世界"))               // "[19990 30028]"
	fmt.Println([]byte("世界"))               // "[228 184 150 231 149 140]"
	fmt.Println("世界")                       // "世界"
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c") // "世界"
	fmt.Println("\u4e16\u754c")             // "世界"
	fmt.Println("\U00004e16\U0000754c")     // "世界"

	s1 := "Hello, 世界"
	fmt.Println(len(s1))                    // "13"
	fmt.Println(utf8.RuneCountInString(s1)) // "9"
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// 0       H
	// 1       e
	// 2       l
	// 3       l
	// 4       o
	// 5       ,
	// 6
	// 7       世
	// 10      界

	for i, r := range s1 {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	// 0       'H'     72
	// 1       'e'     101
	// 2       'l'     108
	// 3       'l'     108
	// 4       'o'     111
	// 5       ','     44
	// 6       ' '     32
	// 7       '世'    19990
	// 10      '界'    30028
	n := 0
	for range s1 {
		n++
	}
	fmt.Println(n) // "9"

	s2 := "プログラム"
	r := []rune(s2)
	fmt.Printf("%x\n", r)        // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Println(string(r))       // "プログラム"
	fmt.Println(string(65))      // "A"
	fmt.Println(string(0x4eac))  // "京"
	fmt.Println(string(1234567)) // "�" 即:\uFFFD
}

/* 测试一个字符串是否是另一个的前缀 */
func HasProfix(s, profix string) bool {
	return len(s) >= len(profix) && s[:len(profix)] == profix
}

/* 后缀 */
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

/* 包含 */
func Contains(s, substr string) bool {
	// 包含
	for i := 0; i < len(s); i++ {
		if HasProfix(s[i:], substr) {
			return true
		}
	}
	return false
}
