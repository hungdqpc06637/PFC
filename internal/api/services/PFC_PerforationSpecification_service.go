package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCPerforationSpecificationService struct {
	*BaseService
}

var PFCPerforationSpecification = &PFCPerforationSpecificationService{}

func (s *PFCPerforationSpecificationService) GetAllPFCPerforationSpecification(pfcModel *types.PFCModel) (*[]types.PFC_PerforationSpecification, error) {
	var arrPerforationSpecification *[]types.PFC_PerforationSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(PerforationSpecificationID AS NVARCHAR(36)) AS PerforationSpecificationID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_PerforationSpecification
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
	).Scan(&arrPerforationSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrPerforationSpecification, nil
}

func (s *PFCPerforationSpecificationService) InsertNewPFCPerforationSpecification(req *types.PFC_PerforationSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var perforationSpecificationID string

	query := `
		INSERT INTO PFC_PerforationSpecification(PerforationSpecificationID, ModelType, ModelName, MaterialNumber, Title, ItemIndex)
		OUTPUT CAST(INSERTED.PerforationSpecificationID AS NVARCHAR(36)) AS PerforationSpecificationID
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
		&perforationSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return perforationSpecificationID, nil
}

func (s *PFCPerforationSpecificationService) UpdatePFCPerforationSpecification(req *types.PFC_PerforationSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var perforationSpecificationID string

	query := `
		UPDATE PFC_PerforationSpecification
		SET Title = ?
		OUTPUT CAST(INSERTED.PerforationSpecificationID AS NVARCHAR(36)) AS PerforationSpecificationID
		WHERE PerforationSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.Title,
		req.PerforationSpecificationID,
	).Scan(
		&perforationSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return perforationSpecificationID, nil
}

func (s *PFCPerforationSpecificationService) DeletePFCPerforationSpecification(req *types.PFC_PerforationSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_PerforationSpecification
		WHERE PerforationSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.PerforationSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// ITEM PFCPerforationSpecification
func (s *PFCPerforationSpecificationService) InsertNewPFCItemPerforationSpecification(req *types.PFC_ItemPerforationSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemPerforationSpecificationID string

	query := `
		INSERT INTO PFC_ItemPerforationSpecification(ItemPerforationSpecificationID, PerforationSpecificationID, Component, ImageContent, SizeGroup1, SizeGroup2, SizeGroup3, ItemIndex)
		OUTPUT CAST(INSERTED.ItemPerforationSpecificationID AS NVARCHAR(36)) AS ItemPerforationSpecificationID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.PerforationSpecificationID,
		req.Component,
		req.ImageContent,
		req.SizeGroup1,
		req.SizeGroup2,
		req.SizeGroup3,
		req.ItemIndex,
	).Scan(
		&ItemPerforationSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemPerforationSpecificationID, nil
}

func (s *PFCPerforationSpecificationService) GetAllPFCItemPerforationSpecification(pfcPerforationSpecification *types.PFC_PerforationSpecification) (*[]types.PFC_ItemPerforationSpecification, error) {
	var arrItemPerforationSpecification *[]types.PFC_ItemPerforationSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemPerforationSpecificationID AS NVARCHAR(36)) AS ItemPerforationSpecificationID,
		CAST(PerforationSpecificationID AS NVARCHAR(36)) AS PerforationSpecificationID,
		Component,
		ImageContent,
		SizeGroup1,
		SizeGroup2,
		SizeGroup3,
		ItemIndex
	FROM PFC_ItemPerforationSpecification
	WHERE PerforationSpecificationID = @PerforationSpecificationID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("PerforationSpecificationID", pfcPerforationSpecification.PerforationSpecificationID),
	).Scan(&arrItemPerforationSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemPerforationSpecification, nil
}

func (s *PFCPerforationSpecificationService) UpdatePFCItemPerforationSpecification(req *types.PFC_ItemPerforationSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemPerforationSpecificationID string

	query := `
		UPDATE PFC_ItemPerforationSpecification
		SET Component = ?,
			ImageContent = ?,
			SizeGroup1 = ?,
			SizeGroup2 = ?,
			SizeGroup3 = ?,
			ItemIndex = ?
		OUTPUT CAST(INSERTED.ItemPerforationSpecificationID AS NVARCHAR(36)) AS ItemPerforationSpecificationID
		WHERE ItemPerforationSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.SizeGroup1,
		req.SizeGroup2,
		req.SizeGroup3,
		req.ItemIndex,
		req.ItemPerforationSpecificationID,
	).Scan(
		&ItemPerforationSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemPerforationSpecificationID, nil
}

func (s *PFCPerforationSpecificationService) DeletePFCItemPerforationSpecification(req *types.PFC_ItemPerforationSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemPerforationSpecification
		WHERE ItemPerforationSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ItemPerforationSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
