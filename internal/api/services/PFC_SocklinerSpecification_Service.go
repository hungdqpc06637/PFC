package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCSocklinerSpecification struct {
	*BaseService
}

var PFCSocklinerSpecificatio = &PFCSocklinerSpecification{}

func (s *PFCSocklinerSpecification) GetAllPFCSocklinerSpecification(pfcModel *types.PFCModel) (*[]types.PFC_SocklinerSpecification, error) {
	var arrSocklinerSpecification *[]types.PFC_SocklinerSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(SocklinerSpecificationID AS NVARCHAR(36)) AS SocklinerSpecificationID,
		ModelType,
		ModelName,
		MaterialNumber,
		Title,
		ItemIndex
	FROM PFC_SocklinerSpecification
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
	).Scan(&arrSocklinerSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrSocklinerSpecification, nil
}

func (s *PFCSocklinerSpecification) InsertNewPFCSocklinerSpecification(req *types.PFC_SocklinerSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SocklinerSpecificationID string

	query := `
		INSERT INTO PFC_SocklinerSpecification
		(
		SocklinerSpecificationID , ModelType, ModelName, MaterialNumber, Title,  ItemIndex
		)
		OUTPUT CAST(INSERTED.SocklinerSpecificationID AS NVARCHAR(36)) AS SocklinerSpecificationID
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
		&SocklinerSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SocklinerSpecificationID, nil
}

func (s *PFCSocklinerSpecification) UpdatePFCSocklinerSpecification(req *types.PFC_SocklinerSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var SocklinerSpecificationID string

	query := `
		UPDATE PFC_SocklinerSpecification
		SET Title = ?
		OUTPUT CAST(INSERTED.SocklinerSpecificationID AS NVARCHAR(36))
		WHERE SocklinerSpecificationID = ?;

		`

	if err := tx.Raw(
		query,
		req.Title,
		req.SocklinerSpecificationID,
	).Scan(
		&SocklinerSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return SocklinerSpecificationID, nil
}

func (s *PFCSocklinerSpecification) DeletePFCSocklinerSpecification(req *types.PFC_SocklinerSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_SocklinerSpecification
		WHERE SocklinerSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.SocklinerSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

// // ITEM PFC_ItemSocklinerSpecification

func (s *PFCSocklinerSpecification) InsertNewPFCItemSocklinerSpecification(req *types.PFC_ItemSocklinerSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSocklinerSpecificationID string

	query := `
		INSERT INTO PFC_ItemSocklinerSpecification
		(
		ItemSocklinerSpecificationID, SocklinerSpecificationID,Front,Back,Size,TableRow1, TableRow2
		)
		OUTPUT CAST(INSERTED.ItemSocklinerSpecificationID AS NVARCHAR(36)) AS ItemSocklinerSpecificationID
		VALUES (NEWID(), ?, ?, ?, ?, ?, ?)

	`
	if err := tx.Raw(
		query,
		req.SocklinerSpecificationID,
		req.Front,
		req.Back,
		req.Size,
		req.TableRow1,
		req.TableRow2,
	).Scan(
		&ItemSocklinerSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSocklinerSpecificationID, nil
}

func (s *PFCSocklinerSpecification) GetAllPFCItemSocklinerSpecification(pfcSocklinerSpecification *types.PFC_SocklinerSpecification) (*[]types.PFC_ItemSocklinerSpecification, error) {
	var arrItemSocklinerSpecification *[]types.PFC_ItemSocklinerSpecification
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemSocklinerSpecificationID AS NVARCHAR(36)) AS ItemSocklinerSpecificationID,
		CAST(SocklinerSpecificationID AS NVARCHAR(36)) AS SocklinerSpecificationID,
		 	Front,Back,Size,TableRow1, TableRow2
	FROM PFC_ItemSocklinerSpecification
	WHERE SocklinerSpecificationID = @SocklinerSpecificationID
	`
	err = db.Raw(query,
		sql.Named("SocklinerSpecificationID", pfcSocklinerSpecification.SocklinerSpecificationID),
	).Scan(&arrItemSocklinerSpecification).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return arrItemSocklinerSpecification, nil
}

func (s *PFCSocklinerSpecification) UpdatePFCItemSocklinerSpecification(req *types.PFC_ItemSocklinerSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var ItemSocklinerSpecificationID string

	query := `
		UPDATE PFC_ItemSocklinerSpecification
		SET 
			Front = ?,Back = ?,Size = ?,TableRow1 = ?, TableRow2 = ?
		OUTPUT CAST(INSERTED.ItemSocklinerSpecificationID AS NVARCHAR(36)) AS ItemSocklinerSpecificationID
		WHERE ItemSocklinerSpecificationID = ?
	`
	if err := tx.Raw(
		query,
		req.Front,
		req.Back,
		req.Size,
		req.TableRow1,
		req.TableRow2,
		req.ItemSocklinerSpecificationID,
	).Scan(
		&ItemSocklinerSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return ItemSocklinerSpecificationID, nil
}

func (s *PFCSocklinerSpecification) DeletePFCItemSocklinerSpecification(req *types.PFC_ItemSocklinerSpecification) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemSocklinerSpecification
		WHERE ItemSocklinerSpecificationID = ?
	`

	if err := tx.Exec(query,
		req.ItemSocklinerSpecificationID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
