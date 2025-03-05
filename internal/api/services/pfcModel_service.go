package services

import (
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCModelService struct {
	*BaseService
}

var PFCModel = &PFCModelService{}

func (s *PFCModelService) GetAllPFCModel() ([]types.PFCModel, error) {
	var list []types.PFCModel
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		ModelType,
		ModelName ,
		MaterialNumber,
		ModelID,
		ColorWayID,
		BOMID,
		Date,
		SizeRange,
		LastCode,
		ToolCode,
		PatternFileName,
		FirstSource,
		MedialSideView,
		LateralSideView,
		BottomView,
		FrontView,
		HeelView,
		KeyManufacturingProcesses,
		IDS,
		Converse
	FROM PFC_Model
`
	err = db.Raw(query).Scan(&list).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return list, nil
}

func (s *PFCModelService) InsertNewModel(requestParams *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		INSERT INTO PFC_Model (
			ModelType, ModelName, MaterialNumber,
			ModelID, ColorWayID, BOMID,
			Date, SizeRange, LastCode,
			ToolCode, PatternFileName, FirstSource,
			MedialSideView, LateralSideView, BottomView,
			FrontView, HeelView, KeyManufacturingProcesses,
			IDS, Converse
		) 
		VALUES (?, ?, ?, 
				?, ?, ?, 
				?, ?, ?, 
				?, ?, ?, 
				?, ?, ?, 
				?, ?, ?, 
				?, ?)
	`

	if err := tx.Exec(query,
		requestParams.ModelType, requestParams.ModelName, requestParams.MaterialNumber,
		requestParams.ModelID, requestParams.ColorWayID, requestParams.BOMID,
		requestParams.Date, requestParams.SizeRange, requestParams.LastCode,
		requestParams.ToolCode, requestParams.PatternFileName, requestParams.FirstSource,
		requestParams.MedialSideView, requestParams.LateralSideView, requestParams.BottomView,
		requestParams.FrontView, requestParams.HeelView, requestParams.KeyManufacturingProcesses,
		requestParams.IDS, requestParams.Converse,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "success", nil
}

func (s *PFCModelService) UpdatePFCModel(requestParams *types.PFCModel) (*types.PFCModel, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		UPDATE PFC_MODEL
		SET 
			ModelType = ?,
			ModelName = ?, 
			MaterialNumber = ?, 
			ModelID = ?,
			ColorWayID = ?,
			BOMID = ?,
			Date = ?, 
			SizeRange = ?, 
			LastCode = ?, 
			ToolCode = ?, 
			PatternFileName = ?, 
			FirstSource = ?, 
			MedialSideView = ?, 
			LateralSideView = ?, 
			BottomView = ?, 
			FrontView = ?, 
			HeelView = ?,
			KeyManufacturingProcesses = ?,
			IDS = ?,
			Converse = ?
		WHERE 
			ModelID = ?
			AND ColorWayID = ?
			AND BOMID = ?
	`

	if err := tx.Exec(query,
		requestParams.ModelType,
		requestParams.ModelName,
		requestParams.MaterialNumber,
		requestParams.ModelID,
		requestParams.ColorWayID,
		requestParams.BOMID,
		requestParams.Date,
		requestParams.SizeRange,
		requestParams.LastCode,
		requestParams.ToolCode,
		requestParams.PatternFileName,
		requestParams.FirstSource,
		requestParams.MedialSideView,
		requestParams.LateralSideView,
		requestParams.BottomView,
		requestParams.FrontView,
		requestParams.HeelView,
		requestParams.KeyManufacturingProcesses,
		requestParams.IDS,
		requestParams.Converse,
		requestParams.ModelID,
		requestParams.ColorWayID,
		requestParams.BOMID,
	).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to execute update query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return requestParams, nil
}

func (s *PFCModelService) DeletePFCModel(requestParams *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_MODEL
		WHERE ModelType = ?
			AND ModelName = ?
			AND MaterialNumber = ?
	`

	if err := tx.Exec(query,
		requestParams.ModelType,
		requestParams.ModelName,
		requestParams.MaterialNumber).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
