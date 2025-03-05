package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCComputerStitchingSchedule struct {
	*BaseService
}

var PFCComputerStitchingSchedu = &PFCComputerStitchingSchedule{}

func (s *PFCComputerStitchingSchedule) GetAllPFCComputerStitchingSchedule(pfcModel *types.PFCModel) (*[]types.PFC_ComputerStitchingSchedule, error) {
	var arrComputerStitchingSchedule *[]types.PFC_ComputerStitchingSchedule
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(ComputerStitchingScheduleID AS NVARCHAR(36)) AS ComputerStitchingScheduleID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_ComputerStitchingSchedule
	WHERE
		ModelType = @ModelType
		AND ModelName = @ModelName
		AND MaterialNumber = @MaterialNumber
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("ModelType", pfcModel.ModelType),
		sql.Named("ModelName", pfcModel.ModelName),
		sql.Named("MaterialNumber", pfcModel.MaterialNumber),
	).Scan(&arrComputerStitchingSchedule).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrComputerStitchingSchedule, nil
}

func (s *PFCComputerStitchingSchedule) InsertNewPFCComputerStitchingSchedule(req *types.PFC_ComputerStitchingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ComputerStitchingScheduleID string

	query := `
		INSERT INTO PFC_ComputerStitchingSchedule(ComputerStitchingScheduleID, ModelType, ModelName, MaterialNumber, Title,ItemIndex)
		OUTPUT CAST(INSERTED.ComputerStitchingScheduleID AS NVARCHAR(36)) AS ComputerStitchingScheduleID
		VALUES (NEWID(), ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,
		req.Title,
		req.ItemIndex,
	).Scan(
		&ComputerStitchingScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ComputerStitchingScheduleID, nil
}

func (s *PFCComputerStitchingSchedule) UpdatePFCComputerStitchingSchedule(req *types.PFC_ComputerStitchingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ComputerStitchingScheduleID string

	query := `
		UPDATE PFC_ComputerStitchingSchedule
		SET Title = ?
		OUTPUT CAST(INSERTED.ComputerStitchingScheduleID AS NVARCHAR(36))
		WHERE ComputerStitchingScheduleID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.ComputerStitchingScheduleID,
	).Scan(
		&ComputerStitchingScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ComputerStitchingScheduleID, nil
}

func (s *PFCComputerStitchingSchedule) DeletePFCComputerStitchingSchedule(req *types.PFC_ComputerStitchingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ComputerStitchingSchedule
		WHERE ComputerStitchingScheduleID = ?
	`

	if err := tx.Exec(query,
		req.ComputerStitchingScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ComputerStitchingSchedule

func (s *PFCComputerStitchingSchedule) InsertNewPFCItemComputerStitchingSchedule(req *types.PFC_ItemComputerStitchingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemComputerStitchingScheduleID string

	query := `
		INSERT INTO PFC_ItemComputerStitchingSchedule
		(
		ItemComputerStitchingScheduleID, ComputerStitchingScheduleID, Component, ImageContent, StitchingMargin ,NeedleTypeSize,StitchPerInch,Size, ItemIndex 
		)
		OUTPUT CAST(INSERTED.ItemComputerStitchingScheduleID AS NVARCHAR(36)) AS ItemComputerStitchingScheduleID
		VALUES (NEWID(), ?, ?, ?, ?, ?,?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.ComputerStitchingScheduleID,
		req.Component,
		req.ImageContent,
		req.StitchingMargin,
		req.NeedleTypeSize,
		req.StitchPerInch,
		req.Size,
		req.ItemIndex,
	).Scan(
		&ItemComputerStitchingScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemComputerStitchingScheduleID, nil
}

func (s *PFCComputerStitchingSchedule) GetAllPFCItemComputerStitchingSchedule(pfcComputerStitchingSchedule *types.PFC_ComputerStitchingSchedule) (*[]types.PFC_ItemComputerStitchingSchedule, error) {
	var arrItemComputerStitchingSchedule *[]types.PFC_ItemComputerStitchingSchedule
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemComputerStitchingScheduleID AS NVARCHAR(36)) AS ItemComputerStitchingScheduleID,
		CAST(ComputerStitchingScheduleID AS NVARCHAR(36)) AS ComputerStitchingScheduleID,
		Component, 
		ImageContent, 
		StitchingMargin,
		NeedleTypeSize,
		StitchPerInch,	
		Size,
		ItemIndex
	FROM PFC_ItemComputerStitchingSchedule
	WHERE ComputerStitchingScheduleID = @ComputerStitchingScheduleID
	ORDER BY ItemIndex ASC
	`
	err = db.Raw(query,
		sql.Named("ComputerStitchingScheduleID", pfcComputerStitchingSchedule.ComputerStitchingScheduleID),
	).Scan(&arrItemComputerStitchingSchedule).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemComputerStitchingSchedule, nil
}

func (s *PFCComputerStitchingSchedule) UpdatePFCItemComputerStitchingSchedule(req *types.PFC_ItemComputerStitchingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemComputerStitchingScheduleID string

	query := `
		UPDATE PFC_ItemComputerStitchingSchedule
		SET Component = ?,
		ImageContent  = ?,
		StitchingMargin = ? ,
		NeedleTypeSize = ?,
		StitchPerInch= ?,
		Size = ?,
		ItemIndex  = ?
		OUTPUT CAST(INSERTED.ItemComputerStitchingScheduleID AS NVARCHAR(36)) AS ItemComputerStitchingScheduleID
		WHERE ItemComputerStitchingScheduleID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.StitchingMargin,
		req.NeedleTypeSize,
		req.StitchPerInch,
		req.Size,
		req.ItemIndex,
		req.ItemComputerStitchingScheduleID,
	).Scan(
		&ItemComputerStitchingScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemComputerStitchingScheduleID, nil
}

func (s *PFCComputerStitchingSchedule) DeletePFCItemComputerStitchingSchedule(req *types.PFC_ItemComputerStitchingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemComputerStitchingSchedule
		WHERE ItemComputerStitchingScheduleID = ?
	`

	if err := tx.Exec(query,
		req.ItemComputerStitchingScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
