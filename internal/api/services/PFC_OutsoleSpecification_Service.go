package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCOutsoleSpecification struct {
	*BaseService
}

var PFCOutsoleSpecificatio = &PFCOutsoleSpecification{}

func (s *PFCOutsoleSpecification) GetAllPFCOutsoleSpecification(pfcModel *types.PFCModel) (*[]types.PFC_OutsoleSpecification, error) {
	var arrOutsoleSpecification *[]types.PFC_OutsoleSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(OutsoleSpecificationID AS NVARCHAR(36)) AS OutsoleSpecificationID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_OutsoleSpecification
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
	).Scan(&arrOutsoleSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrOutsoleSpecification, nil
}

func (s *PFCOutsoleSpecification) InsertNewPFCOutsoleSpecification(req *types.PFC_OutsoleSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var OutsoleSpecificationID string

	query := `
		INSERT INTO PFC_OutsoleSpecification
		(
		OutsoleSpecificationID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.OutsoleSpecificationID AS NVARCHAR(36)) AS OutsoleSpecificationID
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
		&OutsoleSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return OutsoleSpecificationID, nil
}

func (s *PFCOutsoleSpecification) UpdatePFCOutsoleSpecification(req *types.PFC_OutsoleSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var OutsoleSpecificationID string

	query := `
		UPDATE PFC_OutsoleSpecification
		SET Title = ?
		OUTPUT CAST(INSERTED.OutsoleSpecificationID AS NVARCHAR(36))
		WHERE OutsoleSpecificationID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.OutsoleSpecificationID,
	).Scan(
		&OutsoleSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return OutsoleSpecificationID, nil
}

func (s *PFCOutsoleSpecification) DeletePFCOutsoleSpecification(req *types.PFC_OutsoleSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_OutsoleSpecification
		WHERE OutsoleSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.OutsoleSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemOutsoleSpecification

func (s *PFCOutsoleSpecification) InsertNewPFCItemOutsoleSpecification(req *types.PFC_ItemOutsoleSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemOutsoleSpecificationID string

	query := `
		INSERT INTO PFC_ItemOutsoleSpecification
		(
				ItemOutsoleSpecificationID ,OutsoleSpecificationID, TableRow1, PreformPartsNestingDiagram, FinalPart, TableRow2, TotalWeightPerSize
		)
		OUTPUT CAST(INSERTED.ItemOutsoleSpecificationID AS NVARCHAR(36)) AS ItemOutsoleSpecificationID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.OutsoleSpecificationID,
		req.TableRow1,
		req.PreformPartsNestingDiagram,
		req.FinalPart,
		req.TableRow2,
		req.TotalWeightPerSize,
	).Scan(
		&ItemOutsoleSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemOutsoleSpecificationID, nil
}

func (s *PFCOutsoleSpecification) GetAllPFCItemOutsoleSpecification(pfcOutsoleSpecification *types.PFC_OutsoleSpecification) (*[]types.PFC_ItemOutsoleSpecification, error) {
	var arrItemOutsoleSpecification *[]types.PFC_ItemOutsoleSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemOutsoleSpecificationID AS NVARCHAR(36)) AS ItemOutsoleSpecificationID,
		CAST(OutsoleSpecificationID AS NVARCHAR(36)) AS OutsoleSpecificationID,
		 	TableRow1, PreformPartsNestingDiagram, FinalPart,TableRow2,TotalWeightPerSize
	FROM PFC_ItemOutsoleSpecification
	WHERE OutsoleSpecificationID = @OutsoleSpecificationID
	`
	err = db.Raw(query,
		sql.Named("OutsoleSpecificationID", pfcOutsoleSpecification.OutsoleSpecificationID),
	).Scan(&arrItemOutsoleSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemOutsoleSpecification, nil
}

func (s *PFCOutsoleSpecification) UpdatePFCItemOutsoleSpecification(req *types.PFC_ItemOutsoleSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemOutsoleSpecificationID string

	query := `
		UPDATE PFC_ItemOutsoleSpecification
		SET 
			TableRow1 = ?, PreformPartsNestingDiagram=?, FinalPart=?, TableRow2 = ?, TotalWeightPerSize = ?
		OUTPUT CAST(INSERTED.ItemOutsoleSpecificationID AS NVARCHAR(36)) AS ItemOutsoleSpecificationID
		WHERE ItemOutsoleSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.TableRow1,
		req.PreformPartsNestingDiagram,
		req.FinalPart,
		req.TableRow2,
		req.TotalWeightPerSize,
		req.ItemOutsoleSpecificationID,
	).Scan(
		&ItemOutsoleSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemOutsoleSpecificationID, nil
}

func (s *PFCOutsoleSpecification) DeletePFCItemOutsoleSpecification(req *types.PFC_ItemOutsoleSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemOutsoleSpecification
		WHERE ItemOutsoleSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ItemOutsoleSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
