package convert

import (
	"strconv"
	"strings"

	"localhost/lgMini/lagouSpider/downloader"
	"localhost/lgMini/lagouSpider/pipeline"
	"time"
)

func ToPipelineJobs(dJobs []downloader.Result) []pipeline.BytedanceJob {
	var pJobs []pipeline.BytedanceJob
	for _, v := range dJobs {
		longitude, _ := strconv.ParseFloat(v.Longitude, 64)
		latitude, _ := strconv.ParseFloat(v.Latitude, 64)
		pJobs = append(pJobs, pipeline.BytedanceJob{
			City:     v.City,
			District: v.District,

			CompanyShortName: v.CompanyShortName,
			CompanyFullName:  v.CompanyFullName,
			CompanyLabelList: strings.Join(v.CompanyLabelList, ","),
			CompanySize:      v.CompanySize,
			FinanceStage:     v.FinanceStage,

			PositionName:      v.PositionName,
			PositionId:        v.PositionId,
			PositionLables:    strings.Join(v.PositionLables, ","),
			PositionAdvantage: v.PositionAdvantage,
			WorkYear:          v.WorkYear,
			Education:         v.Education,
			Salary:            v.Salary,

			IndustryField:  v.IndustryField,
			IndustryLables: strings.Join(v.IndustryLables, ","),

			Longitude:  longitude,
			Latitude:   latitude,
			Linestaion: v.Linestaion,

			CreateTime: MustDateToUnix(v.CreateTime),
			AddTime:    time.Now().Unix(),
		})
	}

	return pJobs
}
