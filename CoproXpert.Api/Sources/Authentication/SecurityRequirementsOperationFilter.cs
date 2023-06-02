// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Microsoft.AspNetCore.Authorization;
using Microsoft.OpenApi.Models;
using Swashbuckle.AspNetCore.SwaggerGen;

namespace CoproXpert.Api.Sources.Authentication;

/// <summary>
///     Security requirements operation filter, used to add security requirements to swagger.
/// </summary>
public class SecurityRequirementsOperationFilter : IOperationFilter
{
    /// <summary>
    /// </summary>
    /// <param name="operation"></param>
    /// <param name="context"></param>
    public void Apply(OpenApiOperation operation, OperationFilterContext context)
    {
        var authorizeAttributes = context.MethodInfo.GetCustomAttributes(true)
            .OfType<AuthorizeAttribute>()
            .ToList();

        if (authorizeAttributes.Any())
        {
            operation.Security = new List<OpenApiSecurityRequirement>
            {
                new()
                {
                    {
                        new OpenApiSecurityScheme
                        {
                            Reference = new OpenApiReference
                            {
                                Type = ReferenceType.SecurityScheme, Id = "ApiKey"
                            }
                        },
                        authorizeAttributes.Select(attr => attr.Policy).ToList()
                    }
                }
            };
        }
    }
}
