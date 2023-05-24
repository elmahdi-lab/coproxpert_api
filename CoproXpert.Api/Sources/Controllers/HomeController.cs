// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Microsoft.AspNetCore.Mvc;

namespace CoproXpert.Api.Sources.Controllers;

/// <summary>
/// </summary>
[ApiController]
[Route("[controller]", Name = "HomeRoute")]
public class HomeController : ControllerBase
{
    private readonly Translator _translator;

    /// <summary>
    /// </summary>
    /// <param name="translator"></param>
    public HomeController(Translator translator)
    {
        _translator = translator;
    }

    // GET /welcome
    /// <summary>
    /// </summary>
    /// <returns></returns>
    [HttpGet]
    public ActionResult Index()
    {
        return Ok("ok");
    }
}
