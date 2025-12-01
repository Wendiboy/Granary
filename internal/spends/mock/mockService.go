package mockSpends_test

import (
	"fmt"
	"strconv"

	spendsService "finance/internal/spends/service"

	"github.com/google/uuid"
)

var DB []spendsService.RawSpend

func MockServiceGetSpend(id uuid.UUID) (spendsService.RawSpend, error) {

	rawSpend := spendsService.RawSpend{
		Id:         id,
		Account:    "Tbank",
		Category:   "food",
		Amount:     2123.53,
		Currency:   "RUB",
		Labels:     "SPAR",
		Note:       "some text",
		Date:       "2025-12-01 16:25 4",
		Created_at: "timestamp with timezone",
		Updated_at: "timestamp with timezone",
		Deleted_at: "timestamp with timezone",
	}
	fmt.Println("SRV:", rawSpend)

	return rawSpend, nil
}

func MockServiceGetAllSpends() ([]spendsService.RawSpend, error) {
	rawSpends := make([]spendsService.RawSpend, 0, 10)

	for i := 0; i < 10; i++ {
		id, _ := uuid.NewUUID()

		RSp := spendsService.RawSpend{
			Id:         id,
			Account:    "account" + strconv.Itoa(i),
			Category:   "category" + strconv.Itoa(i),
			Amount:     float64(i),
			Currency:   "RUB",
			Labels:     "labels" + strconv.Itoa(i),
			Note:       "note" + strconv.Itoa(i),
			Date:       "date" + strconv.Itoa(i),
			Created_at: "date" + strconv.Itoa(i),
			Updated_at: "date" + strconv.Itoa(i),
			Deleted_at: "date" + strconv.Itoa(i),
		}
		fmt.Println("service:", RSp)
		rawSpends = append(rawSpends, RSp)
	}
	fmt.Println("service: ", rawSpends)
	return rawSpends, nil
}

func MockCreateSpend(spend spendsService.Spend) (spendsService.RawSpend, error) {
	id, _ := uuid.NewUUID()
	rawSpend := spendsService.RawSpend{
		Id:         id,
		Account:    spend.Account,
		Category:   spend.Category,
		Amount:     spend.Amount,
		Currency:   spend.Currency,
		Labels:     spend.Labels,
		Note:       spend.Note,
		Date:       spend.Date,
		Created_at: "timestamp with timezone",
		Updated_at: "timestamp with timezone",
		Deleted_at: "timestamp with timezone",
	}
	fmt.Println("SRV Create:", rawSpend)

	DB = append(DB, rawSpend)

	return rawSpend, nil
}

func MockUpdateSpend(id uuid.UUID, spend spendsService.Spend) (spendsService.RawSpend, error) {
	rawSpend := spendsService.RawSpend{
		Id:         id,
		Account:    spend.Account,
		Category:   spend.Category,
		Amount:     spend.Amount,
		Currency:   spend.Currency,
		Labels:     spend.Labels,
		Note:       spend.Note,
		Date:       spend.Date,
		Created_at: "timestamp with timezone",
		Updated_at: "timestamp with timezone",
		Deleted_at: "timestamp with timezone",
	}
	fmt.Println("SRV Update:", rawSpend)

	DB = append(DB, rawSpend)

	return rawSpend, nil
}

func MockDeleteSpend(id uuid.UUID) error {
	fmt.Println("SRV DELETE:", id)

	return nil
}
