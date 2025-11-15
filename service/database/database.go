package database

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	TransactionSchemaEnumTypeName = "transaction_types"
	CategorySchemaEnumTypeName    = "category_types"
)

type TransactionType string

const (
	TransactionSchemaEnumTypeIncome  TransactionType = "INCOME"
	TransactionSchemaEnumTypeExpense TransactionType = "EXPENSE"
)

type CategoryType string

const (
	CategorySchemaEnumTypeIncome        CategoryType = "INCOME"
	CategorySchemaEnumTypeFood          CategoryType = "FOOD"
	CategorySchemaEnumTypeTransport     CategoryType = "TRANSPORT"
	CategorySchemaEnumTypeShopping      CategoryType = "SHOPPING"
	CategorySchemaEnumTypeEntertainment CategoryType = "ENTERTAINMENT"
	CategorySchemaEnumTypeHealth        CategoryType = "HEALTH"
	CategorySchemaEnumTypeOther         CategoryType = "OTHER"
)

type PersonalFinancialTracking struct {
	ID              uuid.UUID       `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey;index:,type:btree" json:"ID"`
	TransactionType TransactionType `gorm:"column:transaction_type;type:transaction_types;not null"`
	Category        CategoryType    `gorm:"column:category;type:category_types;not null"`
	Amount          float64         `gorm:"column:amount;type:decimal(17,2);not null" json:"amount"`
	Note            string          `gorm:"column:note;type:varchar;size:100" json:"note"`
	CreatedAt       time.Time       `gorm:"column:created_at;type:timestamp WITH TIME ZONE; not null;default:now()" json:"createdAt"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;type:timestamp WITH TIME ZONE; not null;default:now()" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt  `gorm:"column:deleted_at;index" json:"deletedAt"`
}

func (PersonalFinancialTracking) TableName() string {
	return "personal_financial_trackings"
}

type PersonalFinancialTrackingTargetSpendPerMonth struct {
	ID          uuid.UUID      `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey;index:,type:btree" json:"ID"`
	Category    CategoryType   `gorm:"column:category;type:category_types;not null" json:"category"`
	SpendAmount float64        `gorm:"column:spend_amount;type:decimal(17,2);not null" json:"spendAmount"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp WITH TIME ZONE; not null;default:now()" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:timestamp WITH TIME ZONE; not null;default:now()" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deletedAt"`
}

func (PersonalFinancialTrackingTargetSpendPerMonth) TableName() string {
	return "tracking_target_spend_per_months"
}

type EnumValidationSchema struct {
	IsExit string `gorm:"column:is_exit;"`
}

func MigrateDB(dbConn *gorm.DB) (err error) {
	migrateEnumSchema(dbConn)
	err = dbConn.AutoMigrate(
		&PersonalFinancialTracking{},
		&PersonalFinancialTrackingTargetSpendPerMonth{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func migrateEnumSchema(gormDB *gorm.DB) {
	// Vendor type
	enumTransactionTypeValidation := &EnumValidationSchema{}
	gormDB.Raw(`SELECT 1 as is_exit FROM pg_type WHERE typname = ?`, TransactionSchemaEnumTypeName).
		Scan(&enumTransactionTypeValidation)
	if enumTransactionTypeValidation.IsExit == "" {
		res := gormDB.Exec(`CREATE TYPE ` + TransactionSchemaEnumTypeName + ` AS ENUM (` + fmt.Sprintf(
			`'%s','%s'`,
			TransactionSchemaEnumTypeIncome,
			TransactionSchemaEnumTypeExpense,
		) + `);`)
		if res.Error != nil {
			log.Fatal("failed to create enum type", res.Error)
		}
	}

	enumCategoryTypeValidation := &EnumValidationSchema{}
	gormDB.Raw(`SELECT 1 as is_exit FROM pg_type WHERE typname = ?`, CategorySchemaEnumTypeName).
		Scan(&enumCategoryTypeValidation)
	if enumCategoryTypeValidation.IsExit == "" {
		res := gormDB.Exec(`CREATE TYPE ` + CategorySchemaEnumTypeName + ` AS ENUM (` + fmt.Sprintf(
			`'%s','%s','%s','%s','%s','%s','%s'`,
			CategorySchemaEnumTypeIncome,
			CategorySchemaEnumTypeFood,
			CategorySchemaEnumTypeTransport,
			CategorySchemaEnumTypeShopping,
			CategorySchemaEnumTypeEntertainment,
			CategorySchemaEnumTypeHealth,
			CategorySchemaEnumTypeOther,
		) + `);`)
		if res.Error != nil {
			log.Fatal("failed to create enum type", res.Error)
		}
	}
}
