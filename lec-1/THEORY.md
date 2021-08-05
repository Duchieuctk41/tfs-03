## Package, variable and function...
- Golang là ngôn ngữ biên dịch
### Package
- Chia nhỏ chức năng của chương trình ra, giúp dễ quản lý
- Tạo local package phải trùng tên với folder mẹ
- Có một số package có sẵn như:
- [fmt](https://pkg.go.dev/fmt) - triển khai I/O như printing, scanning
- [math/rand](https://pkg.go.dev/math/rand) - tạo số random
- [os](https://pkg.go.dev/os) - mở file, đọc file, nhận giá trị console
- [io/ioutil](https://pkg.go.dev/io/ioutil) - triển khai I/O như read file, write file
- [bufio]() - triển khai buffered I/O
- [strings](https://pkg.go.dev/strings) - thao tác với chuỗi (compare, replace, ToUpper...)
- [net/http](https://pkg.go.dev/net/http) - triển khai http client - server
- [strconv](https://pkg.go.dev/strconv) - chuyển đổi: chuỗi <=> số
- [encoding/json](https://pkg.go.dev/encoding/json) - decode và encode định dạng json (Marshal - Unmarshal)

### Variable
- byte, int, float32, float64, string, bool, uint
- Khai báo dữ liệu
 ```sh
var variableName dataType = initialValue (bỏ dataType nó vẫn hiểu)
variableName := initialValue
var (
variableName1 dataType = initialValue1
variableName2 dataType = initialValue2, 
)
const cho hằng số
```
### Function
- Viết hoa chữ đầu là public, để camelCase là private
- [Closures](https://tour.golang.org/moretypes/25) - hàm tham chiếu đến các giá trị ở bên ngoài nó, lưu lại giá trị ngay cả khi func ngoài đã kết thúc

### Map
- Kiểu dử liệu gồm key/value (giống object trong javascript)
- Map là dạng reference type (pass by reference)
- Khi khởi tạo mà ko có make >> giá trị trả về nil >> panic
```sh
mapName := make(map[keyType]valueType)
```
### Struct
- Kiểu dữ liệu tự định nghĩa
```sh
type nameStruc struct {
	variableName1 dataType1
	variableName2  dataType2
}
```

### Interface
- Xác định hành động của đối tượng
  ![](https://res.cloudinary.com/duchieu/image/upload/v1628161531/test/pmt45zpfxsmsqycoej43.png)
- nil của interfaces không có giá trị và loại cụ thể, Gọi 1 phương thức trong nil interfaces mà chưa được định nghĩa sẽ sinh ra lỗi.
- Khai báo 1 empty rỗng và có thể implement mọi thứ vào trong thằng này

### Pointer
- Trỏ đến địa chỉ vùng nhớ của giá trị
```sh
var variableName1 *nameType = &variableName2
```

### Method
- [method](https://techmaster.vn/posts/35044/series-golang-co-ban-phan-17-methods-phuong-thuc) giống classes trong OOP
- Được khai báo trong hàm gồm receiver và receiverType
```sh
func (t Type) methodName(parameter list) {
    //something
}
```

### Array và Slice
- [Array]() để lưu giá trị các biến, các biến sẽ được lưu trong vị trí nhớ liền kề nhau. vị trí nhớ nhỏ nhất là biến đầu tiên
- Để lưu 1 biến trong mảng mất 4 byte = 4 ô nhớ = 32bit
- Khi lưu trữ 1 mảng các ngôn ngữ thường lưu với hệ số gia tăng từ 1.5 - 2, vì nếu fix cứng hệ số là 1, sẽ ko thể thêm phần tử mới vào trong mảng đc. phải clone nó ra một mảng mới với sức chứa lớn hơn. Khi này độ phức tạp sẽ là O(n)
- Ngược lại, khi fix với hệ số 1.5 - 2 thì khi cần thêm phần tử mới sẽ rất đơn giản, nhanh hơn. độ phức tạp sẽ là O(1)
- Trong go ko cần fix cứng bộ nhớ array
- [Slice]() Mảng tham chiếu tới 1 mảng
```sh
num1 := [3]int{78, 79 ,80} // khai báo Array
num2 := num1[:] // khai báo slice
```

### Array vs Slice => Pass by value / Pass by Reference

![](https://res.cloudinary.com/duchieu/image/upload/v1628163307/test/ictoic1bchayf2vqll2t.gif)

###  Concurrency và Parallelism
- [Concurrency](https://www.youtube.com/watch?v=KmJ-Phn49jA): Giống với async await trong javascript (xử lý bất đồng bộ)
- [Parallelism](): Xử lý đồng bộ

### Unit test
- [log.fatal với panic](https://stackoverflow.com/questions/35996966/what-are-the-differences-in-outcome-for-panic-vs-log-fatalln-in-golang)  khác nhau chỗ nào
- [benchmark](https://pkg.go.dev/testing): test chạy nhanh hay chậm, tốn bao nhiêu resource


