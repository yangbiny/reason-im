package it

import (
	"context"
	"reason-im/internal/config/mysql"
	"reason-im/internal/repo"
	mysql2 "reason-im/internal/utils/mysql"
	"testing"
)

func TestDatabaseTpl_FindOne(t *testing.T) {
	datasource := mysql.Datasource()
	tpl := mysql2.NewDatabaseTpl(datasource)

	background := context.Background()
	tpl.WithTransaction(&background, func(ctx *context.Context) error {
		_, err := tpl.Insert(ctx, "insert into im_group ( name, description, group_type, group_member_cnt, gmt_create, gmt_update) values ('test', 'test',1, 1, now(), now())")
		if err != nil {
			println(err.Error())
		}

		result, err := tpl.FindOne(ctx, "select * from im_group where id = 1", repo.GroupDO{})
		if result != nil {
			println(result)
		}
		if err != nil {
			return err
		}

		_, err = tpl.Insert(ctx, "insert into im_group_member (group_id, user_id, nick_name, group_member_role, gmt_create, gmt_update) values ( 1, 1, 'test', 1, now(), now())")
		if err != nil {
			return err
		}

		return nil
	})
}
