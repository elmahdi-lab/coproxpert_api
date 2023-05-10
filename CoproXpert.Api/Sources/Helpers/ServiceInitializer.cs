// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database.Models;
using CoproXpert.Database.Services;
using CoProXpert.Database.Models;

namespace CoproXpert.Api.Sources.Helpers;

/// <summary>
///   Initializes the services
/// </summary>
public static class ServiceInitializer
{
    private static readonly List<string> s_targetNamespaces = new() { "CoproXpert.Database.Repositories" };

    private static readonly List<Type> s_excludedTypes = new()
    {
        typeof(BaseModel), typeof(BaseService<>), typeof(IModel)
    };

    /// <summary>
    ///    Initializes the services in the specified service collection
    /// </summary>
    /// <param name="serviceCollection"></param>
    public static void Init(IServiceCollection serviceCollection)
    {
        // Get all the services in the target namespace
        var services = GetServicesInNamespace(s_targetNamespaces, s_excludedTypes);

        // Loop through the services and add them to the service collection
        foreach (var service in services)
        {
            serviceCollection.AddScoped(service);
        }
    }

    private static List<Type> GetServicesInNamespace(ICollection<string> targetNamespaces,
        ICollection<Type> excludedTypes)
    {
        // Get all the currently loaded assemblies
        var assemblies = AppDomain.CurrentDomain.GetAssemblies();

        // Get all the types in the assemblies that match any of the target namespaces and exclude the specified types
        var services = assemblies
            .SelectMany(assembly => assembly.GetTypes())
            .Where(type => type.Namespace != null
                           && targetNamespaces.Any(targetNamespace => type.Namespace.Equals(targetNamespace))
                           && !excludedTypes.Contains(type))
            .ToList();

        return services;
    }
}
