package bagique

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type CommonApi struct{}

type Rate struct {
	Rate float64 `json:"rate"`
}

type Cur struct {
	Price string `json:"price"`
}

type Result struct {
	Cur Cur `json:"cur"`
}

type Api struct {
	Result Result `json:"Result"`
}

// GetRate 获取美元汇率
// @Tags Common
// @Summary 获取美元汇率
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query {object} true "获取美元汇率"
// @Success 200 {object} response.Response{data=Rate,msg=string} "查询成功"
// @Router /common/getRate [get]
func (commonApi *CommonApi) GetRate(c *gin.Context) {
	url := "https://finance.pae.baidu.com/vapi/v1/getquotation?group=huilv_minute&need_reverse_real=1&code=USDCNY&finClientType=pc"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		global.GVA_LOG.Error("调用失败,创建请求错误", zap.Error(err))
		response.FailWithMessage("调用失败,创建请求错误"+err.Error(), c)
		return
	}
	req.Header.Add("Cookie", "PSTM=1702862660; BIDUPSID=ACF181E7AD7E8CDD789A3CFDC1BB3555; BAIDUID=61BC7DE0829BF6F8882B4CFE64F799AF:SL=0:NR=10:FG=1; BDUSS=2l6YXdTZXRTfm82dHBvaERKMU4yQVV1N0I3c3pBd3pUeX54eGFMfnF2ekRESzltRVFBQUFBJCQAAAAAAAAAAAEAAAAEuKW0TXVzdGFuZ196ZXJvAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMN~h2bDf4dmV; BDUSS_BFESS=2l6YXdTZXRTfm82dHBvaERKMU4yQVV1N0I3c3pBd3pUeX54eGFMfnF2ekRESzltRVFBQUFBJCQAAAAAAAAAAAEAAAAEuKW0TXVzdGFuZ196ZXJvAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMN~h2bDf4dmV; MCITY=-236%3A; H_PS_PSSID=60275_60599_60724_60360_60797; H_WISE_SIDS=60275_60599_60724_60360_60797; H_WISE_SIDS_BFESS=60275_60599_60724_60360_60797; BDORZ=FFFB88E999055A3F8A630C64834BD6D0; BDSFRCVID=H48OJexroG308n3tgRO-owDNbgKKz7RTDYLEgD-k666_WH8VFvD5EG0PtEfL4Zu-oxFFogKKymOTHrAF_2uxOjjg8UtVJeC6EG0Ptf8g0M5; H_BDCLCKID_SF=tbkD_I8KfI_3fP36q4Q2MJ8eKp8X5-CsLTvi2hcH0KLKV66d2lbC5--1h4Ty-M3wKH5vMbRdbfb1MRjvhpKM0MANMUOm2hRr3aTa3h5TtUJFeCnTDMRNXtK4XfbyKMniMCj9-pn-0lQrh459XP68bTkA5bjZKxtq3mkjbPbDfn02JKKuDTLBejoWjH0sh4c8KC_X3JjV5PK_Hn7zenb8Qx4pbq7H2M-eBIJBLtbj3hnrDP3eX45Zy-IbWM6n0pcH3Rb7KpR82fo0eIo33lrWQ5FkQN3T-nJGa6crLn_K2qo-Dn3oynrVXp0nK4Qly5jtMgOBBJ0yQ4b4OR5JjxonDh83bG7MJUutfJCfVC_atCK3fP36q4Q_MRLfbeT22-usKDvi2hcHMPoosI8xe-I55--w5fny-M3wKDvia-F5LfbUoqRHbU8VMT5BjRnUBTvp3jcR_l5TtUJMsJQsjh6hqt4bQ4jyKMniMCj9-pn-0lQrh459XP68bTkA5bjZKxtq3mkjbPbDfn028DKuDjtBD6QLjHRf-b-XKCoBLRvHHJOoDDv1j63cy4LbKxnxJ58OaJ7J0MoVXb7bshrPbURv26-w3-OkWUQM3PjW_DLKXM31bMQD5x8hQfbQ0hO4KfAJyjILW-nI3J7JOpvwhfnxybDFQRPH-Rv92DQMVU52QqcqEIQHQT3m5-5bbN3ut6IDJbKH_CDaf-3bfTrph46hMPQH-UnLqMkJfgOZ04n-ah05hKnkDnOJ0xRDhpJRWtjMW23E_Mom3UTdsJ3eqt4-j6Km3PAOLf5q3nT4KKJxbp5t8KbL5Dco2MKzhUJiBbcLBan7QqjIXKohJh7FM4tW3J0ZyxomtfQxtNRJ0DnjtpChbRO4-TFMDj3Bjx5; BDSFRCVID_BFESS=H48OJexroG308n3tgRO-owDNbgKKz7RTDYLEgD-k666_WH8VFvD5EG0PtEfL4Zu-oxFFogKKymOTHrAF_2uxOjjg8UtVJeC6EG0Ptf8g0M5; H_BDCLCKID_SF_BFESS=tbkD_I8KfI_3fP36q4Q2MJ8eKp8X5-CsLTvi2hcH0KLKV66d2lbC5--1h4Ty-M3wKH5vMbRdbfb1MRjvhpKM0MANMUOm2hRr3aTa3h5TtUJFeCnTDMRNXtK4XfbyKMniMCj9-pn-0lQrh459XP68bTkA5bjZKxtq3mkjbPbDfn02JKKuDTLBejoWjH0sh4c8KC_X3JjV5PK_Hn7zenb8Qx4pbq7H2M-eBIJBLtbj3hnrDP3eX45Zy-IbWM6n0pcH3Rb7KpR82fo0eIo33lrWQ5FkQN3T-nJGa6crLn_K2qo-Dn3oynrVXp0nK4Qly5jtMgOBBJ0yQ4b4OR5JjxonDh83bG7MJUutfJCfVC_atCK3fP36q4Q_MRLfbeT22-usKDvi2hcHMPoosI8xe-I55--w5fny-M3wKDvia-F5LfbUoqRHbU8VMT5BjRnUBTvp3jcR_l5TtUJMsJQsjh6hqt4bQ4jyKMniMCj9-pn-0lQrh459XP68bTkA5bjZKxtq3mkjbPbDfn028DKuDjtBD6QLjHRf-b-XKCoBLRvHHJOoDDv1j63cy4LbKxnxJ58OaJ7J0MoVXb7bshrPbURv26-w3-OkWUQM3PjW_DLKXM31bMQD5x8hQfbQ0hO4KfAJyjILW-nI3J7JOpvwhfnxybDFQRPH-Rv92DQMVU52QqcqEIQHQT3m5-5bbN3ut6IDJbKH_CDaf-3bfTrph46hMPQH-UnLqMkJfgOZ04n-ah05hKnkDnOJ0xRDhpJRWtjMW23E_Mom3UTdsJ3eqt4-j6Km3PAOLf5q3nT4KKJxbp5t8KbL5Dco2MKzhUJiBbcLBan7QqjIXKohJh7FM4tW3J0ZyxomtfQxtNRJ0DnjtpChbRO4-TFMDj3Bjx5; BA_HECTOR=21052la1a12h052lak2k058h0fh4sb1jem2td1v; BAIDUID_BFESS=61BC7DE0829BF6F8882B4CFE64F799AF:SL=0:NR=10:FG=1; ZFY=JhxDdDi9hr:BW1oqteW9LH4UMRdyq1FDbP9izUUvYKGw:C; delPer=0; PSINO=2; ab_sr=1.0.1_MGE0ZWE1YTA0ZDYyZjdhMTUzZmMxNDY4YTU0YzAzOTNmMTY0ZmFhNTgyMWE4M2IwYzRmNWJjNWRhYzZmNzcyMjU1Y2Y4YmYxNTgwYWY3MzZhNzE5OGM5NmM1NDE5NWNmMDg4OTlkOWNkNWYyMTFlNGEzZDk3NGNkODJhZWE0Y2JjZmUxYTVkOTdiZDMyODM1ODBlZTRjYzFiN2Q5YjM1MQ==")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("调用失败,发起请求错误", zap.Error(err))
		response.FailWithMessage("调用失败,发起请求错误"+err.Error(), c)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			global.GVA_LOG.Error("调用失败,创建io读取错误", zap.Error(err))
			response.FailWithMessage("调用失败,创建io读取错误"+err.Error(), c)
			return
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		global.GVA_LOG.Error("调用失败,访问错误", zap.Error(err))
		response.FailWithMessage("调用失败,访问错误"+err.Error(), c)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("调用失败,ioutil读取错误", zap.Error(err))
		response.FailWithMessage("调用失败,ioutil读取错误"+err.Error(), c)
		return
	}

	var api Api
	if err := json.Unmarshal(body, &api); err != nil {
		global.GVA_LOG.Error("调用失败,json解析失败", zap.Error(err))
		response.FailWithMessage("调用失败,json解析失败"+err.Error(), c)
		return
	}
	price, err := strconv.ParseFloat(api.Result.Cur.Price, 64)
	if err != nil {
		global.GVA_LOG.Error("调用失败,价格转换错误", zap.Error(err))
		response.FailWithMessage("调用失败,价格转换错误"+err.Error(), c)
		return
	}
	rate := Rate{Rate: price}
	response.OkWithData(rate, c)
}
