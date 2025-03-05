package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCAdhesiveTypeService struct {
	*BaseService
}

var PFCAdhesiveType = &PFCAdhesiveTypeService{}

func (s *PFCAdhesiveTypeService) GetAllPFCAdhesiveType(laminationProcessID string) ([]types.PFC_AdhesiveType, error) {
	var list []types.PFC_AdhesiveType
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(AdhesiveTypeID AS NVARCHAR(36)) AS AdhesiveTypeID, 
		CAST(LaminationProcessID AS NVARCHAR(36)) AS LaminationProcessID, 
		Type, 
		Name, 
		Vendor, 
		Thickness, 
		MeltingPoint
	FROM PFC_AdhesiveType
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

func (s *PFCAdhesiveTypeService) InsertPFCAdhesiveType(req *types.PFC_AdhesiveType) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		INSERT INTO PFC_AdhesiveType(AdhesiveTypeID, LaminationProcessID, Type, Name, Vendor, Thickness, MeltingPoint)
		VALUES
		(NEWID(), ?, ?, ?, ?, ?, ?)
	`
	if err := tx.Exec(
		query,
		req.LaminationProcessID,
		req.Type,
		req.Name,
		req.Vendor,
		req.Thickness,
		req.MeltingPoint,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "success", nil
}

func (s *PFCAdhesiveTypeService) UpdatePFCAdhesiveType(req *types.PFC_AdhesiveType) (*types.PFC_AdhesiveType, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		UPDATE PFC_AdhesiveType
		SET 
			Type = ?, 
			Name = ?, 
			Vendor = ?, 
			Thickness = ?, 
			MeltingPoint = ?
		WHERE
			AdhesiveTypeID = ?
	`
	if err := tx.Exec(query,
		req.Type,
		req.Name,
		req.Vendor,
		req.Thickness,
		req.MeltingPoint,
		req.AdhesiveTypeID,
	).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to execute update query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return req, nil
}

func (s *PFCAdhesiveTypeService) DeletePFCAdhesiveType(req *types.PFC_AdhesiveType) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_AdhesiveType
		WHERE AdhesiveTypeID = ?
	`

	if err := tx.Exec(query,
		req.AdhesiveTypeID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

func (s *PFCAdhesiveTypeService) DeletePFCAdhesiveTypeByModelID(req *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_AdhesiveType
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
