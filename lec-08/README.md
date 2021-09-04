# So sánh tốc độ mysql với elasticsearch
| Số row | MySQL |  ES  |
| ------ |  ----- | ----- |
| 150 | 2.399505ms | 22.300538ms es |
| 150 | 2.33152ms | 18.06978ms |
| 1 triệu 048 nghìn 576 | 5.962471045s | 29.943037ms |
| 1 triệu 048 nghìn 576 | 6.90573595s | 45.033811ms |
- MySQL chỉ trả về 1 row tìm thấy đầu tiên, ES trả về một mảng các kết quả, sắp xếp theo thang điểm
- Mới một số tìm kiếm bị thiếu từ, ES vẫn trả về đúng kết quả, còn MySQL không thể tìm thấy
- Ngoài ra thời gian insert Mysql cũng khá lâu (~ 30phút) còn với ES chỉ mất 3m30.924384323s


# Repository's info

|     Fullname    |    Class     |
|-----------------|--------------|
| Nguyen Duc Hieu |   TFS - 03   |

![class_diagram](https://res.cloudinary.com/duchieu/image/upload/v1630716224/test/itbjzexbfdel3qsy6rhp.png)
![people_diagram](https://res.cloudinary.com/duchieu/image/upload/v1630716198/test/zqyxsuq9q0bq8cxytlfy.png)

### directory's tree

```
├── learn_transaction
│   ├── transaction_gorm.go
│   ├── transaction_manual.go
│   └── transfer.sql
├── people
│   ├── controllers
│   │   ├── calendarContro.go
│   │   ├── classContro.go
│   │   ├── homecontro.go
│   │   ├── studentContro.go
│   │   └── teacherContro.go
│   ├── database
│   │   └── connectDB.go
│   ├── main.go
│   └── routes
│       └── routes.go
├── people.png
├── people.sql
├── README.md
└── search_engine
    ├── controllers
    │   ├── HomeContro.go
    │   └── SearchContro.go
    ├── docker-compose.yml
    ├── main.go
    ├── models
    │   └── elastic_search.go
    ├── routes
    │   └── routes.go
    └── views
        └── home.html
```
