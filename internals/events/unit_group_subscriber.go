package events

import (
	"fmt"
	"log/slog"

	"ithumans.com/coproxpert/data/models"
)

type UnitGroupSubscriber struct{}

func (ugs *UnitGroupSubscriber) EntityName() {
	fmt.Println("unit_group")
}

func (ugs *UnitGroupSubscriber) HandleMessage(message PubSubMessage) {
	switch message.EventType {

	case Created:
		//services.CreatePermission(message.UserID, message.EntityID, models.AdminRole, models.UnitGroupEntity)
		slog.Info("Permission created", "user_id", message.UserID, "entity_id", message.EntityID, "role", models.AdminRole)
	case Deleted:
		//err := services.DeletePermission(message.UserID, message.EntityID)
		//if err != nil {
		//	slog.Error("Error deleting permission", err)
		//}
		slog.Info("Permission deleted", "user_id", message.UserID, "entity_id", message.EntityID)

	}
}
