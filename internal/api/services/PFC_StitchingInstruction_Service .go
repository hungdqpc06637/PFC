package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCStitchingInstruction struct {
	*BaseService
}

var PFCStitchingInstructio = &PFCStitchingInstruction{}

func (s *PFCStitchingInstruction) GetAllPFCStitchingInstruction(pfcModel *types.PFCModel) (*[]types.PFC_StitchingInstruction, error) {
	var arrStitchingInstruction *[]types.PFC_StitchingInstruction
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(StitchingInstructionID AS NVARCHAR(36)) AS StitchingInstructionID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_StitchingInstruction
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
	).Scan(&arrStitchingInstruction).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrStitchingInstruction, nil
}

func (s *PFCStitchingInstruction) InsertNewPFCStitchingInstruction(req *types.PFC_StitchingInstruction) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var StitchingInstructionID string

	query := `
		INSERT INTO PFC_StitchingInstruction(StitchingInstructionID, ModelType, ModelName, MaterialNumber, Title,ItemIndex)
		OUTPUT CAST(INSERTED.StitchingInstructionID AS NVARCHAR(36)) AS StitchingInstructionID
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
		&StitchingInstructionID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return StitchingInstructionID, nil
}

func (s *PFCStitchingInstruction) UpdatePFCStitchingInstruction(req *types.PFC_StitchingInstruction) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var StitchingInstructionID string

	query := `
		UPDATE PFC_StitchingInstruction
		SET Title = ?
		OUTPUT CAST(INSERTED.StitchingInstructionID AS NVARCHAR(36))
		WHERE StitchingInstructionID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.StitchingInstructionID,
	).Scan(
		&StitchingInstructionID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return StitchingInstructionID, nil
}

func (s *PFCStitchingInstruction) DeletePFCStitchingInstruction(req *types.PFC_StitchingInstruction) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_StitchingInstruction
		WHERE StitchingInstructionID = ?
	`

	if err := tx.Exec(query,
		req.StitchingInstructionID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_StitchingInstruction

func (s *PFCStitchingInstruction) InsertNewPFCItemStitchingInstruction(req *types.PFC_ItemStitchingInstruction) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemStitchingInstructionID string

	query := `
		INSERT INTO PFC_ItemStitchingInstruction
		(
		ItemStitchingInstructionID,StitchingInstructionID, Component, ImageContent, McType ,NeedleSystem, NeedleSize, 
		NeedlePointType,ThreadType,StitchingMargin,StitchPerInch,AttachingMethod ,  StitchingGuideName, ItemIndex 
		)
		OUTPUT CAST(INSERTED.ItemStitchingInstructionID AS NVARCHAR(36)) AS ItemStitchingInstructionID
		VALUES (NEWID(), ?, ?, ?, ?, ?,?,?, ?, ?, ?, ?,?,?)

	`
	if err := tx.Raw(
		query,
		req.StitchingInstructionID,
		req.Component,
		req.ImageContent,
		req.McType,
		req.NeedleSystem,
		req.NeedleSize,
		req.NeedlePointType,
		req.ThreadType,
		req.StitchingMargin,
		req.StitchPerInch,
		req.AttachingMethod,
		req.StitchingGuideName,
		req.ItemIndex,
	).Scan(
		&ItemStitchingInstructionID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemStitchingInstructionID, nil
}

func (s *PFCStitchingInstruction) GetAllPFCItemStitchingInstruction(pfcStitchingInstruction *types.PFC_StitchingInstruction) (*[]types.PFC_ItemStitchingInstruction, error) {
	var arrItemStitchingInstruction *[]types.PFC_ItemStitchingInstruction
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemStitchingInstructionID AS NVARCHAR(36)) AS ItemStitchingInstructionID,
		CAST(StitchingInstructionID AS NVARCHAR(36)) AS StitchingInstructionID,
		Component, 
		ImageContent, 
		McType ,
		NeedleSystem,
		NeedleSize, 
		NeedlePointType,
		ThreadType,
		StitchingMargin,
		StitchPerInch,AttachingMethod , 
		StitchingGuideName, 
		ItemIndex
	FROM PFC_ItemStitchingInstruction
	WHERE StitchingInstructionID = @StitchingInstructionID
	ORDER BY ItemIndex ASC
	`
	err = db.Raw(query,
		sql.Named("StitchingInstructionID", pfcStitchingInstruction.StitchingInstructionID),
	).Scan(&arrItemStitchingInstruction).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemStitchingInstruction, nil
}

func (s *PFCStitchingInstruction) UpdatePFCItemStitchingInstruction(req *types.PFC_ItemStitchingInstruction) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemStitchingInstructionID string

	query := `
		UPDATE PFC_ItemStitchingInstruction
		SET Component = ?, 
		ImageContent = ?, 
		McType  = ?,
		NeedleSystem = ?,
		NeedleSize = ?, 
		NeedlePointType = ?,
		ThreadType = ?,
		StitchingMargin = ?,
		StitchPerInch  =?,
		AttachingMethod  = ?, 
		StitchingGuideName = ?, 
		ItemIndex = ?
		OUTPUT CAST(INSERTED.ItemStitchingInstructionID AS NVARCHAR(36)) AS ItemStitchingInstructionID
		WHERE ItemStitchingInstructionID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.McType,
		req.NeedleSystem,
		req.NeedleSize,
		req.NeedlePointType,
		req.ThreadType,
		req.StitchingMargin,
		req.StitchPerInch,
		req.AttachingMethod,
		req.StitchingGuideName,
		req.ItemIndex,
		req.ItemStitchingInstructionID,
	).Scan(
		&ItemStitchingInstructionID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemStitchingInstructionID, nil
}

func (s *PFCStitchingInstruction) DeletePFCItemStitchingInstruction(req *types.PFC_ItemStitchingInstruction) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemStitchingInstruction
		WHERE ItemStitchingInstructionID = ?
	`

	if err := tx.Exec(query,
		req.ItemStitchingInstructionID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
