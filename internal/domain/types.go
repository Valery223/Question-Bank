package domain

import (
	"errors"
)

// Ошибки домена
var (
	ErrInvalidDifficulty = errors.New("difficulty must be between 1 and 5")
	ErrEmptyText         = errors.New("text cannot be empty")
	ErrNoOptions         = errors.New("choice question must have at least 2 options")
	ErrInvalidType       = errors.New("unknown question type")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
)

// ID - строгий тип для идентификаторов сущностей
//
// Везде используется string(UUID) для простоты
type ID string

// --- Enums ---

// QuestionType
type QuestionType string

const (
	TypeSingleChoice QuestionType = "single_choice"
	TypeMultiChoice  QuestionType = "multi_choice"
	TypeText         QuestionType = "text"
)

// IsValid проверяет, является ли тип вопроса допустимым
func (qt QuestionType) IsValid() bool {
	switch qt {
	case TypeSingleChoice, TypeMultiChoice, TypeText:
		return true
	default:
		return false
	}
}

// RoleQuestionnaire - роль, для которой предназначен тестовый шаблон
type RoleQuestionnaire string

const (
	RoleBackendJunior  RoleQuestionnaire = "backend_junior"
	RoleFrontendJunior RoleQuestionnaire = "frontend_junior"
	// Можно добавить новые
)

// IsValid проверяет, что роль является допустимой
func (r RoleQuestionnaire) IsValid() bool {
	switch r {
	case RoleBackendJunior, RoleFrontendJunior:
		return true
	default:
		return false
	}
}

// TemplatePurpose - цель использования тестового шаблона
type TemplatePurpose string

const (
	PurposeAssessment    TemplatePurpose = "skills_assessment"
	PurposeMockInterview TemplatePurpose = "mock_interview"
)

// IsValid проверяет, что цель является допустимой
func (tp TemplatePurpose) IsValid() bool {
	switch tp {
	case PurposeAssessment, PurposeMockInterview:
		return true
	default:
		return false
	}
}

// Difficulty - делаем строгим типом, чтобы нельзя было случайно передать просто int
type Difficulty int

const (
	MinDifficulty Difficulty = 1
	MaxDifficulty Difficulty = 5
)

// IsValid проверяет, что уровень сложности в допустимых пределах
func (d Difficulty) IsValid() bool {
	return d >= MinDifficulty && d <= MaxDifficulty
}
