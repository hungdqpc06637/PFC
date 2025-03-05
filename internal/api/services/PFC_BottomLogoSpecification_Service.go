package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCBottomLogoSpecification struct {
	*BaseService
}

var PFCBottomLogoSpecificatio = &PFCBottomLogoSpecification{}

func (s *PFCBottomLogoSpecification) GetAllPFCBottomLogoSpecification(pfcModel *types.PFCModel) (*[]types.PFC_BottomLogoSpecification, error) {
	var arrBottomLogoSpecification *[]types.PFC_BottomLogoSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(BottomLogoSpecificationID AS NVARCHAR(36)) AS BottomLogoSpecificationID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_BottomLogoSpecification
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
	).Scan(&arrBottomLogoSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrBottomLogoSpecification, nil
}

func (s *PFCBottomLogoSpecification) InsertNewPFCBottomLogoSpecification(req *types.PFC_BottomLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var BottomLogoSpecificationID string

	query := `
		INSERT INTO PFC_BottomLogoSpecification(BottomLogoSpecificationID, ModelType, ModelName, MaterialNumber, Title,ItemIndex)
		OUTPUT CAST(INSERTED.BottomLogoSpecificationID AS NVARCHAR(36)) AS BottomLogoSpecificationID
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
		&BottomLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return BottomLogoSpecificationID, nil
}

func (s *PFCBottomLogoSpecification) UpdatePFCBottomLogoSpecification(req *types.PFC_BottomLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var BottomLogoSpecificationID string

	query := `
		UPDATE PFC_BottomLogoSpecification
		SET Title = ?
		OUTPUT CAST(INSERTED.BottomLogoSpecificationID AS NVARCHAR(36))
		WHERE BottomLogoSpecificationID = ?;
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.BottomLogoSpecificationID,
	).Scan(
		&BottomLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return BottomLogoSpecificationID, nil
}

func (s *PFCBottomLogoSpecification) DeletePFCBottomLogoSpecification(req *types.PFC_BottomLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_BottomLogoSpecification
		WHERE BottomLogoSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.BottomLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_BottomLogoSpecification

func (s *PFCBottomLogoSpecification) InsertNewPFCItemBottomLogoSpecification(req *types.PFC_ItemBottomLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemBottomLogoSpecificationID string

	query := `
		INSERT INTO PFC_ItemBottomLogoSpecification
		(
		ItemBottomLogoSpecificationID, BottomLogoSpecificationID, Component, ImageContent, Vendor ,MaterialApplication,Model,Size,ItemIndex 
		)
		OUTPUT CAST(INSERTED.ItemBottomLogoSpecificationID AS NVARCHAR(36)) AS ItemBottomLogoSpecificationID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.BottomLogoSpecificationID,
		req.Component,
		req.ImageContent,
		req.Vendor,
		req.MaterialApplication,
		req.Model,
		req.Size,
		req.ItemIndex,
	).Scan(
		&ItemBottomLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemBottomLogoSpecificationID, nil
}

func (s *PFCBottomLogoSpecification) GetAllPFCItemBottomLogoSpecification(pfcBottomLogoSpecification *types.PFC_BottomLogoSpecification) (*[]types.PFC_ItemBottomLogoSpecification, error) {
	var arrItemBottomLogoSpecification *[]types.PFC_ItemBottomLogoSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemBottomLogoSpecificationID AS NVARCHAR(36)) AS ItemBottomLogoSpecificationID,
		CAST(BottomLogoSpecificationID AS NVARCHAR(36)) AS BottomLogoSpecificationID,
		Component, 
		ImageContent, 
		Vendor ,
		MaterialApplication,
		Model,
		Size,
		ItemIndex 
	FROM PFC_ItemBottomLogoSpecification
	WHERE BottomLogoSpecificationID = @BottomLogoSpecificationID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("BottomLogoSpecificationID", pfcBottomLogoSpecification.BottomLogoSpecificationID),
	).Scan(&arrItemBottomLogoSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemBottomLogoSpecification, nil
}

func (s *PFCBottomLogoSpecification) UpdatePFCItemBottomLogoSpecification(req *types.PFC_ItemBottomLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemBottomLogoSpecificationID string

	query := `
		UPDATE PFC_ItemBottomLogoSpecification
		SET Component = ?, 
		ImageContent = ?, 
		Vendor = ? ,
		MaterialApplication = ?,
		Model = ?,
		Size = ?,
		ItemIndex  = ?
		OUTPUT CAST(INSERTED.ItemBottomLogoSpecificationID AS NVARCHAR(36)) AS ItemBottomLogoSpecificationID
		WHERE ItemBottomLogoSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.Component,
		req.ImageContent,
		req.Vendor,
		req.MaterialApplication,
		req.Model,
		req.Size,
		req.ItemIndex,
		req.ItemBottomLogoSpecificationID,
	).Scan(
		&ItemBottomLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemBottomLogoSpecificationID, nil
}

func (s *PFCBottomLogoSpecification) DeletePFCItemBottomLogoSpecification(req *types.PFC_ItemBottomLogoSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemBottomLogoSpecification
		WHERE ItemBottomLogoSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ItemBottomLogoSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
