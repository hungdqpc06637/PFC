package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCItemUpperCuttingDieScheduleService struct {
	*BaseService
}

var PFCItemUpperCuttingDieSchedule = &PFCItemUpperCuttingDieScheduleService{}

func (s *PFCItemUpperCuttingDieScheduleService) GetAllItemUpperCuttingDieScheduleByModelID(req *types.PFCModel) ([]types.PFC_ItemUpperCuttingSchedule, error) {
	var list []types.PFC_ItemUpperCuttingSchedule
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(i.ItemUpperCuttingDieScheduleID AS NVARCHAR(36)) AS ItemUpperCuttingDieScheduleID,
		CAST(i.UpperCuttingDieScheduleID AS NVARCHAR(36)) AS UpperCuttingDieScheduleID,
		i.ComponentName,
		i.ItemIndex,
		i.SizeRange1,
		i.SizeRange2,
		i.SizeRange,
		i.SizeRangeAreSame,
		i.ImageContent,
		i.ImageRemark,
		i.NumberOfPlayers
	FROM PFC_ItemUpperCuttingDieSchedule i
	INNER JOIN PFC_UpperCuttingDieSchedule u
    ON i.UpperCuttingDieScheduleID = u.UpperCuttingDieScheduleID
	WHERE  u.ModelType = @ModelType AND u.ModelName = @ModelName AND u.MaterialNumber = @MaterialNumber
	ORDER BY i.ItemIndex ASC
`
	err = db.Raw(query,
		sql.Named("ModelType", req.ModelType),
		sql.Named("ModelName", req.ModelName),
		sql.Named("MaterialNumber", req.MaterialNumber),
	).Scan(&list).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return list, nil
}

func (s *PFCItemUpperCuttingDieScheduleService) GetAllItemUpperCuttingDieSchedule(req *types.PFC_UpperCuttingSchedule) ([]types.PFC_ItemUpperCuttingSchedule, error) {
	var list []types.PFC_ItemUpperCuttingSchedule
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(ItemUpperCuttingDieScheduleID AS NVARCHAR(36)) AS ItemUpperCuttingDieScheduleID,
		CAST(UpperCuttingDieScheduleID AS NVARCHAR(36)) AS UpperCuttingDieScheduleID,
		ComponentName,
		ItemIndex,
		SizeRange1,
		SizeRange2,
		SizeRange,
		SizeRangeAreSame,
		ImageContent,
		ImageRemark,
		NumberOfPlayers
	FROM PFC_ItemUpperCuttingDieSchedule
	WHERE UpperCuttingDieScheduleID = @UpperCuttingDieScheduleID
	ORDER BY ItemIndex ASC
`
	err = db.Raw(query, sql.Named("UpperCuttingDieScheduleID", req.UpperCuttingDieScheduleID)).Scan(&list).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return list, nil
}

func (s *PFCItemUpperCuttingDieScheduleService) InsertNewPFCItemUpperCuttingDieSchedule(req *types.PFC_ItemUpperCuttingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		INSERT INTO PFC_ItemUpperCuttingDieSchedule(ItemUpperCuttingDieScheduleID, UpperCuttingDieScheduleID, ComponentName, 
													ItemIndex, SizeRange1, SizeRange2, 
													SizeRange, ImageContent, ImageRemark, 
													NumberOfPlayers, SizeRangeAreSame)
		VALUES
		(NEWID(), ?, ?,
		?, ?, ?,
		?, ?, ?,
		?, ?)
	`
	if err := tx.Exec(
		query,
		req.UpperCuttingDieScheduleID,
		req.ComponentName,
		req.ItemIndex,
		req.SizeRange1,
		req.SizeRange2,
		req.SizeRange,
		req.ImageContent,
		req.ImageRemark,
		req.NumberOfPlayers,
		req.SizeRangeAreSame,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "success", nil
}

func (s *PFCItemUpperCuttingDieScheduleService) UpdatePFCItemUpperCuttingDieSchedule(req *types.PFC_ItemUpperCuttingSchedule) (*types.PFC_ItemUpperCuttingSchedule, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		UPDATE PFC_ItemUpperCuttingDieSchedule
		SET 
			ComponentName = ?,
			ItemIndex = ?,
			SizeRange1 = ?,
			SizeRange2 = ?,
			SizeRange = ?,
			SizeRangeAreSame = ?,
			ImageContent = ?,
			ImageRemark = ?,
			NumberOfPlayers = ?
		WHERE
			ItemUpperCuttingDieScheduleID = ?
	`
	if err := tx.Exec(query,
		req.ComponentName,
		req.ItemIndex,
		req.SizeRange1,
		req.SizeRange2,
		req.SizeRange,
		req.SizeRangeAreSame,
		req.ImageContent,
		req.ImageRemark,
		req.NumberOfPlayers,
		req.ItemUpperCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to execute update query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return req, nil
}

func (s *PFCItemUpperCuttingDieScheduleService) DeletePFCItemUpperCuttingDieSchedule(req *types.PFC_ItemUpperCuttingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemUpperCuttingDieSchedule
		WHERE ItemUpperCuttingDieScheduleID = ?
	`

	if err := tx.Exec(query,
		req.ItemUpperCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

func (s *PFCItemUpperCuttingDieScheduleService) DeletePFCItemUpperCuttingDieScheduleByModelID(req *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_ItemUpperCuttingDieSchedule
		WHERE UpperCuttingDieScheduleID = (SELECT TOP(1) UpperCuttingDieScheduleID FROM PFC_UpperCuttingDieSchedule WHERE ModelType = ? AND ModelName = ? AND MaterialNumber = ?)
	`

	if err := tx.Exec(query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}
