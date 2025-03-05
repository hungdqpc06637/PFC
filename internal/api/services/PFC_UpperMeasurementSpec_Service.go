package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCUpperMeasurementSpec struct {
	*BaseService
}

var PFCUpperMeasurementSpe = &PFCUpperMeasurementSpec{}

func (s *PFCUpperMeasurementSpec) GetAllPFCUpperMeasurementSpec(pfcModel *types.PFCModel) (*[]types.PFC_UpperMeasurementSpec, error) {
	var arrUpperMeasurementSpec *[]types.PFC_UpperMeasurementSpec
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(UpperMeasurementSpecID AS NVARCHAR(36)) AS UpperMeasurementSpecID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_UpperMeasurementSpec
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
	).Scan(&arrUpperMeasurementSpec).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrUpperMeasurementSpec, nil
}

func (s *PFCUpperMeasurementSpec) InsertNewPFCUpperMeasurementSpec(req *types.PFC_UpperMeasurementSpec) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var UpperMeasurementSpecID string

	query := `
		INSERT INTO PFC_UpperMeasurementSpec
		(
		UpperMeasurementSpecID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.UpperMeasurementSpecID AS NVARCHAR(36)) AS UpperMeasurementSpecID
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
		&UpperMeasurementSpecID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return UpperMeasurementSpecID, nil
}

func (s *PFCUpperMeasurementSpec) UpdatePFCUpperMeasurementSpec(req *types.PFC_UpperMeasurementSpec) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var UpperMeasurementSpecID string

	query := `
		UPDATE PFC_UpperMeasurementSpec
		SET Title = ?
		OUTPUT CAST(INSERTED.UpperMeasurementSpecID AS NVARCHAR(36))
		WHERE UpperMeasurementSpecID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.UpperMeasurementSpecID,
	).Scan(
		&UpperMeasurementSpecID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return UpperMeasurementSpecID, nil
}

func (s *PFCUpperMeasurementSpec) DeletePFCUpperMeasurementSpec(req *types.PFC_UpperMeasurementSpec) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_UpperMeasurementSpec
		WHERE UpperMeasurementSpecID = ?
	`

	if err := tx.Exec(query,
		req.UpperMeasurementSpecID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemUpperMeasurementSpec

func (s *PFCUpperMeasurementSpec) InsertNewPFCItemUpperMeasurementSpec(req *types.PFC_ItemUpperMeasurementSpec) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemUpperMeasurementSpecID string

	query := `
		INSERT INTO PFC_ItemUpperMeasurementSpec
		(
		ItemUpperMeasurementSpecID ,UpperMeasurementSpecID, ImagesContent,TableRow1 
		)
		OUTPUT CAST(INSERTED.ItemUpperMeasurementSpecID AS NVARCHAR(36)) AS ItemUpperMeasurementSpecID
		VALUES (NEWID(), ?, ? ,?)

	`
	if err := tx.Raw(
		query,
		req.UpperMeasurementSpecID,
		req.ImagesContent,
		req.TableRow1,
	).Scan(
		&ItemUpperMeasurementSpecID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemUpperMeasurementSpecID, nil
}

func (s *PFCUpperMeasurementSpec) GetAllPFCItemUpperMeasurementSpec(pfcUpperMeasurementSpec *types.PFC_UpperMeasurementSpec) (*[]types.PFC_ItemUpperMeasurementSpec, error) {
	var arrItemUpperMeasurementSpec *[]types.PFC_ItemUpperMeasurementSpec
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemUpperMeasurementSpecID AS NVARCHAR(36)) AS ItemUpperMeasurementSpecID,
		CAST(UpperMeasurementSpecID AS NVARCHAR(36)) AS UpperMeasurementSpecID,
		 	ImagesContent,TableRow1 
	FROM PFC_ItemUpperMeasurementSpec
	WHERE UpperMeasurementSpecID = @UpperMeasurementSpecID
	`
	err = db.Raw(query,
		sql.Named("UpperMeasurementSpecID", pfcUpperMeasurementSpec.UpperMeasurementSpecID),
	).Scan(&arrItemUpperMeasurementSpec).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemUpperMeasurementSpec, nil
}

func (s *PFCUpperMeasurementSpec) UpdatePFCItemUpperMeasurementSpec(req *types.PFC_ItemUpperMeasurementSpec) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemUpperMeasurementSpecID string

	query := `
		UPDATE PFC_ItemUpperMeasurementSpec
		SET  ImagesContent = ? ,TableRow1  = ? 
		OUTPUT CAST(INSERTED.ItemUpperMeasurementSpecID AS NVARCHAR(36)) AS ItemUpperMeasurementSpecID
		WHERE ItemUpperMeasurementSpecID = ?
	`
	if err := tx.Raw(
		query,
		req.ImagesContent,
		req.TableRow1,
		req.ItemUpperMeasurementSpecID,
	).Scan(
		&ItemUpperMeasurementSpecID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemUpperMeasurementSpecID, nil
}

func (s *PFCUpperMeasurementSpec) DeletePFCItemUpperMeasurementSpec(req *types.PFC_ItemUpperMeasurementSpec) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemUpperMeasurementSpec
		WHERE ItemUpperMeasurementSpecID = ?
	`

	if err := tx.Exec(query,
		req.ItemUpperMeasurementSpecID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
