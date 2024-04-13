package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
	"ithumans.com/coproxpert/services"
	"log/slog"
)

func CheckSubscriptionLimit(su models.SubscriptionLimitType) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		user := c.Locals("user").(*models.User)

		// super admin bypasses the subscription limit
		if user != nil {
			if user.IsSuperAdmin() {
				slog.Info("Super admin bypassed the subscription limit", "userID:", user.ID)
				return c.Next()
			}
		}

		organizationID := c.Params("organizationID")
		organizationUUID, _ := uuid.Parse(organizationID)

		organizationRepository := repositories.NewOrganizationRepository()
		organization, err := organizationRepository.FindByID(organizationUUID)

		if organization == nil || err != nil {
			slog.Error("Organization not found")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Organization not found"})
		}

		// check if user has permission to access the organization
		if !services.HasPermission(user.ID, models.OrganizationEntity, organization.ID, models.AdminRole) {
			slog.Error("User has no permission to access the organization", "userID:", user.ID, "orgID:", organization.ID)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User has no permission to access the organization"})
		}

		if organization.Subscription == nil {
			slog.Error("No subscription found")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "No subscription found"})
		}

		// if subscription is enterprise, return false
		if organization.Subscription.SubscriptionType == models.Enterprise {
			return c.Next()
		}

		// if subscription is not enterprise, check the limit
		switch su {

		case models.UnitLimit:
			unitRepo := repositories.NewUnitRepository()
			units, err := unitRepo.CountByOrganizationID(organization.ID)
			if err != nil {
				slog.Error("Error counting units by organization ID")
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Error counting units by organization ID"})
			}
			if units >= models.SubscriptionTiers[organization.Subscription.SubscriptionType].UnitsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit limit reached"})
			}
			return c.Next()

		case models.UnitGroupLimit:
			unitGroupRepo := repositories.NewUnitGroupRepository()
			unitGroups, err := unitGroupRepo.CountByOrganizationID(organization.ID)
			if err != nil {
				slog.Error("Error counting unit groups by organization ID")
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Error counting unit groups by organization ID"})
			}
			if unitGroups >= models.SubscriptionTiers[organization.Subscription.SubscriptionType].UnitGroupsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit group limit reached"})
			}
			return c.Next()
		}

		return c.Next()
	}
}
