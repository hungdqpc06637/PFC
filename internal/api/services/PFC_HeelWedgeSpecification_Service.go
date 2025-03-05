package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCHeelWedgeSpecification struct {
	*BaseService
}

var PFCHeelWedgeSpecificatio = &PFCHeelWedgeSpecification{}

func (s *PFCHeelWedgeSpecification) GetAllPFCHeelWedgeSpecification(pfcModel *types.PFCModel) (*[]types.PFC_HeelWedgeSpecification, error) {
	var arrHeelWedgeSpecification *[]types.PFC_HeelWedgeSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(HeelWedgeSpecificationID AS NVARCHAR(36)) AS HeelWedgeSpecificationID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_HeelWedgeSpecification
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
	).Scan(&arrHeelWedgeSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrHeelWedgeSpecification, nil
}

func (s *PFCHeelWedgeSpecification) InsertNewPFCHeelWedgeSpecification(req *types.PFC_HeelWedgeSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var HeelWedgeSpecificationID string

	query := `
		INSERT INTO PFC_HeelWedgeSpecification
		(
		HeelWedgeSpecificationID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.HeelWedgeSpecificationID AS NVARCHAR(36)) AS HeelWedgeSpecificationID
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
		&HeelWedgeSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return HeelWedgeSpecificationID, nil
}

func (s *PFCHeelWedgeSpecification) UpdatePFCHeelWedgeSpecification(req *types.PFC_HeelWedgeSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var HeelWedgeSpecificationID string

	query := `
		UPDATE PFC_HeelWedgeSpecification
		SET Title = ?
		OUTPUT CAST(INSERTED.HeelWedgeSpecificationID AS NVARCHAR(36))
		WHERE HeelWedgeSpecificationID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.HeelWedgeSpecificationID,
	).Scan(
		&HeelWedgeSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return HeelWedgeSpecificationID, nil
}

func (s *PFCHeelWedgeSpecification) DeletePFCHeelWedgeSpecification(req *types.PFC_HeelWedgeSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_HeelWedgeSpecification
		WHERE HeelWedgeSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.HeelWedgeSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemHeelWedgeSpecification

func (s *PFCHeelWedgeSpecification) InsertNewPFCItemHeelWedgeSpecification(req *types.PFC_ItemHeelWedgeSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemHeelWedgeSpecificationID string

	query := `
		INSERT INTO PFC_ItemHeelWedgeSpecification
		(
		ItemHeelWedgeSpecificationID, HeelWedgeSpecificationID ,TableRow1, Thickness, ImagesContent
		)
		OUTPUT CAST(INSERTED.ItemHeelWedgeSpecificationID AS NVARCHAR(36)) AS ItemHeelWedgeSpecificationID
		VALUES (NEWID(), ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.HeelWedgeSpecificationID,
		req.TableRow1,
		req.Thickness,
		req.ImagesContent,
	).Scan(
		&ItemHeelWedgeSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemHeelWedgeSpecificationID, nil
}

func (s *PFCHeelWedgeSpecification) GetAllPFCItemHeelWedgeSpecification(pfcHeelWedgeSpecification *types.PFC_HeelWedgeSpecification) (*[]types.PFC_ItemHeelWedgeSpecification, error) {
	var arrItemHeelWedgeSpecification *[]types.PFC_ItemHeelWedgeSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemHeelWedgeSpecificationID AS NVARCHAR(36)) AS ItemHeelWedgeSpecificationID,
		CAST(HeelWedgeSpecificationID AS NVARCHAR(36)) AS HeelWedgeSpecificationID,
		 	TableRow1,Thickness , ImagesContent
	FROM PFC_ItemHeelWedgeSpecification
	WHERE HeelWedgeSpecificationID = @HeelWedgeSpecificationID
	`
	err = db.Raw(query,
		sql.Named("HeelWedgeSpecificationID", pfcHeelWedgeSpecification.HeelWedgeSpecificationID),
	).Scan(&arrItemHeelWedgeSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemHeelWedgeSpecification, nil
}

func (s *PFCHeelWedgeSpecification) UpdatePFCItemHeelWedgeSpecification(req *types.PFC_ItemHeelWedgeSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemHeelWedgeSpecificationID string

	query := `
		UPDATE PFC_ItemHeelWedgeSpecification
		SET 
			TableRow1 = ?,Thickness  = ?, ImagesContent = ?
		OUTPUT CAST(INSERTED.ItemHeelWedgeSpecificationID AS NVARCHAR(36)) AS ItemHeelWedgeSpecificationID
		WHERE ItemHeelWedgeSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.TableRow1,
		req.Thickness,
		req.ImagesContent,
		req.ItemHeelWedgeSpecificationID,
	).Scan(
		&ItemHeelWedgeSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemHeelWedgeSpecificationID, nil
}

func (s *PFCHeelWedgeSpecification) DeletePFCItemHeelWedgeSpecification(req *types.PFC_ItemHeelWedgeSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemHeelWedgeSpecification
		WHERE ItemHeelWedgeSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ItemHeelWedgeSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
