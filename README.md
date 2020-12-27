# Zeinab Local Library
## To run server:
>go run main.go

## To run client:
- **First way**
    - > cd client
    - > go run client.go --help
- **Second way**
    - > go run client/client.go --help
---
## Some client args:

**To add book**

*`-ab BookData spreated by space`*
> go run client.go -ab 1 CPL 2/3/2000 OmarSaad Sci AlexLib English
 

**To save book info**

*`-ub BookID NewData`*
> go run client.go -ub 1 OOD 2/3/2000 OmarSaad Sci AlexLib English

**To search book by id**

*`-Sb BookID`*
> go run client.go -Sb 1

**To search book by title**

*`-sb BookTitle`*
> go run client.go -sb CPL

**To sort book by title**

*`-otb`*
> go run client.go -otb

**To search book by date**

*`-odb`*
> go run client.go -odb

---

**To add reader**

*`-ar ReaderData spreated by space`*
> go run client.go -ar 1 OmarSaad male 1/1/2000 170 75

**To remove reader**

*`-rb BookId`*
> go run client.go -rb 1

**To search reader by id**

*`-Sr`*
> go run client.go -Sr 1

**To search reader by name**

*`-sr`*
> go run client.go -sr OmarSaad