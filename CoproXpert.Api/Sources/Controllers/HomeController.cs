// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Api.Sources.Authentication;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace CoproXpert.Api.Sources.Controllers;

/// <summary>
///     Home controller for handling welcome endpoint.
/// </summary>
[ApiController]
[Route("[controller]", Name = "HomeRoute")]
public class HomeController : ControllerBase
{
    private readonly Translator _translator;

    /// <summary>
    ///     Initializes a new instance of the <see cref="HomeController" /> class.
    /// </summary>
    /// <param name="translator">The translator instance.</param>
    public HomeController(Translator translator)
    {
        _translator = translator;
    }

    /// <summary>
    ///     Retrieves a welcome message.
    /// </summary>
    /// <returns>The welcome message.</returns>
    [HttpGet("index")]
    [ServiceFilter(typeof(ApiKeyAuthFilterAttribute))]
    [Authorize]
    public ActionResult Index()
    {
        var message = _translator.Translate("index.welcome");
        return Ok(message);
    }
}
