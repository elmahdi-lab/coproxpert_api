package events

import (
	"fmt"
	"log/slog"

	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

type UnitSubscriber struct{}

func (ugs *UnitSubscriber) EntityName() {
	fmt.Println("unit")
}

func (ugs *UnitSubscriber) HandleMessage(message PubSubMessage) {
	switch message.EventType {

	case Created:
		services.CreatePermission(
			message.UserID, message.EntityID, models.AdminRole, models.UnitEntity)
		slog.Info("Permission created", "user_id", message.UserID, "entity_id", message.EntityID, "role", models.AdminRole)
	case Deleted:
		err := services.DeletePermission(message.UserID, message.EntityID)
		if err != nil {
			slog.Error("Error deleting permission", err)
		}
	}
}
