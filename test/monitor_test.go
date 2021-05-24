package test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type AutoGenerated struct {
	Product         string `json:"product"`
	LogUUID         string `json:"log_uuid"`
	Channel         string `json:"channel"`
	ClientIP        string `json:"client_ip"`
	ServerIP        string `json:"server_ip"`
	BeginTime       string `json:"begin_time"`
	EndTime         string `json:"end_time"`
	UsedTime        int    `json:"used_time"`
	DmsCode         int    `json:"dms_code"`
	DmsUsedTime     int    `json:"dms_used_time"`
	DmsSize         int    `json:"dms_size"`
	CandidateSize   int    `json:"candidate_size"`
	Number          int    `json:"number"`
	ProfileUsedTime int    `json:"profile_used_time"`
	ProfileCode     int    `json:"profile_code"`
	CacheSize       int    `json:"cache_size"`
	CostLevel       struct {
		Handlerfillscore struct {
			Total float64 `json:"total"`
		} `json:"handlerfillscore"`
		Handlergetcandidateset struct {
			ForceRecall                  float64 `json:"ForceRecall"`
			GetIsYoungPlanV1             float64 `json:"GetIsYoungPlanV1"`
			OnIndexRecallFinish          float64 `json:"OnIndexRecallFinish"`
			OperationPlatformRecall      float64 `json:"OperationPlatformRecall"`
			Roughing                     float64 `json:"Roughing"`
			SmartMarketing               float64 `json:"SmartMarketing"`
			TDebug                       float64 `json:"TDebug"`
			FillExecutableFilterDqAndDMP float64 `json:"fillExecutableFilterDqAndDMP"`
			FilldataTime                 float64 `json:"filldata_time"`
			GetInterestPointMap          float64 `json:"getInterestPointMap"`
			GetRecallInfoMap             float64 `json:"getRecallInfoMap"`
			IndexTime                    float64 `json:"index_time"`
			IntentTime                   float64 `json:"intent_time"`
			Pack                         float64 `json:"pack"`
			RecallByDMP                  float64 `json:"recallByDMP"`
			Total                        float64 `json:"total"`
		} `json:"handlergetcandidateset"`
		Handlerintervention struct {
			Total float64 `json:"total"`
		} `json:"handlerintervention"`
		Handlersort struct {
			Total float64 `json:"total"`
		} `json:"handlersort"`
		Handlerweightsum struct {
			Total float64 `json:"total"`
		} `json:"handlerweightsum"`
	} `json:"cost_level"`
	URL         string      `json:"url"`
	RecallTime  string      `json:"recall_time"`
	GetAddr     string      `json:"get_addr"`
	IntentData  string      `json:"intent_data"`
	Cost        interface{} `json:"cost"`
	UserCostMap struct {
		GetUsrAppInfo   int `json:"GetUsrAppInfo"`
		Adsplantime     int `json:"adsplantime"`
		Getuserinfotime int `json:"getuserinfotime"`
		Mapredistime    int `json:"mapredistime"`
		PortrayalTime   int `json:"portrayal_time"`
		Preparetime     int `json:"preparetime"`
		ProfileTime     int `json:"profile_time"`
		Ssbctime        int `json:"ssbctime"`
		UserseriTime    int `json:"userseri_time"`
	} `json:"user_cost_map"`
	ExtTime struct {
		AdjustAcceptReq         float64 `json:"AdjustAcceptReq"`
		BizRank                 float64 `json:"BizRank"`
		ExpInfo2BitMap          float64 `json:"ExpInfo2BitMap"`
		ExpInfoFromReq          float64 `json:"ExpInfoFromReq"`
		FillTestData            float64 `json:"FillTestData"`
		GetFlagValByMap         float64 `json:"GetFlagValByMap"`
		GetReqInfoFromURL       float64 `json:"GetReqInfoFromUrl"`
		GetUserInfo             float64 `json:"GetUserInfo"`
		InfoOut2Res             float64 `json:"InfoOut2Res"`
		Rank                    float64 `json:"Rank"`
		SpecificHandleProcedure int     `json:"SpecificHandleProcedure"`
		WriteData2Redis         float64 `json:"WriteData2Redis"`
		InitQueryContext        float64 `json:"initQueryContext"`
		RunHandlerProcedure     float64 `json:"runHandlerProcedure"`
		Tc1                     float64 `json:"tc1"`
		Tc2                     int     `json:"tc2"`
		Tc3                     float64 `json:"tc3"`
		Tc4                     int     `json:"tc4"`
	} `json:"ext_time"`
	Ddflag          bool        `json:"ddflag"`
	ExpLowctr       bool        `json:"exp_lowctr"`
	ItemKeys        []string    `json:"item_keys"`
	MissItems       interface{} `json:"miss_items"`
	RecallInfos     interface{} `json:"recall_infos"`
	SummaryUsedTime int         `json:"summary_used_time"`
	SummaryMsg      string      `json:"summary_msg"`
}

func TestJson(t *testing.T) {
	//
	a := 56
	v := strconv.Itoa(a)
	//s := string(a)
	fmt.Println("11--", v)

	f, err := os.Open("/Users/lizhen/Desktop/monitor.autohmpg2.log.20200817")
	if err != nil {
		fmt.Println("read file fail", err)
	}
	buf := bufio.NewReader(f)
	all := 0.0
	timeout := 0.0
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		lineList := strings.Split(line, "request log")
		if len(lineList) == 2 {
			monitorjson := strings.TrimSpace(lineList[1])
			autoGenerated := &AutoGenerated{}

			json.Unmarshal([]byte(monitorjson), autoGenerated)
			if autoGenerated.UsedTime > 300 && strings.Contains(autoGenerated.URL, "pid=100087") {
				timeout++
				//fmt.Println(autoGenerated.UsedTime)
			}
			if strings.Contains(autoGenerated.URL, "pid=100087") {
				all++
			}
		}
		if err != nil {
			if err == io.EOF {
				fmt.Printf("总计：【%v】，超时【%v】，占比【%v】", all, timeout, timeout/all*100)
				return
			}
		}
	}
}

func TestHandler(t *testing.T) {
	for i := 0; i < 100; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Nanosecond)
		fmt.Printf("%d\n", r.Intn(2))
	}
}