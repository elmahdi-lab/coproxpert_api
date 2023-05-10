// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models;
using CoproXpert.Database.Services;
using Microsoft.AspNetCore.Mvc;

namespace CoproXpert.Api.Sources.Controllers;

/// <summary>
/// </summary>
[ApiController]
[Route("[controller]", Name = "UserRoute")]
public class UserController : ControllerBase
{
    private readonly UserService _userService;

    /// <summary>
    /// </summary>
    /// <param name="userService"></param>
    public UserController(UserService userService)
    {
        _userService = userService;
    }

    // GET /user
    [HttpGet]
    public ActionResult<IEnumerable<User>> GetAll()
    {
        return _userService.GetAll().ToList();
    }

    // GET /user/1
    [HttpGet("{id}")]
    public ActionResult<User> Get(int id)
    {
        var user = _userService.GetById(id);

        if (user == null)
        {
            return NotFound();
        }

        return user;
    }

    // POST /user
    [HttpPost]
    public ActionResult<User> Create(User user)
    {
        _userService.Create(user);
        return CreatedAtAction(nameof(Get), new { id = user.Id }, user);
    }

    // PUT /user/1
    [HttpPut("{id}")]
    public IActionResult Update(int id, User user)
    {
        if (id != user.Id)
        {
            return BadRequest();
        }

        var success = _userService.Update(user);

        if (!success)
        {
            return NotFound();
        }

        return NoContent();
    }

    // DELETE /user/1
    [HttpDelete("{id}")]
    public IActionResult Delete(int id)
    {
        var success = _userService.Delete(id);

        if (!success)
        {
            return NotFound();
        }

        return NoContent();
    }
}
