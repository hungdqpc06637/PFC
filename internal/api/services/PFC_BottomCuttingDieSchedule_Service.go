package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCBottomCuttingDieSchedule struct {
	*BaseService
}

var PFCBottomCuttingDieSchedul = &PFCBottomCuttingDieSchedule{}

func (s *PFCBottomCuttingDieSchedule) GetAllPFCBottomCuttingDieSchedule(pfcModel *types.PFCModel) (*[]types.PFC_BottomCuttingDieSchedule, error) {
	var arrBottomCuttingDieSchedule *[]types.PFC_BottomCuttingDieSchedule
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(BottomCuttingDieScheduleID AS NVARCHAR(36)) AS BottomCuttingDieScheduleID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_BottomCuttingDieSchedule
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
	).Scan(&arrBottomCuttingDieSchedule).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrBottomCuttingDieSchedule, nil
}

func (s *PFCBottomCuttingDieSchedule) InsertNewPFCBottomCuttingDieSchedule(req *types.PFC_BottomCuttingDieSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var BottomCuttingDieScheduleID string

	query := `
		INSERT INTO PFC_BottomCuttingDieSchedule(BottomCuttingDieScheduleID, ModelType, ModelName, MaterialNumber, Title,ItemIndex)
		OUTPUT CAST(INSERTED.BottomCuttingDieScheduleID AS NVARCHAR(36)) AS BottomCuttingDieScheduleID
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
		&BottomCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return BottomCuttingDieScheduleID, nil
}

func (s *PFCBottomCuttingDieSchedule) UpdatePFCBottomCuttingDieSchedule(req *types.PFC_BottomCuttingDieSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var BottomCuttingDieScheduleID string

	query := `
		UPDATE PFC_BottomCuttingDieSchedule
		SET Title = ?
		OUTPUT CAST(INSERTED.BottomCuttingDieScheduleID AS NVARCHAR(36))
		WHERE BottomCuttingDieScheduleID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.BottomCuttingDieScheduleID,
	).Scan(
		&BottomCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return BottomCuttingDieScheduleID, nil
}

func (s *PFCBottomCuttingDieSchedule) DeletePFCBottomCuttingDieSchedule(req *types.PFC_BottomCuttingDieSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_BottomCuttingDieSchedule
		WHERE BottomCuttingDieScheduleID = ?
	`

	if err := tx.Exec(query,
		req.BottomCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_BottomCuttingDieSchedule

func (s *PFCBottomCuttingDieSchedule) InsertNewPFCItemBottomCuttingDieSchedule(req *types.PFC_ItemBottomCuttingDieSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemBottomCuttingDieScheduleID string

	query := `
		INSERT INTO PFC_ItemBottomCuttingDieSchedule
		(
		ItemBottomCuttingDieScheduleID,BottomCuttingDieScheduleID, Component, ImageContent, SizeRange1 ,SizeRange2, SizeRange,SizeRangeAreSame, Remarks,NumberOfLayers,Thickness,Width,Hardness, ItemIndex 
		)
		OUTPUT CAST(INSERTED.ItemBottomCuttingDieScheduleID AS NVARCHAR(36)) AS ItemBottomCuttingDieScheduleID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)

	`
	if err := tx.Raw(
		query,
		req.BottomCuttingDieScheduleID,
		req.Component,
		req.ImageContent,
		req.SizeRange1,
		req.SizeRange2,
		req.SizeRange,
		req.SizeRangeAreSame,
		req.Remarks,
		req.NumberOfLayers,
		req.Thickness,
		req.Width,
		req.Hardness,
		req.ItemIndex,
	).Scan(
		&ItemBottomCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemBottomCuttingDieScheduleID, nil
}

func (s *PFCBottomCuttingDieSchedule) GetAllPFCItemBottomCuttingDieSchedule(pfcBottomCuttingDieSchedule *types.PFC_BottomCuttingDieSchedule) (*[]types.PFC_ItemBottomCuttingDieSchedule, error) {
	var arrItemBottomCuttingDieSchedule *[]types.PFC_ItemBottomCuttingDieSchedule
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemBottomCuttingDieScheduleID AS NVARCHAR(36)) AS ItemBottomCuttingDieScheduleID,
		CAST(BottomCuttingDieScheduleID AS NVARCHAR(36)) AS BottomCuttingDieScheduleID,
		Component, 
		ImageContent,
		SizeRange1 ,
		SizeRange2,
		SizeRange,
		SizeRangeAreSame, 
		Remarks,
		NumberOfLayers,
		Thickness,
		Width,
		Hardness, 
		ItemIndex 
	FROM PFC_ItemBottomCuttingDieSchedule
	WHERE BottomCuttingDieScheduleID = @BottomCuttingDieScheduleID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("BottomCuttingDieScheduleID", pfcBottomCuttingDieSchedule.BottomCuttingDieScheduleID),
	).Scan(&arrItemBottomCuttingDieSchedule).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemBottomCuttingDieSchedule, nil
}

func (s *PFCBottomCuttingDieSchedule) UpdatePFCItemBottomCuttingDieSchedule(req *types.PFC_ItemBottomCuttingDieSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemBottomCuttingDieScheduleID string

	query := `
		UPDATE PFC_ItemBottomCuttingDieSchedule
		SET Component = ?, 
		ImageContent = ?, 
		SizeRange1  = ?,
		SizeRange2 = ?, 
		SizeRange = ?,
		SizeRangeAreSame = ?, 
		Remarks = ?,
		NumberOfLayers = ?,
		Thickness = ?,
		Width = ?,
		Hardness = ?, 
		ItemIndex  = ?
		OUTPUT CAST(INSERTED.ItemBottomCuttingDieScheduleID AS NVARCHAR(36)) AS ItemBottomCuttingDieScheduleID
		WHERE ItemBottomCuttingDieScheduleID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.SizeRange1,
		req.SizeRange2,
		req.SizeRange,
		req.SizeRangeAreSame,
		req.Remarks,
		req.NumberOfLayers,
		req.Thickness,
		req.Width,
		req.Hardness,
		req.ItemIndex,
		req.ItemBottomCuttingDieScheduleID,
	).Scan(
		&ItemBottomCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemBottomCuttingDieScheduleID, nil
}

func (s *PFCBottomCuttingDieSchedule) DeletePFCItemBottomCuttingDieSchedule(req *types.PFC_ItemBottomCuttingDieSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemBottomCuttingDieSchedule
		WHERE ItemBottomCuttingDieScheduleID = ?
	`

	if err := tx.Exec(query,
		req.ItemBottomCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
