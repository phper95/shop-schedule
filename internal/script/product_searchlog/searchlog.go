package product_searchlog

import (
	"fmt"
	"gitee.com/phper95/pkg/timeutil"
	"go.mongodb.org/mongo-driver/bson"
	"shop-schedule/global"
	"time"
)

type Searchlog struct {
	UserID     int64  `json:"userid" bson:"userid"`
	Keyword    string `json:"keyword" bson:"keyword"`
	New        *int   `json:"new" bson:"new"`
	Sales      string `json:"sales" bson:"sales"`
	Price      string `json:"price" bson:"price"`
	PageNum    int    `json:"page_num" bson:"page_num"`
	PageSize   int    `json:"page_size" bson:"page_size"`
	CreateTime int64  `json:"create_time" bson:"create_time"`
}

var productSearchedUserCount map[int64]int

func AnalysisSearchLog() {
	global.LOG.Warn("AnalysisSearchLog start ....")
	tables := getTables()
	productSearchedUserCount = make(map[int64]int, 0)
	for _, table := range tables {
		count, err := global.Mongo.EstimatedDocumentCount(global.ProductSearchLogDbName, table)
		if err != nil {
			global.LOG.Error("EstimatedDocumentCount error", err)
		} else {
			global.LOG.Warnf("table %s ; count %d", table, count)
		}
		global.Mongo.FindUseCursor(global.ProductSearchLogDbName, table, 1000, bson.D{}, &Searchlog{}, cursorCallback)
	}

	global.LOG.Warnf("AnalysisSearchLog finished , productSearchedUserCount %+v", productSearchedUserCount)
}

func getTables() (tables []string) {
	days := 10
	for i := 0; i < days; i++ {
		t := time.Now().AddDate(0, 0, -i)
		//if i > 0 {
		//	t = t.AddDate(0, 0, -i)
		//}
		tables = append(tables, fmt.Sprintf(global.ProductSearchLogCollectionNamePrefix, timeutil.YMDLayoutInt64(t)))
	}
	global.LOG.Warn("getTables", tables)
	return tables
}

func cursorCallback(res interface{}, err error) {
	if err != nil {
		global.LOG.Error("cursorCallback error", err, res)
		return
	} else {
		log := res.(*Searchlog)
		global.LOG.Warn(log)
		// cursor有超时时间，这里不宜做耗时操作
		productSearchedUserCount[log.UserID] += 1
	}
}
