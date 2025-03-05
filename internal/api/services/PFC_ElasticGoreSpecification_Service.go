package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCElasticGoreSpecificationService struct {
	*BaseService
}

var PFCElasticGoreSpecification = &PFCElasticGoreSpecificationService{}

func (s *PFCElasticGoreSpecificationService) GetAllPFCElasticGoreSpecification(pfcModel *types.PFCModel) (*[]types.PFC_ElasticGoreSpecification, error) {
	var arrElasticGoreSpecification *[]types.PFC_ElasticGoreSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(ElasticGoreSpecificationID AS NVARCHAR(36)) AS ElasticGoreSpecificationID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_ElasticGoreSpecification
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
	).Scan(&arrElasticGoreSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrElasticGoreSpecification, nil
}

func (s *PFCElasticGoreSpecificationService) InsertNewPFCElasticGoreSpecification(req *types.PFC_ElasticGoreSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var elasticGoreSpecificationID string

	query := `
		INSERT INTO PFC_ElasticGoreSpecification(ElasticGoreSpecificationID, ModelType, ModelName, MaterialNumber, Title, ItemIndex)
		OUTPUT CAST(INSERTED.ElasticGoreSpecificationID AS NVARCHAR(36)) AS ElasticGoreSpecificationID
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
		&elasticGoreSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return elasticGoreSpecificationID, nil
}

func (s *PFCElasticGoreSpecificationService) UpdatePFCElasticGoreSpecification(req *types.PFC_ElasticGoreSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var elasticGoreSpecificationID string

	query := `
		UPDATE PFC_ElasticGoreSpecification
		SET Title = ?
		OUTPUT CAST(INSERTED.ElasticGoreSpecificationID AS NVARCHAR(36)) AS ElasticGoreSpecificationID
		WHERE ElasticGoreSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.Title,
		req.ElasticGoreSpecificationID,
	).Scan(
		&elasticGoreSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return elasticGoreSpecificationID, nil
}

func (s *PFCElasticGoreSpecificationService) DeletePFCElasticGoreSpecification(req *types.PFC_ElasticGoreSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ElasticGoreSpecification
		WHERE ElasticGoreSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ElasticGoreSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// ITEM PFC_ElasticGoreSpecification
func (s *PFCElasticGoreSpecificationService) InsertNewPFCItemElasticGoreSpecification(req *types.PFC_ItemElasticGoreSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemElasticGoreSpecificationID string

	query := `
		INSERT INTO PFC_ItemElasticGoreSpecification(ItemElasticGoreSpecificationID, ElasticGoreSpecificationID, Component, Vendor, ImageContent, Material, Model, TableRow1, TableRow2, TableRow3, TableRow4, TableRow5, ItemIndex)
		OUTPUT CAST(INSERTED.ItemElasticGoreSpecificationID AS NVARCHAR(36)) AS ItemElasticGoreSpecificationID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.ElasticGoreSpecificationID,
		req.Component,
		req.Vendor,
		req.ImageContent,
		req.Material,
		req.Model,
		req.TableRow1,
		req.TableRow2,
		req.TableRow3,
		req.TableRow4,
		req.TableRow5,
		req.ItemIndex,
	).Scan(
		&ItemElasticGoreSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemElasticGoreSpecificationID, nil
}

func (s *PFCElasticGoreSpecificationService) GetAllPFCItemElasticGoreSpecification(pfcElasticGoreSpecification *types.PFC_ElasticGoreSpecification) (*[]types.PFC_ItemElasticGoreSpecification, error) {
	var arrItemElasticGoreSpecification *[]types.PFC_ItemElasticGoreSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemElasticGoreSpecificationID AS NVARCHAR(36)) AS ItemElasticGoreSpecificationID,
		CAST(ElasticGoreSpecificationID AS NVARCHAR(36)) AS ElasticGoreSpecificationID,
		Component,
		Vendor,
		ImageContent,
		Material,
		Model,
		TableRow1,
		TableRow2,
		TableRow3,
		TableRow4,
		TableRow5,
		ItemIndex
	FROM PFC_ItemElasticGoreSpecification
	WHERE ElasticGoreSpecificationID = @ElasticGoreSpecificationID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("ElasticGoreSpecificationID", pfcElasticGoreSpecification.ElasticGoreSpecificationID),
	).Scan(&arrItemElasticGoreSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemElasticGoreSpecification, nil
}

func (s *PFCElasticGoreSpecificationService) UpdatePFCItemElasticGoreSpecification(req *types.PFC_ItemElasticGoreSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemElasticGoreSpecificationID string

	query := `
		UPDATE PFC_ItemElasticGoreSpecification
		SET Component = ?,
			Vendor = ?,
			ImageContent = ?,
			Material = ?,
			Model = ?,
			TableRow1 = ?,
			TableRow2 = ?,
			TableRow3 = ?,
			TableRow4 = ?,
			TableRow5 = ?,
			ItemIndex = ?
		OUTPUT CAST(INSERTED.ItemElasticGoreSpecificationID AS NVARCHAR(36)) AS ItemElasticGoreSpecificationID
		WHERE ItemElasticGoreSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.Vendor,
		req.ImageContent,
		req.Material,
		req.Model,
		req.TableRow1,
		req.TableRow2,
		req.TableRow3,
		req.TableRow4,
		req.TableRow5,
		req.ItemIndex,
		req.ItemElasticGoreSpecificationID,
	).Scan(
		&ItemElasticGoreSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemElasticGoreSpecificationID, nil
}

func (s *PFCElasticGoreSpecificationService) DeletePFCItemElasticGoreSpecification(req *types.PFC_ItemElasticGoreSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemElasticGoreSpecification
		WHERE ItemElasticGoreSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ItemElasticGoreSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
