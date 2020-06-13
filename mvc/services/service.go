package services

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/cast"
	M "github.com/yuw-mvc/yuw/modules"
	"io/ioutil"
	"math"
	"net/http"
)

const(
	defaultPage int = 1
	defaultSize int = 10

	ToJSON string = "JSON"
	ToHTML string = "HTML"
)

type (
	PoT struct {
		Type string
		Code int
		HTML string
		Response *gin.H
	}

	Services struct {
		F *M.File
		U *M.Utils
		T *M.Token
	}
)

func New() *Services {
	return &Services {
		F: M.NewFile(),
		U: M.NewUtils(),
		T: M.NewToken(),
	}
}

func (srv *Services) Rd(rPoT *M.RdPoT) (r *redis.Client, err error) {
	rd, err := M.InstanceRd(rPoT)
	if err == nil {
		r = rd.Engine().R
	}

	return
}

func (srv *Services) Curl(reqURL string) (res []byte, err error) {
	client := &http.Client{}

	response, err := client.Get(reqURL)
	defer response.Body.Close()

	if err != nil {
		res = []byte("{}")
		return
	}

	if response.StatusCode == 200 {
		res, _ = ioutil.ReadAll(response.Body)
	} else {
		res = []byte("{}")
	}

	return
}

func (srv *Services) Paginator(page int, pageNums int, pageSize ...int) (paginator map[string]interface{}) {
	var (
		nums int = pageNums
		size int
		prePage int
		sufPage int
	)

	if len(pageSize) > 0 {
		size = pageSize[0]
	} else {
		size = defaultSize
	}

	var totalPage int = int(math.Ceil(float64(nums) / float64(size)))

	if page > totalPage {
		page = totalPage
	}

	if page <= 0 {
		page = 1
	}

	var pages []int

	switch {
	case page >= totalPage-5 && totalPage > 5:
		start := totalPage-5+1
		prePage = page-1
		sufPage = int(math.Min(float64(totalPage), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start+i
		}

	case page >= 3 && totalPage > 5:
		start := page-3+1
		pages = make([]int, 5)
		prePage = page-3
		for i, _ := range pages {
			pages[i] = start + i
		}

		prePage = page-1
		sufPage = page+1

	default:
		pages = make([]int, int(math.Min(5, float64(totalPage))))
		for i, _ := range pages {
			pages[i] = i + 1
		}

		prePage = int(math.Max(float64(1), float64(page-1)))
		sufPage = page+1
	}

	paginator = map[string]interface{}{
		"pages": pages,
		"total": totalPage,
		"cur_page": page,
		"pre_page": prePage,
		"suf_page": sufPage,
	}

	return
}

func (srv *Services) PaginatorParams(strPage string, strSize string) (page int, size int) {
	if strPage != "" {
		page = cast.ToInt(strPage)
	} else {
		page = defaultPage
	}

	if strSize != "" {
		size = cast.ToInt(strSize)
	} else {
		size = defaultSize
	}

	return
}