// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models;
using CoproXpert.Database.Repositories;
using Microsoft.AspNetCore.Mvc;

namespace CoproXpert.Api.Sources.Controllers;

/// <summary>
///     Controller for managing organizations.
/// </summary>
[ApiController]
[Route("[controller]", Name = "OrganizationRoute")]
public class OrganizationController : ControllerBase
{
    private readonly OrganizationRepository _organizationRepository;

    /// <summary>
    ///     Initializes a new instance of the <see cref="OrganizationController" /> class.
    /// </summary>
    /// <param name="organizationRepository">The organization repository.</param>
    public OrganizationController(OrganizationRepository organizationRepository)
    {
        _organizationRepository = organizationRepository;
    }

    /// <summary>
    ///     Gets all organizations.
    /// </summary>
    /// <returns>The list of organizations.</returns>
    [HttpGet(Name = "GetAllOrganizations")]
    public ActionResult<IEnumerable<Organization>> GetAll()
    {
        return _organizationRepository.GetAll().ToList();
    }

    /// <summary>
    ///     Gets an organization by its ID.
    /// </summary>
    /// <param name="organizationId">The ID of the organization.</param>
    /// <returns>The organization with the specified ID.</returns>
    [HttpGet("{organizationId}", Name = "GetOrganizationById")]
    public ActionResult<Organization> Get(Guid organizationId)
    {
        var organization = _organizationRepository.GetById(organizationId);

        if (organization == null)
        {
            return NotFound();
        }

        return organization;
    }

    /// <summary>
    ///     Creates a new organization.
    /// </summary>
    /// <param name="organization">The organization to create.</param>
    /// <returns>The created organization.</returns>
    [HttpPost("{organization}", Name = "CreateOrganization")]
    public ActionResult<Organization> Create(Organization organization)
    {
        organization = _organizationRepository.Create(organization);
        if (organization.Id == Guid.Empty)
        {
            return BadRequest();
        }

        return CreatedAtAction(nameof(Get), new { id = organization.Id }, organization);
    }

    /// <summary>
    ///     Updates an existing organization.
    /// </summary>
    /// <param name="organization">The updated organization.</param>
    /// <returns>
    ///     No content if the update is successful, bad request if the IDs do not match, or not found if the organization
    ///     does not exist.
    /// </returns>
    [HttpPut("{organization}", Name = "UpdateOrganization")]
    public IActionResult Update(Organization organization)
    {
        var success = _organizationRepository.Update(organization);

        if (!success)
        {
            return NotFound();
        }

        return NoContent();
    }

    /// <summary>
    ///     Deletes an organization.
    /// </summary>
    /// <param name="organizationId">The ID of the organization to delete.</param>
    /// <returns>No content if the deletion is successful, or not found if the organization does not exist.</returns>
    [HttpDelete("{organizationId}", Name = "DeleteOrganization")]
    public IActionResult Delete(Guid organizationId)
    {
        var success = _organizationRepository.Delete(organizationId);

        if (!success)
        {
            return NotFound();
        }

        return NoContent();
    }
}
