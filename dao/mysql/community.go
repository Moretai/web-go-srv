package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"web_app/models"
)

func GetCommunityList()(communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}

	return
}

func GetCommunityDetail(communityID int64)(*models.CommunityDetail, error) {
	sqlStr := `select 
       id, community_id,  introduction, create_time,update_time
       from community
	   where community_id=?`
	community := new(models.CommunityDetail)
	if err := db.Get(community, sqlStr, communityID); err != nil {
		zap.L().Error("db.Get(community, sqlStr, communityID)", zap.Error(err))
		if errors.Is(err,sql.ErrNoRows) {
			zap.L().Warn("there is no community detail where community_id is")
			err = ErrorInvalidID
			return nil, err
		}
	}

	zap.L().Info("GetCommunityDetail ", zap.Any("ID", communityID),zap.Any("data", community))

	return community, nil
}