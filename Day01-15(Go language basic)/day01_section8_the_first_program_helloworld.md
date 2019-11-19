# Day 01: Hello World

> @author: PVKDRAGON
> Copyright: Han Ru version chinese

## Cấu trúc của một dự án golang
- Phần này chưa rõ lắm

### 1.1 Thư mục gopath
- tạm pass

### 1.2 Viết chương trình đầu tiên
Bất kỳ môt ngôn ngữ lập trình nào cũng bắt đầu với chương trình "Hello, World", được đề cập lần đầu trong sách "The C Programming Language" xuất bản năm 1978. 
Đó là một câu chuyện hay về từ "Hello, World", mà tất cả các lập trình viên đều kỳ vọng máy tính một ngày nào đó sẽ có trí thông minh thực sự để trả lời lại cho người đã tạo ra
nó bằng một câu từ tận đáy lòng của nó như một đứa trẻ vừa chào đời.

1.Trong thư mục chứa code làm việc, tạo thư mục tên là src, sau đó tạo thư mục tên là hello trong thư mục vừa tạo, rồi tạo file helloworld.go, và gõ như sau:
```go
package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")
}
```
2.Chạy chương trình
Có vài cách để thực thi chương trình

Cách 1: Sử dụng lệnh go run

Bước 1: Mở terminal
Tuỳ hệ điều hành sẽ có cách mở khác nhau
Bước 2: Tìm đến nơi chứa file helloworld.go
Bước 3: gõ lệnh go run helloworld.go

Cách 2: Sử dụng go build command
Vẫn như 2 bước trên nhưng khác bước 3
Bước 3: gõ lệnh go build helloworld.go
Lệnh này sẽ tạo ra một chương trình tên là helloworld và có thể gọi thực thi bằng cách ./helloworld(Unix, Linux, Mac)
- Lệnh gop install với package chưa hiểu lắm

### 1.3 Giải thích các câu lệnh trong chương trình

#### 3.2.1 package

#### 3.2.2 import

#### 3.3.3 main
