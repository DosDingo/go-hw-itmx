# go-hw-itmx
### golang itmx test

## Instruction
if you have a problem with go run main.go with **CGO_ENABLED=0**

**Try this**
```
go env -w CGO_ENABLED=1
```
if you have a problem with **gcc**

https://jmeubank.github.io/tdm-gcc/ 


download and install this **(FOR WINDOW)**

## Get Start

```
go run main.go
```

มี scripts สำหรับใช้เพื่อเพิ่มข้อมูลใน customers เผื่อไว้ลองเล่นกันครับ

สามารถนำคำสั่งนี้ไปรันใน terminal ได้เลยครับ
```
go run .\scripts\create_customers.go
```

**ส่วนตอนลองเล่นก็สามารถใช้ postman หรือ curl ได้ตามสะดวกเลยนะครับ**

ในส่วนของ code coverage สามารถใช้คำสั่งตามด้านล่างได้เลยครับ

```
go test ./... -coverprofile=coverage.out
```

