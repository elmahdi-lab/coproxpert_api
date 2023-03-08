using Microsoft.AspNetCore.Mvc;

namespace CoproXpert.Sources.Controllers;

[ApiController]
[Route("[controller]")]
public class UserController : ControllerBase
{
    [HttpGet(Name = "IndexRoute")]
    public string Index()
    {
        return "Hello World!";
    }
}