package services

import (
	"database/sql"
	"fmt"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/types"
)

type PFCRollService struct {
	*BaseService
}

var PFCRoll = &PFCRollService{}

func (s *PFCRollService) GetAllPFCRolls(laminationProcessID string) ([]types.PFC_Roll, error) {
	var list []types.PFC_Roll
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}
	dbInstance, _ := db.DB()
	query := `
	SELECT
		CAST(RollID AS NVARCHAR(36)) AS RollID, 
		CAST(LaminationProcessID AS NVARCHAR(36)) AS LaminationProcessID, 
		Temp, 
		RollSize, 
		Time
	FROM PFC_Roll
	WHERE LaminationProcessID = @LaminationProcessID
`
	err = db.Raw(query, sql.Named("LaminationProcessID", laminationProcessID)).Scan(&list).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	dbInstance.Close()

	return list, nil
}

func (s *PFCRollService) InsertNewPFCRoll(req *types.PFC_Roll) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		INSERT INTO PFC_Roll(RollID, LaminationProcessID, Temp, RollSize, Time)
		VALUES
		(NEWID(), ?, ?, ?, ?)
	`
	if err := tx.Exec(
		query,
		req.LaminationProcessID,
		req.Temp,
		req.RollSize,
		req.Time,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute query: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "success", nil
}

func (s *PFCRollService) UpdatePFCRoll(req *types.PFC_Roll) (*types.PFC_Roll, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		UPDATE PFC_Roll
		SET 
			Temp = ?, 
			RollSize = ?, 
			Time = ?
		WHERE
			RollID = ?
	`
	if err := tx.Exec(query,
		req.Temp,
		req.RollSize,
		req.Time,
		req.RollID,
	).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to execute update query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return req, nil
}

func (s *PFCRollService) DeletePFCRoll(req *types.PFC_Roll) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_Roll
		WHERE RollID = ?
	`

	if err := tx.Exec(query,
		req.RollID,
	).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to execute delete query: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return "delete success", nil
}

func (s *PFCRollService) DeletePFCRollByModelID(req *types.PFCModel) (string, error) {
	db, err := database.DatabaseConnection()
	if err != nil {
		return "", fmt.Errorf("failed to connect database: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	query := `
		DELETE FROM PFC_Roll 
		WHERE LaminationProcessID = (SELECT TOP(1) LaminationProcessID FROM PFC_LaminationProcess WHERE ModelType = ? AND ModelName = ? AND MaterialNumber = ?)
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
