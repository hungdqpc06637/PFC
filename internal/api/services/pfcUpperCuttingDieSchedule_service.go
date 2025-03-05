package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCUpperCuttingDieScheduleService struct {
	*BaseService
}

var PFCUpperCuttingDieSchedule = &PFCUpperCuttingDieScheduleService{}

func (s *PFCUpperCuttingDieScheduleService) GetAllPFCUpperCuttingDieSchedule(requestParams *types.PFCModel) (*[]types.PFC_UpperCuttingSchedule, error) {
	var upperCuttingDieSchedule *[]types.PFC_UpperCuttingSchedule
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT 
		CAST(UpperCuttingDieScheduleID AS NVARCHAR(36)) AS UpperCuttingDieScheduleID,
		ModelType,
		ModelName,
		MaterialNumber,
		ComponentName,
		ImageRemark,
		Remark,
		PageIndex
	FROM PFC_UpperCuttingDieSchedule
	WHERE
		ModelType = @ModelType
		AND ModelName = @ModelName
		AND MaterialNumber = @MaterialNumber
	ORDER BY PageIndex ASC
`
	err = db.Raw(query,
		sql.Named("ModelType", requestParams.ModelType),
		sql.Named("ModelName", requestParams.ModelName),
		sql.Named("MaterialNumber", requestParams.MaterialNumber),
	).Scan(&upperCuttingDieSchedule).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()
	// res, _ := PFCItemUpperCuttingDieSchedule.GetAllItemUpperCuttingDieSchedule(upperCuttingDieSchedule.UpperCuttingDieScheduleID)

	// upperCuttingDieSchedule.ItemsUpperCuttingDieScheduleID = res

	return upperCuttingDieSchedule, nil
}

func (s *PFCUpperCuttingDieScheduleService) InsertNewPFCUpperCuttingDieSchedule(req *types.PFC_UpperCuttingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var UpperCuttingDieScheduleID string

	query := `
		INSERT INTO PFC_UpperCuttingDieSchedule(UpperCuttingDieScheduleID, ModelType, ModelName, MaterialNumber, ComponentName, ImageRemark, Remark, PageIndex)
		OUTPUT CAST(INSERTED.UpperCuttingDieScheduleID AS NVARCHAR(36)) AS UpperCuttingDieScheduleID
		VALUES
		(NEWID(), ?, ?, ?, ?, ?, ?, ?)
	`
	if err := tx.Raw(
		query,
		req.ModelType,
		req.ModelName,
		req.MaterialNumber,
		req.ComponentName,
		req.ImageRemark,
		req.Remark,
		req.PageIndex,
	).Scan(
		&UpperCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return UpperCuttingDieScheduleID, nil
}

func (s *PFCUpperCuttingDieScheduleService) UpdatePFCUpperCuttingDieSchedule(req *types.PFC_UpperCuttingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var UpperCuttingDieScheduleID string

	query := `
		UPDATE PFC_UpperCuttingDieSchedule
		SET 
			ComponentName = ?, 
			ImageRemark = ?, 
			Remark = ?, 
			PageIndex = ?
		OUTPUT CAST(INSERTED.UpperCuttingDieScheduleID AS NVARCHAR(36)) AS UpperCuttingDieScheduleID
		WHERE 
			UpperCuttingDieScheduleID = ?
	`
	if err := tx.Raw(
		query,
		req.ComponentName,
		req.ImageRemark,
		req.Remark,
		req.PageIndex,
		req.UpperCuttingDieScheduleID,
	).Scan(
		&UpperCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return UpperCuttingDieScheduleID, nil
}

func (s *PFCUpperCuttingDieScheduleService) DeletePFCUpperCuttingDieSchedule(req *types.PFC_UpperCuttingSchedule) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	var UpperCuttingDieScheduleID string

	query := `
		DELETE PFC_UpperCuttingDieSchedule
		OUTPUT CAST(DELETED.UpperCuttingDieScheduleID AS NVARCHAR(36)) AS UpperCuttingDieScheduleID
		WHERE UpperCuttingDieScheduleID = ?
	`
	if err := tx.Raw(
		query,
		req.UpperCuttingDieScheduleID,
	).Scan(
		&UpperCuttingDieScheduleID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return UpperCuttingDieScheduleID, nil
}

func (s *PFCUpperCuttingDieScheduleService) DeletePFCUpperCuttingDieScheduleByModelID(req *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_UpperCuttingDieSchedule
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
