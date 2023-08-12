package service

import "tiktok/dao"

func HandleSentMes(act int64, mes string, uid int64, tid int64) error {
	if act == 1 {
		err := dao.CreateMes(uid, tid, mes)
		if err != nil {
			return err
		}
	}
	return nil
}
