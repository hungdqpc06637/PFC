package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCRubberComponentSpecification struct {
	*BaseService
}

var PFCRubberComponentSpecificatio = &PFCRubberComponentSpecification{}

func (s *PFCRubberComponentSpecification) GetAllPFCRubberComponentSpecification(pfcModel *types.PFCModel) (*[]types.PFC_RubberComponentSpecification, error) {
	var arrRubberComponentSpecification *[]types.PFC_RubberComponentSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(RubberComponentSpecificationID AS NVARCHAR(36)) AS RubberComponentSpecificationID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_RubberComponentSpecification
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
	).Scan(&arrRubberComponentSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrRubberComponentSpecification, nil
}

func (s *PFCRubberComponentSpecification) InsertNewPFCRubberComponentSpecification(req *types.PFC_RubberComponentSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var RubberComponentSpecificationID string

	query := `
		INSERT INTO PFC_RubberComponentSpecification
		(
		RubberComponentSpecificationID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.RubberComponentSpecificationID AS NVARCHAR(36)) AS RubberComponentSpecificationID
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
		&RubberComponentSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return RubberComponentSpecificationID, nil
}

func (s *PFCRubberComponentSpecification) UpdatePFCRubberComponentSpecification(req *types.PFC_RubberComponentSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var RubberComponentSpecificationID string

	query := `
		UPDATE PFC_RubberComponentSpecification
		SET Title = ?
		OUTPUT CAST(INSERTED.RubberComponentSpecificationID AS NVARCHAR(36))
		WHERE RubberComponentSpecificationID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.RubberComponentSpecificationID,
	).Scan(
		&RubberComponentSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return RubberComponentSpecificationID, nil
}

func (s *PFCRubberComponentSpecification) DeletePFCRubberComponentSpecification(req *types.PFC_RubberComponentSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_RubberComponentSpecification
		WHERE RubberComponentSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.RubberComponentSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemRubberComponentSpecification

func (s *PFCRubberComponentSpecification) InsertNewPFCItemRubberComponentSpecification(req *types.PFC_ItemRubberComponentSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemRubberComponentSpecificationID string

	query := `
		INSERT INTO PFC_ItemRubberComponentSpecification
		(
				ItemRubberComponentSpecificationID ,RubberComponentSpecificationID, TableRow1, PartInformation,LengthGrading, TableRow2
		)
		OUTPUT CAST(INSERTED.ItemRubberComponentSpecificationID AS NVARCHAR(36)) AS ItemRubberComponentSpecificationID
		VALUES (NEWID(), ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.RubberComponentSpecificationID,
		req.TableRow1,
		req.PartInformation,
		req.LengthGrading,
		req.TableRow2,
	).Scan(
		&ItemRubberComponentSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemRubberComponentSpecificationID, nil
}

func (s *PFCRubberComponentSpecification) GetAllPFCItemRubberComponentSpecification(pfcRubberComponentSpecification *types.PFC_RubberComponentSpecification) (*[]types.PFC_ItemRubberComponentSpecification, error) {
	var arrItemRubberComponentSpecification *[]types.PFC_ItemRubberComponentSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemRubberComponentSpecificationID AS NVARCHAR(36)) AS ItemRubberComponentSpecificationID,
		CAST(RubberComponentSpecificationID AS NVARCHAR(36)) AS RubberComponentSpecificationID,
		 	TableRow1, PartInformation,LengthGrading,TableRow2
	FROM PFC_ItemRubberComponentSpecification
	WHERE RubberComponentSpecificationID = @RubberComponentSpecificationID
	`
	err = db.Raw(query,
		sql.Named("RubberComponentSpecificationID", pfcRubberComponentSpecification.RubberComponentSpecificationID),
	).Scan(&arrItemRubberComponentSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemRubberComponentSpecification, nil
}

func (s *PFCRubberComponentSpecification) UpdatePFCItemRubberComponentSpecification(req *types.PFC_ItemRubberComponentSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemRubberComponentSpecificationID string

	query := `
		UPDATE PFC_ItemRubberComponentSpecification
		SET 
			TableRow1 = ?, PartInformation =?, LengthGrading=?, TableRow2 = ? 
		OUTPUT CAST(INSERTED.ItemRubberComponentSpecificationID AS NVARCHAR(36)) AS ItemRubberComponentSpecificationID
		WHERE ItemRubberComponentSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.TableRow1,
		req.PartInformation,
		req.LengthGrading,
		req.TableRow2,
		req.ItemRubberComponentSpecificationID,
	).Scan(
		&ItemRubberComponentSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemRubberComponentSpecificationID, nil
}

func (s *PFCRubberComponentSpecification) DeletePFCItemRubberComponentSpecification(req *types.PFC_ItemRubberComponentSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemRubberComponentSpecification
		WHERE ItemRubberComponentSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ItemRubberComponentSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
