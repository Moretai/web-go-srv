package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = int64(snowflake.GenID())
	return mysql.CreatePost(p)
}

func GetPostById(id int64) (data *models.ApiPostDetail, err error) {

	post, err := mysql.GetPostById(id)
	if err != nil {
		zap.L().Error("mysql.GetPostById(id) failed",
			zap.Int64("post_id", id),
			zap.Error(err))
		return
	}

	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(id) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}

	community, err := mysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetail(post.CommunityID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	return
}
