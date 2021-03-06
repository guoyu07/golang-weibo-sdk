package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesUpdateBizString : 发布一条微博信息
// http://open.weibo.com/wiki/C/2/statuses/update/biz
func StatusesUpdateBizString(param apiParam.StatusesUpdateBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.StatusesUpdateBizURL, URLParam)
	return apiHelper.APIPost(apiURL.StatusesUpdateBizURL, URLParam)
}

// StatusesUpdateBizObject : 以对象的方式返回
func StatusesUpdateBizObject(param apiParam.StatusesUpdateBizParam, accessToken string) (*apiResp.StatusesUpdateBizResp, *apiResp.ErrorResp) {
	body, error := StatusesUpdateBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesUpdateBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesUpdateBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
