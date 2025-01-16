package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

// TODO: in case we limit the unit groups and units, consider counting

func CheckSubscriptionLimit(limitType models.SubscriptionLimitType) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)

		if user.IsSuperAdmin() {
			slog.Info("Super admin bypassed the subscription limit", "userID:", user.ID)
			return c.Next()
		}

		//if limitType == models.OrganizationLimit {
		//	// Count the number of permissions where the role is an admin from user.Permissions:
		//	organizationsCount := 0
		//	for _, permission := range user.Permissions {
		//		if permission.OrganizationID != uuid.Nil && permission.Role == models.AdminRole {
		//			organizationsCount++
		//		}
		//	}
		//	if
		//}

		organizationID := c.Params("organizationID")
		organizationUUID, _ := uuid.Parse(organizationID)

		organizationRepository := repositories.NewOrganizationRepository()
		organization, err := organizationRepository.FindByID(organizationUUID)

		if err != nil || organization == nil {
			slog.Error("Organization not found")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Organization not found"})
		}

		//if !services.HasPermission(user.ID, models.OrganizationEntity, organization.ID, models.AdminRole) {
		//	slog.Error("User has no permission to access the organization", "userID:", user.ID, "orgID:", organization.ID)
		//	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User has no permission to access the organization"})
		//}

		if organization.Subscription == nil {
			slog.Error("No subscription found")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "No subscription found"})
		}

		subscription := models.SubscriptionTiers[organization.Subscription.SubscriptionType]

		switch limitType {
		case models.UnitLimit:
			err := countUnitsAndCheckLimit(c, organization, subscription)
			if err != nil {
				return err
			}
		case models.UnitGroupLimit:
			err := countUnitGroupsAndCheckLimit(c, organization, subscription)
			if err != nil {
				return err
			}
		}

		return c.Next()
	}
}

func countUnitsAndCheckLimit(c *fiber.Ctx, organization *models.Organization, subscription models.SubscriptionTier) error {
	unitRepo := repositories.NewUnitRepository()
	units, err := unitRepo.CountByOrganizationID(organization.ID)
	if err != nil {
		slog.Error("Error counting units by organization ID")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Error counting units by organization ID"})
	}
	if units >= subscription.UnitsLimit {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit limit reached"})
	}
	return nil
}

func countUnitGroupsAndCheckLimit(c *fiber.Ctx, organization *models.Organization, subscription models.SubscriptionTier) error {
	unitGroupRepo := repositories.NewUnitGroupRepository()
	unitGroups, err := unitGroupRepo.CountByOrganizationID(organization.ID)
	if err != nil {
		slog.Error("Error counting unit groups by organization ID")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Error counting unit groups by organization ID"})
	}
	if unitGroups >= subscription.UnitGroupsLimit {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit group limit reached"})
	}
	return nil
}
