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

// Пусть  везде будет string ID
type ID string

// --- Enums ---

type QuestionType string

const (
	TypeSingleChoice QuestionType = "single_choice"
	TypeMultiChoice  QuestionType = "multi_choice"
	TypeText         QuestionType = "text"
)

func (qt QuestionType) IsValid() bool {
	switch qt {
	case TypeSingleChoice, TypeMultiChoice, TypeText:
		return true
	default:
		return false
	}
}

type RoleQuestionnaire string

const (
	RoleBackendJunior  RoleQuestionnaire = "backend_junior"
	RoleFrontendJunior RoleQuestionnaire = "frontend_junior"
	// Можно добавить новые
)

func (r RoleQuestionnaire) IsValid() bool {
	switch r {
	case RoleBackendJunior, RoleFrontendJunior:
		return true
	default:
		return false
	}
}

type TemplatePurpose string

const (
	PurposeAssessment    TemplatePurpose = "skills_assessment"
	PurposeMockInterview TemplatePurpose = "mock_interview"
)

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

// Вспомогательный метод для проверки
func (d Difficulty) IsValid() bool {
	return d >= MinDifficulty && d <= MaxDifficulty
}
