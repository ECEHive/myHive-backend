package service

import (
	"github.com/ECEHive/myHive-backend/db"
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/util"
	"sync"
)

var sequenceServiceLogger = util.GetLogger("sequence-service")
var sequenceLocks = map[string]*sync.Mutex{}

func getSequenceWithName(sequenceName string) *entity.Sequence {
	conn := db.GetDB()
	var model entity.Sequence

	if err := conn.FirstOrCreate(&entity.Sequence{
		SequenceName: &sequenceName,
	}).Find(&model).Error; err != nil {
		sequenceServiceLogger.Errorf("Error getting / creating a sequence with name %s: %+v", sequenceName, err)
		return nil
	}
	return &model
}

func SequenceIncrementAndGet(sequenceName string) int64 {
	sequenceServiceLogger.Infof("Requesting Sequence: %s", sequenceName)
	if sequenceLocks[sequenceName] == nil {
		sequenceLocks[sequenceName] = &sync.Mutex{}
	}
	lock := sequenceLocks[sequenceName]
	lock.Lock()
	defer lock.Unlock()

	seq := getSequenceWithName(sequenceName)
	if seq == nil {
		return -1
	}
	seq.SequenceValue += 1
	value := seq.SequenceValue
	if err := db.GetDB().Save(seq).Error; err != nil {
		sequenceServiceLogger.Panicf("Error while saving back sequence with name `%s`: %+v", sequenceName, err)
	}
	return value
}
