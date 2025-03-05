package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCPageSetupService struct {
	*BaseService
}

var PFCPageSetup = &PFCPageSetupService{}

func (s *PFCPageSetupService) GetPFCPageSetup(req *types.PFCModel) (*types.PFC_PageSetup, error) {
	var pageSetup *types.PFC_PageSetup
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(PageSetupID AS NVARCHAR(36)) AS PageSetupID,
		ModelType,
		ModelName,
		MaterialNumber,
		LeftSelectionHeader,
		RightSelectionHeader
	FROM PFC_PageSetup
	WHERE
		ModelType = @ModelType
		AND ModelName = @ModelName
		AND MaterialNumber = @MaterialNumber
`
	err = db.Raw(query,
		sql.Named("ModelType", req.ModelType),
		sql.Named("ModelName", req.ModelName),
		sql.Named("MaterialNumber", req.MaterialNumber),
	).Scan(&pageSetup).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return pageSetup, nil
}

func (s *PFCPageSetupService) InsertNewPFCPageSetup(req *types.PFC_PageSetup) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		INSERT INTO PFC_PageSetup(PageSetupID, ModelType, ModelName, MaterialNumber, LeftSelectionHeader, RightSelectionHeader)
		VALUES
		(NEWID(), ?, ?, ?, ?, ?)
	`
	if err := tx.Exec(
		query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,
		req.LeftSelectionHeader,
		req.RightSelectionHeader,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "success", nil
}

func (s *PFCPageSetupService) UpdatePFCPageSetup(req *types.PFC_PageSetup) (*types.PFC_PageSetup, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		UPDATE PFC_PageSetup
		SET 
			LeftSelectionHeader = ?, 
			RightSelectionHeader = ?
		WHERE
			PageSetupID = ?
	`
	if err := tx.Exec(query,
		req.LeftSelectionHeader,
		req.RightSelectionHeader,
		req.PageSetupID,
	).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to execute update query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return req, nil
}

func (s *PFCPageSetupService) DeletePFCPageSetupByModelID(req *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_PageSetup 
		WHERE PageSetupID = (SELECT TOP(1) PageSetupID FROM PFC_PageSetup WHERE ModelType = ? AND ModelName = ? AND MaterialNumber = ?)
	`

	if err := tx.Exec(query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
