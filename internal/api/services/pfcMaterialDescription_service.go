package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCMaterialDescriptionService struct {
	*BaseService
}

var PFCMaterialDescription = &PFCMaterialDescriptionService{}

func (s *PFCMaterialDescriptionService) GetAllPFCMaterialDescription(laminationProcessID string) ([]types.PFC_MaterialDescription, error) {
	var list []types.PFC_MaterialDescription
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(MaterialDescriptionID AS NVARCHAR(36)) AS MaterialDescriptionID, 
		CAST(LaminationProcessID AS NVARCHAR(36)) AS LaminationProcessID, 
		Name, 
		Mat
	FROM PFC_MaterialDescription
	WHERE LaminationProcessID = @LaminationProcessID
`
	err = db.Raw(query, sql.Named("LaminationProcessID", laminationProcessID)).Scan(&list).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return list, nil
}

func (s *PFCMaterialDescriptionService) InsertNewPFCMaterialDescription(req *types.PFC_MaterialDescription) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		INSERT INTO PFC_MaterialDescription(MaterialDescriptionID, LaminationProcessID, Name, Mat)
		VALUES
		(NEWID(), ?, ?, ?)
	`
	if err := tx.Exec(
		query,
		req.LaminationProcessID,
		req.Name,
		req.Mat,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "success", nil
}

func (s *PFCMaterialDescriptionService) UpdatePFCMaterialDescription(req *types.PFC_MaterialDescription) (*types.PFC_MaterialDescription, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		UPDATE PFC_MaterialDescription
		SET 
			Name = ?, 
			Mat = ?
		WHERE
			MaterialDescriptionID = ?
	`
	if err := tx.Exec(query,
		req.Name,
		req.Mat,
		req.MaterialDescriptionID,
	).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to execute update query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return req, nil
}

func (s *PFCMaterialDescriptionService) DeletePFCMaterialDescription(req *types.PFC_MaterialDescription) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_MaterialDescription
		WHERE MaterialDescriptionID = ?
	`

	if err := tx.Exec(query,
		req.MaterialDescriptionID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

func (s *PFCMaterialDescriptionService) DeletePFCMaterialDescriptionByModelID(req *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_MaterialDescription
		WHERE LaminationProcessID = (SELECT TOP(1) LaminationProcessID FROM PFC_LaminationProcess WHERE ModelType = ? AND ModelName = ? AND MaterialNumber = ?)
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
