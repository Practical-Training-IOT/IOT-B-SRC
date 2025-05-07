package scene

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/devicePkg"
	rulesReq "github.com/flipped-aurora/gin-vue-admin/server/model/rules/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/scene"
	sceneReq "github.com/flipped-aurora/gin-vue-admin/server/model/scene/request"
	"strconv"
	"time"
)

type ScenesService struct{}

// CreateScenes 创建scenes表记录
// Author [yourname](https://github.com/yourname)
func (scenesService *ScenesService) CreateScenes(ctx context.Context, scenes *scene.Scenes) (err error) {
	err = global.GVA_DB.Create(scenes).Error
	return err
}

// DeleteScenes 删除scenes表记录
// Author [yourname](https://github.com/yourname)
func (scenesService *ScenesService) DeleteScenes(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&scene.Scenes{}, "id = ?", ID).Error
	return err
}

// DeleteScenesByIds 批量删除scenes表记录
// Author [yourname](https://github.com/yourname)
func (scenesService *ScenesService) DeleteScenesByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]scene.Scenes{}, "id in ?", IDs).Error
	return err
}

// UpdateScenes 更新scenes表记录
// Author [yourname](https://github.com/yourname)
func (scenesService *ScenesService) UpdateScenes(ctx context.Context, scenes scene.Scenes) (err error) {
	err = global.GVA_DB.Model(&scene.Scenes{}).Where("id = ?", scenes.ID).Updates(&scenes).Error
	return err
}

// GetScenes 根据ID获取scenes表记录
// Author [yourname](https://github.com/yourname)
func (scenesService *ScenesService) GetScenes(ctx context.Context, ID string) (scenes scene.Scenes, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&scenes).Error
	return
}

// GetScenesInfoList 分页获取scenes表记录
// Author [yourname](https://github.com/yourname)
func (scenesService *ScenesService) GetScenesInfoList(ctx context.Context, info sceneReq.ScenesSearch) (list []scene.Scenes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&scene.Scenes{})
	var sceness []scene.Scenes
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.Scenename != nil && *info.Scenename != "" {
		db = db.Where("scenename LIKE ?", "%"+*info.Scenename+"%")
	}
	if info.Enabledstatus != nil {
		db = db.Where("enabledstatus = ?", *info.Enabledstatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sceness).Error
	return sceness, total, err
}
func (scenesService *ScenesService) GetScenesPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (scenesService *ScenesService) HandleSwitchChange(ctx context.Context, info rulesReq.HandSearch) error {
	err := global.GVA_DB.Model(&scene.Scenes{}).Where("id = ?", info.ID).Update("enabledstatus", info.Status).Error
	if err != nil {
		return err
	}
	return nil
}

func (scenesService *ScenesService) GetScenDevicesList(ctx context.Context, value string) ([]devicePkg.Devices, error) {
	var dev []devicePkg.Devices
	err := global.GVA_DB.Where("belonging_products=?", value).Find(&dev).Error
	if err != nil {
		return nil, err
	}
	return dev, nil
}

func (scenesService *ScenesService) GetScenFuncList(ctx context.Context, value string) ([]scene.Property, error) {
	var cate scene.Category
	err := global.GVA_DB.Model(&scene.Category{}).Where("category_name=?", value).Find(&cate).Error
	var property []scene.Property
	err = global.GVA_DB.Model(&scene.Property{}).Where("category_key=?", cate.CategoryKey).Find(&property).Error
	if err != nil {
		return nil, err
	}
	return property, nil
}

func (scenesService *ScenesService) EnterCreateScenes(ctx context.Context, s *sceneReq.SceneRequest) error {
	var db = global.GVA_DB.Begin()
	var times *time.Time
	var week string
	if s.TriggerMethod == "TimeTriggered" {
		for _, w := range s.TriggerWeekdays {
			week += w + ","
		}
		parsedTime, _ := time.Parse("15:04", s.TriggerConfig.Time)
		times = &parsedTime
	} else {
		times = nil
	}
	triggerCondition := scene.TriggerCondition{
		SceneID:           s.ID,
		TriggerType:       s.TriggerType,
		Product:           s.Product,
		Device:            strconv.Itoa(s.Device),
		Function:          s.Function,
		ValueType:         s.ValueType,
		JudgmentCondition: s.JudgeCondition,
		Time:              times,
		DaysOfWeek:        week,
		TriggerMode:       s.TriggerMethod,
	}
	err := db.Create(&triggerCondition).Error
	if err != nil {
		db.Rollback()
		return err
	}
	var sli []scene.ExecutionAction
	for _, i := range s.HTTPHeaders {
		executionAction := scene.ExecutionAction{
			SceneID:    s.ID,
			ActionType: i.Key,
			Product:    i.Product,
			Device:     strconv.Itoa(i.Device),
			Function:   i.Function,
			Value:      i.Value,
		}
		sli = append(sli, executionAction)
	}
	err = db.Create(&sli).Error
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}
