package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCAdhesiveOtherTypeService struct {
	*BaseService
}

var PFCAdhesiveOtherType = &PFCAdhesiveOtherTypeService{}

func (s *PFCAdhesiveOtherTypeService) GetAllPFCAdhesiveOtherType(laminationProcessID string) ([]types.PFC_AdhesiveOtherType, error) {
	var list []types.PFC_AdhesiveOtherType
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(AdhesiveOtherTypeID AS NVARCHAR(36)) AS AdhesiveOtherTypeID, 
		CAST(LaminationProcessID AS NVARCHAR(36)) AS LaminationProcessID, 
		Name, 
		Description
	FROM PFC_AdhesiveOtherType
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

func (s *PFCAdhesiveOtherTypeService) InsertNewPFCAdhesiveOtherType(req *types.PFC_AdhesiveOtherType) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		INSERT INTO PFC_AdhesiveOtherType(AdhesiveOtherTypeID, LaminationProcessID, Name, Description)
		VALUES
		(NEWID(), ?, ?, ?)
	`
	if err := tx.Exec(
		query,
		req.LaminationProcessID,
		req.Name,
		req.Description,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "success", nil
}

func (s *PFCAdhesiveOtherTypeService) UpdatePFCAdhesiveOtherType(req *types.PFC_AdhesiveOtherType) (*types.PFC_AdhesiveOtherType, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		UPDATE PFC_AdhesiveOtherType
		SET 
			Name = ?, 
			Description = ?
		WHERE
			AdhesiveOtherTypeID = ?
	`
	if err := tx.Exec(query,
		req.Name,
		req.Description,
		req.AdhesiveOtherTypeID,
	).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to execute update query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return req, nil
}

func (s *PFCAdhesiveOtherTypeService) DeletePFCAdhesiveOtherType(req *types.PFC_AdhesiveOtherType) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_AdhesiveOtherType
		WHERE AdhesiveOtherTypeID = ?
	`

	if err := tx.Exec(query,
		req.AdhesiveOtherTypeID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

func (s *PFCAdhesiveOtherTypeService) DeletePFCAdhesiveOtherTypeByModelID(req *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_AdhesiveOtherType
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
