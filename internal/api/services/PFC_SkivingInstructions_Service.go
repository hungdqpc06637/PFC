package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCSkivingInstructions struct {
	*BaseService
}

var PFCSkivingInstruction = &PFCSkivingInstructions{}

func (s *PFCSkivingInstructions) GetAllPFCSkivingInstructions(pfcModel *types.PFCModel) (*[]types.PFC_SkivingInstructions, error) {
	var arrSkivingInstructions *[]types.PFC_SkivingInstructions
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(SkivingInstructionsID AS NVARCHAR(36)) AS SkivingInstructionsID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		SkivingKey,
		ItemIndex
	FROM PFC_SkivingInstructions
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
	).Scan(&arrSkivingInstructions).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrSkivingInstructions, nil
}

func (s *PFCSkivingInstructions) InsertNewPFCSkivingInstructions(req *types.PFC_SkivingInstructions) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SkivingInstructionsID string

	query := `
		INSERT INTO PFC_SkivingInstructions(SkivingInstructionsID, ModelType, ModelName, MaterialNumber, Title, SkivingKey,ItemIndex)
		OUTPUT CAST(INSERTED.SkivingInstructionsID AS NVARCHAR(36)) AS SkivingInstructionsID
		VALUES (NEWID(), ?, ?, ?, ?, ?,?)

	`
	if err := tx.Raw(
		query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,
		req.Title,
		req.SkivingKey,
		req.ItemIndex,
	).Scan(
		&SkivingInstructionsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SkivingInstructionsID, nil
}

func (s *PFCSkivingInstructions) UpdatePFCSkivingInstructions(req *types.PFC_SkivingInstructions) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SkivingInstructionsID string

	query := `
		UPDATE PFC_SkivingInstructions
		SET 
			Title = ?, 
			SkivingKey = ?
		OUTPUT CAST(INSERTED.SkivingInstructionsID AS NVARCHAR(36))
		WHERE SkivingInstructionsID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.SkivingKey,
		req.SkivingInstructionsID,
	).Scan(
		&SkivingInstructionsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SkivingInstructionsID, nil
}

func (s *PFCSkivingInstructions) DeletePFCSkivingInstructions(req *types.PFC_SkivingInstructions) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_SkivingInstructions
		WHERE SkivingInstructionsID = ?
	`

	if err := tx.Exec(query,
		req.SkivingInstructionsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_SkivingInstructions

func (s *PFCSkivingInstructions) InsertNewPFCItemSkivingInstructions(req *types.PFC_ItemSkivingInstructions) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSkivingInstructionsID string

	query := `
		INSERT INTO PFC_ItemSkivingInstructions(ItemSkivingInstructionsID, SkivingInstructionsID, Component, ImageContent, SkivedEdgeThickness, SkivingWidth, ItemIndex)
		OUTPUT CAST(INSERTED.ItemSkivingInstructionsID AS NVARCHAR(36)) AS ItemSkivingInstructionsID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.SkivingInstructionsID,
		req.Component,
		req.ImageContent,
		req.SkivedEdgeThickness,
		req.SkivingWidth,
		req.ItemIndex,
	).Scan(
		&ItemSkivingInstructionsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSkivingInstructionsID, nil
}

func (s *PFCSkivingInstructions) GetAllPFCItemSkivingInstructions(pfcSkivingInstructions *types.PFC_SkivingInstructions) (*[]types.PFC_ItemSkivingInstructions, error) {
	var arrItemSkivingInstructions *[]types.PFC_ItemSkivingInstructions
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemSkivingInstructionsID AS NVARCHAR(36)) AS ItemSkivingInstructionsID,
		CAST(SkivingInstructionsID AS NVARCHAR(36)) AS SkivingInstructionsID,
		Component, 
		ImageContent,
		SkivedEdgeThickness,
		SkivingWidth,
		ItemIndex
	FROM PFC_ItemSkivingInstructions
	WHERE SkivingInstructionsID = @SkivingInstructionsID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("SkivingInstructionsID", pfcSkivingInstructions.SkivingInstructionsID),
	).Scan(&arrItemSkivingInstructions).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemSkivingInstructions, nil
}

func (s *PFCSkivingInstructions) UpdatePFCItemSkivingInstructions(req *types.PFC_ItemSkivingInstructions) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSkivingInstructionsID string

	query := `
		UPDATE PFC_ItemSkivingInstructions
		SET Component = ?,
		ImageContent  = ?,
		SkivedEdgeThickness  = ?,
		SkivingWidth  = ?,
		ItemIndex  = ?
		OUTPUT CAST(INSERTED.ItemSkivingInstructionsID AS NVARCHAR(36)) AS ItemSkivingInstructionsID
		WHERE ItemSkivingInstructionsID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.SkivedEdgeThickness,
		req.SkivingWidth,
		req.ItemIndex,
		req.ItemSkivingInstructionsID,
	).Scan(
		&ItemSkivingInstructionsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSkivingInstructionsID, nil
}

func (s *PFCSkivingInstructions) DeletePFCItemSkivingInstructions(req *types.PFC_ItemSkivingInstructions) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemSkivingInstructions
		WHERE ItemSkivingInstructionsID = ?
	`

	if err := tx.Exec(query,
		req.ItemSkivingInstructionsID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
