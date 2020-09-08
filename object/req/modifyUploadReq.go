/*
@Time   : 2020/9/8 13:04
@Author : Haichao Wang
*/
package req

type ModifyUpload struct {
	MnId       int64 `json:"mn_id"`
	ModifyType int   `json:"modify_type"`
}
