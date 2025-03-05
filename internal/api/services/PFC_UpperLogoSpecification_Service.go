package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCUpperLogoSpecificationService struct {
	*BaseService
}

var PFCUpperLogoSpecification = &PFCUpperLogoSpecificationService{}

func (s *PFCUpperLogoSpecificationService) GetAllPFCUpperLogoSpecification(pfcModel *types.PFCModel) (*[]types.PFC_UpperLogoSpecification, error) {
	var arrUpperLogoSpecification *[]types.PFC_UpperLogoSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(UpperLogoSpecificationID AS NVARCHAR(36)) AS UpperLogoSpecificationID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_UpperLogoSpecification
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
	).Scan(&arrUpperLogoSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrUpperLogoSpecification, nil
}

func (s *PFCUpperLogoSpecificationService) InsertNewPFCUpperLogoSpecification(req *types.PFC_UpperLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var upperLogoSpecificationID string

	query := `
		INSERT INTO PFC_UpperLogoSpecification(UpperLogoSpecificationID, ModelType, ModelName, MaterialNumber, Title, ItemIndex)
		OUTPUT CAST(INSERTED.UpperLogoSpecificationID AS NVARCHAR(36)) AS UpperLogoSpecificationID
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
		&upperLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return upperLogoSpecificationID, nil
}

func (s *PFCUpperLogoSpecificationService) UpdatePFCUpperLogoSpecification(req *types.PFC_UpperLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var upperLogoSpecificationID string

	query := `
		UPDATE PFC_UpperLogoSpecification
		SET Title = ?
		OUTPUT CAST(INSERTED.UpperLogoSpecificationID AS NVARCHAR(36)) AS UpperLogoSpecificationID
		WHERE UpperLogoSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.Title,
		req.UpperLogoSpecificationID,
	).Scan(
		&upperLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return upperLogoSpecificationID, nil
}

func (s *PFCUpperLogoSpecificationService) DeletePFCUpperLogoSpecification(req *types.PFC_UpperLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_UpperLogoSpecification
		WHERE UpperLogoSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.UpperLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// ITEM PFC_UpperLogoSpecification
func (s *PFCUpperLogoSpecificationService) InsertNewPFCItemUpperLogoSpecification(req *types.PFC_ItemUpperLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemUpperLogoSpecificationID string

	query := `
		INSERT INTO PFC_ItemUpperLogoSpecification(ItemUpperLogoSpecificationID, UpperLogoSpecificationID, Component, Vendor, ImageContent, Material, TableRow1, TableRow2, TableRow3, TableRow4, TableRow5, TableRow6, ItemIndex)
		OUTPUT CAST(INSERTED.ItemUpperLogoSpecificationID AS NVARCHAR(36)) AS ItemUpperLogoSpecificationID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.UpperLogoSpecificationID,
		req.Component,
		req.Vendor,
		req.ImageContent,
		req.Material,
		req.TableRow1,
		req.TableRow2,
		req.TableRow3,
		req.TableRow4,
		req.TableRow5,
		req.TableRow6,
		req.ItemIndex,
	).Scan(
		&ItemUpperLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemUpperLogoSpecificationID, nil
}

func (s *PFCUpperLogoSpecificationService) GetAllPFCItemUpperLogoSpecification(pfcUpperLogoSpecification *types.PFC_UpperLogoSpecification) (*[]types.PFC_ItemUpperLogoSpecification, error) {
	var arrItemUpperLogoSpecification *[]types.PFC_ItemUpperLogoSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemUpperLogoSpecificationID AS NVARCHAR(36)) AS ItemUpperLogoSpecificationID,
		CAST(UpperLogoSpecificationID AS NVARCHAR(36)) AS UpperLogoSpecificationID,
		Component,
		Vendor,
		ImageContent,
		Material,
		TableRow1,
		TableRow2,
		TableRow3,
		TableRow4,
		TableRow5,
		TableRow6,
		ItemIndex
	FROM PFC_ItemUpperLogoSpecification
	WHERE UpperLogoSpecificationID = @UpperLogoSpecificationID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("UpperLogoSpecificationID", pfcUpperLogoSpecification.UpperLogoSpecificationID),
	).Scan(&arrItemUpperLogoSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemUpperLogoSpecification, nil
}

func (s *PFCUpperLogoSpecificationService) UpdatePFCItemUpperLogoSpecification(req *types.PFC_ItemUpperLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemUpperLogoSpecificationID string

	query := `
		UPDATE PFC_ItemUpperLogoSpecification
		SET Component = ?,
			Vendor = ?,
			ImageContent = ?,
			Material = ?,
			TableRow1 = ?,
			TableRow2 = ?,
			TableRow3 = ?,
			TableRow4 = ?,
			TableRow5 = ?,
			TableRow6 = ?,
			ItemIndex = ?
		OUTPUT CAST(INSERTED.ItemUpperLogoSpecificationID AS NVARCHAR(36)) AS ItemUpperLogoSpecificationID
		WHERE ItemUpperLogoSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.Vendor,
		req.ImageContent,
		req.Material,
		req.TableRow1,
		req.TableRow2,
		req.TableRow3,
		req.TableRow4,
		req.TableRow5,
		req.TableRow6,
		req.ItemIndex,
		req.ItemUpperLogoSpecificationID,
	).Scan(
		&ItemUpperLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemUpperLogoSpecificationID, nil
}

func (s *PFCUpperLogoSpecificationService) DeletePFCItemUpperLogoSpecification(req *types.PFC_ItemUpperLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemUpperLogoSpecification
		WHERE ItemUpperLogoSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ItemUpperLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
