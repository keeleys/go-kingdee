```go
package main

import (
	"log"

    "github.com/keeleys/go-kingdee/erp"
)

type fhQueryModel struct {
	FID             int
	FBillNo         string
	FDocumentStatus string
	FSaleDeptId     string `json:"fSaleDeptId.FNumber"`
	FSaleDeptName   string `json:"fSaleDeptId.FName"`
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	erp.InitBaseUrl("https://xxx.xxx.com/K3Cloud/")
	erp.InitSession("xxx", "xxx@", "xxx")
}
func main() {
	result := erp.PostSelectHandler("SAL_SaleOrder", fhQueryModel{}, func(query *erp.ErpQueryModel) {
		query.FilterString = "FBillNo = 'XSDD202212200001'"
	})
	log.Printf("返回结果:%+v\n", result)
}

```