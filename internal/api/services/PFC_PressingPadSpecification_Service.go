package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCPressingPadSpecification struct {
	*BaseService
}

var PFCPressingPadSpecificatio = &PFCPressingPadSpecification{}

func (s *PFCPressingPadSpecification) GetAllPFCPressingPadSpecification(pfcModel *types.PFCModel) (*[]types.PFC_PressingPadSpecification, error) {
	var arrPressingPadSpecification *[]types.PFC_PressingPadSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(PressingPadSpecificationID AS NVARCHAR(36)) AS PressingPadSpecificationID,
		ModelType,
		ModelName,
		MaterialNumber,

		ColorWayID ,
		TechLevel ,
		FirstSource ,

		Title,
		ItemIndex
	FROM PFC_PressingPadSpecification
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
	).Scan(&arrPressingPadSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrPressingPadSpecification, nil
}

func (s *PFCPressingPadSpecification) InsertNewPFCPressingPadSpecification(req *types.PFC_PressingPadSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var PressingPadSpecificationID string

	query := `
		INSERT INTO PFC_PressingPadSpecification
		(
		PressingPadSpecificationID , ModelType, ModelName, MaterialNumber,ColorWayID , TechLevel , FirstSource , Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.PressingPadSpecificationID AS NVARCHAR(36)) AS PressingPadSpecificationID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?, ?, ?)
		`
	if err := tx.Raw(
		query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,

		req.ColorWayID,
		req.TechLevel,
		req.FirstSource,

		req.Title,
		req.ItemIndex,
	).Scan(
		&PressingPadSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return PressingPadSpecificationID, nil
}

func (s *PFCPressingPadSpecification) UpdatePFCPressingPadSpecification(req *types.PFC_PressingPadSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var PressingPadSpecificationID string

	query := `
		UPDATE PFC_PressingPadSpecification
		SET Title = ? , ColorWayID = ? , TechLevel = ? , FirstSource = ? 
		OUTPUT CAST(INSERTED.PressingPadSpecificationID AS NVARCHAR(36))
		WHERE PressingPadSpecificationID = ?; 
		`

	if err := tx.Raw(
		query,
		req.Title,
		req.ColorWayID,
		req.TechLevel,
		req.FirstSource,
		req.PressingPadSpecificationID,
	).Scan(
		&PressingPadSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return PressingPadSpecificationID, nil
}

func (s *PFCPressingPadSpecification) DeletePFCPressingPadSpecification(req *types.PFC_PressingPadSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_PressingPadSpecification
		WHERE PressingPadSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.PressingPadSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemPressingPadSpecification

func (s *PFCPressingPadSpecification) InsertNewPFCItemPressingPadSpecification(req *types.PFC_ItemPressingPadSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemPressingPadSpecificationID string

	query := `
		INSERT INTO PFC_ItemPressingPadSpecification
		(
		ItemPressingPadSpecificationID ,PressingPadSpecificationID, TableRow1 ,TableRow2 
		)
		OUTPUT CAST(INSERTED.ItemPressingPadSpecificationID AS NVARCHAR(36)) AS ItemPressingPadSpecificationID
		VALUES (NEWID(), ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.PressingPadSpecificationID,
		req.TableRow1,
		req.TableRow2,
	).Scan(
		&ItemPressingPadSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemPressingPadSpecificationID, nil
}

func (s *PFCPressingPadSpecification) GetAllPFCItemPressingPadSpecification(pfcPressingPadSpecification *types.PFC_PressingPadSpecification) (*[]types.PFC_ItemPressingPadSpecification, error) {
	var arrItemPressingPadSpecification *[]types.PFC_ItemPressingPadSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemPressingPadSpecificationID AS NVARCHAR(36)) AS ItemPressingPadSpecificationID,
		CAST(PressingPadSpecificationID AS NVARCHAR(36)) AS PressingPadSpecificationID,
		 	TableRow1 ,TableRow2 
	FROM PFC_ItemPressingPadSpecification
	WHERE PressingPadSpecificationID = @PressingPadSpecificationID
	`
	err = db.Raw(query,
		sql.Named("PressingPadSpecificationID", pfcPressingPadSpecification.PressingPadSpecificationID),
	).Scan(&arrItemPressingPadSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemPressingPadSpecification, nil
}

func (s *PFCPressingPadSpecification) UpdatePFCItemPressingPadSpecification(req *types.PFC_ItemPressingPadSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemPressingPadSpecificationID string

	query := `
		UPDATE PFC_ItemPressingPadSpecification
		SET 
			TableRow1 = ?  ,TableRow2 = ?  
		OUTPUT CAST(INSERTED.ItemPressingPadSpecificationID AS NVARCHAR(36)) AS ItemPressingPadSpecificationID
		WHERE ItemPressingPadSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.TableRow1,
		req.TableRow2,
		req.ItemPressingPadSpecificationID,
	).Scan(
		&ItemPressingPadSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemPressingPadSpecificationID, nil
}

func (s *PFCPressingPadSpecification) DeletePFCItemPressingPadSpecification(req *types.PFC_ItemPressingPadSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemPressingPadSpecification
		WHERE ItemPressingPadSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ItemPressingPadSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
