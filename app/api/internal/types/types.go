// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type CustomerList struct {
	LimitNum int64 `form:"limitNum"`
	Page     int64 `form:"page"`
}

type CustomerListReply struct {
}

type InGoods struct {
}

type AddGoods struct {
}

type Goods struct {
}

type GoodsList struct {
	LimitNum int64 `form:"limitNum"`
	Page     int64 `form:"page"`
}

type GoodsListReply struct {
}

type RepoList struct {
	GoodsID  string `form:"goodsId"`
	LimitNum int64  `form:"limitNum"`
	Page     int64  `form:"page"`
}

type RepoListReply struct {
}
