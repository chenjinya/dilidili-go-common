package model

var (
	// Tags emmm...
	Tags = map[string]map[string][]*Tag{
		"main": map[string][]*Tag{
			"ios": []*Tag{
				{95, "账号相关", 0, ""},
				{481, "播放相关", 0, ""},
				{475, "投稿相关", 0, ""},
				{163, "番剧相关", 0, ""},
				{132, "弹幕", 0, ""},
				{178, "首页推荐", 0, ""},
				{152, "网站内容", 0, ""},
				{192, "免流卡", 0, ""},
				{365, "会员购", 0, ""},
				{478, "创作中心", 0, ""},
				{98, "其他功能建议", 0, ""},
			},
			"android": []*Tag{
				{14, "账号相关", 0, ""},
				{182, "播放相关", 0, ""},
				{476, "投稿相关", 0, ""},
				{162, "番剧相关", 0, ""},
				{27, "弹幕", 0, ""},
				{177, "首页推荐", 0, ""},
				{150, "网站内容", 0, ""},
				{191, "免流卡", 0, ""},
				{363, "会员购", 0, ""},
				{477, "创作中心", 0, ""},
				{43, "其他功能建议", 0, ""},
			},
		},
		"player": map[string][]*Tag{
			"ios": []*Tag{
				{100, "无限小电视", 0, ""},
				//{5, "播放卡顿"},
				//{11, "网页被劫持"},
				//{11, "拖动进度条失败"},
				{101, "黑屏", 0, ""},
				{143, "音画不同步", 0, ""},
				{134, "弹幕无法显示", 0, ""},
			},
			"android": []*Tag{
				{183, "无限小电视", 0, ""},
				{184, "播放卡顿", 0, ""},
				{185, "网页劫持", 0, ""},
				{186, "拖动进度条失败", 0, ""},
				{187, "黑屏", 0, ""},
				{188, "音画不同步", 0, ""},
				{70, "弹幕无法显示", 0, ""},
			},
		},
		"live": map[string][]*Tag{
			"ios": []*Tag{
				{444, "播放卡顿", 0, ""},
				{445, "打开直播间速度慢", 0, ""},
				{452, "无限小电视", 0, ""},
				{450, "音画不同步", 0, ""},
				{446, "黑屏", 0, ""},
				{448, "画面比例异常", 0, ""},
				{447, "绿屏", 0, ""},
				{451, "片段循环", 0, ""},
				{449, "弹幕卡顿", 0, ""},
				{453, "其他问题或建议", 0, ""},
			},
			"android": []*Tag{
				{454, "播放卡顿", 0, ""},
				{455, "打开直播间速度慢", 0, ""},
				{462, "无限小电视", 0, ""},
				{460, "音画不同步", 0, ""},
				{456, "黑屏", 0, ""},
				{457, "画面比例异常", 0, ""},
				{458, "绿屏", 0, ""},
				{461, "片段循环", 0, ""},
				{459, "弹幕卡顿", 0, ""},
				{463, "其他问题或建议", 0, ""},
			},
		},
	}
)

// UGCTag ugc tag.
type UGCTag struct {
	Tags  map[string][]*Tag `json:"tags"`
	Limit int               `json:"limit"`
}

// Tag tag.
type Tag struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Platform string `json:"platform"`
}
