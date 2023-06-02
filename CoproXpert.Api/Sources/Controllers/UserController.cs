// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models;
using CoproXpert.Database.Repositories;
using Microsoft.AspNetCore.Mvc;

namespace CoproXpert.Api.Sources.Controllers;

/// <summary>
///     Controller for managing users.
/// </summary>
[ApiController]
[Route("[controller]", Name = "UserRoute")]
public class UserController : ControllerBase
{
    private readonly UserRepository _userRepository;

    /// <summary>
    ///     Initializes a new instance of the <see cref="UserController" /> class.
    /// </summary>
    /// <param name="userRepository">The user repository.</param>
    public UserController(UserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    /// <summary>
    ///     Gets all users.
    /// </summary>
    /// <returns>The list of users.</returns>
    [HttpGet]
    public ActionResult<IEnumerable<User>> GetAll()
    {
        return _userRepository.GetAll().ToList();
    }

    /// <summary>
    ///     Gets a user by its ID.
    /// </summary>
    /// <param name="id">The ID of the user.</param>
    /// <returns>The user with the specified ID.</returns>
    [HttpGet("{id}")]
    public ActionResult<User> Get(Guid id)
    {
        var user = _userRepository.GetById(id);

        if (user == null)
        {
            return NotFound();
        }

        return user;
    }

    /// <summary>
    ///     Creates a new user.
    /// </summary>
    /// <param name="user">The user to create.</param>
    /// <returns>The created user.</returns>
    [HttpPost]
    public ActionResult<User> Create(User user)
    {
        user = _userRepository.Create(user);
        // make sure we have a valid user.Id before returning it
        if (user.Id == Guid.Empty)
        {
            return BadRequest();
        }

        return CreatedAtAction(nameof(Get), new { id = user.Id }, user);
    }

    /// <summary>
    ///     Updates an existing user.
    /// </summary>
    /// <param name="id">The ID of the user to update.</param>
    /// <param name="user">The updated user.</param>
    /// <returns>
    ///     No content if the update is successful, bad request if the IDs do not match, or not found if the user does not
    ///     exist.
    /// </returns>
    [HttpPut("{id}")]
    public IActionResult Update(Guid id, User user)
    {
        if (user.Id != Guid.Empty && id != user.Id)
        {
            return BadRequest();
        }

        var success = _userRepository.Update(user);

        if (!success)
        {
            return NotFound();
        }

        return NoContent();
    }

    /// <summary>
    ///     Deletes a user.
    /// </summary>
    /// <param name="userId">The ID of the user to delete.</param>
    /// <returns>No content if the deletion is successful, or not found if the user does not exist.</returns>
    [HttpDelete("{userId}", Name = "DeleteUser")]
    public IActionResult Delete(Guid userId)
    {
        var success = _userRepository.Delete(userId);

        if (!success)
        {
            return NotFound();
        }

        return NoContent();
    }
}
